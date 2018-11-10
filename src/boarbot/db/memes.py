from datetime import datetime
from sqlalchemy import Column, DateTime, ForeignKey, Index, Integer, String
from sqlalchemy.orm.session import Session
from sqlalchemy.orm import relationship

from boarbot.db.common import Base

class MemeAlreadyExistsException(Exception): pass

class Meme(Base):
    __tablename__ = 'memes'
    
    id = Column(Integer, primary_key=True)
    names = relationship('MemeName', back_populates='meme',
        cascade='all, delete, delete-orphan')
    urls = relationship('MemeURL', back_populates='meme',
        cascade='all, delete, delete-orphan')

    def __repr__(self) -> str:
        return f'<Meme(id=`{self.id}`)>'

class MemeName(Base):
    __tablename__ = 'meme_names'

    id = Column(Integer, primary_key=True)    
    name = Column(String, unique=True)
    timestamp = Column(DateTime, nullable=False)
    author = Column(String, nullable=False)
    meme = relationship('Meme', back_populates='names')
    meme_id = Column(Integer, ForeignKey('memes.id'))

class MemeURL(Base):
    __tablename__ = 'meme_urls'

    id = Column(Integer, primary_key=True)
    url = Column(String, nullable=False)
    timestamp = Column(DateTime, nullable=False)
    author = Column(String, nullable=False)
    meme = relationship('Meme', back_populates='urls')
    meme_id = Column(Integer, ForeignKey('memes.id'))

def get_meme(db_session: Session, name: str) -> Meme:
    meme = db_session.query(Meme).join(MemeName).filter(MemeName.name == name.lower()).one_or_none()
    if not meme:
        return None
    return meme

def search_memes(db_session: Session, query: str) -> [Meme]:
    return db_session.query(Meme).join(MemeName).filter(
        MemeName.name.ilike(f'%{query.lower()}%'),
    ).all() 

def new_meme(db_session: Session, name: str, url: str, author: str) -> Meme:
    timestamp = datetime.utcnow()

    if get_meme(db_session, name):
        # Already exists, error out
        raise MemeAlreadyExistsException()

    meme = Meme(
        names = [
            MemeName(name=name, timestamp=timestamp, author=author),
        ],
        urls = [
            MemeURL(url=url, timestamp=timestamp, author=author),
        ],
    )
    db_session.add(meme)
    return meme

def add_url(db_session, name: str, url: str, author: str):
    timestamp = datetime.utcnow()
    meme = get_meme(db_session, name)
    if not meme:
        raise KeyError()

    meme.urls.append(MemeURL(url=url, timestamp=timestamp, author=author))

def add_alias(db_session, name: str, new_name: str, author: str):
    timestamp = datetime.utcnow()
    meme = get_meme(db_session, name)
    if not meme:
        raise KeyError()

    meme.names.append(MemeName(name=new_name, timestamp=timestamp, author=author))


def delete_meme_name(db_session: Session, name: str):
    meme_name = db_session.query(MemeName).filter(MemeName.name == name.lower()).one_or_none()
    if meme_name:
        meme = get_meme(db_session, name)
        if len(meme.names) == 1:
            db_session.delete(meme)
        db_session.delete(meme_name)

def delete_meme_url(db_session, meme_id: int, url: str):
    meme_url = db_session.query(MemeURL).join(Meme).filter(
        Meme.id == meme_id,
        MemeURL.url == url,
    ).one_or_none()

    if meme_url:
        db_session.delete(meme_url)
