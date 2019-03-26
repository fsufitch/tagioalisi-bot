import discord
from sqlalchemy import Column, Index, Integer, Sequence, String
from sqlalchemy.orm.session import Session

from boarbot.db.common import Base

class UserACLEntry(Base):
    __tablename__ = 'user_acl'

    row_id = Column(Integer, Sequence('seq_user_acl_row_id'), primary_key=True)
    acl_id = Column(String, nullable=False)
    user_id = Column(String, nullable=False)
    details = Column(String)

    def __repr__(self) -> str:
        return f'<UserACL(id=`{self.row_id}`, acl_id=`{self.acl_id}`, user_id=`{self.user_id}`, details=`{self.details}`)>'

    __table_args__ = (
        Index(f'idx_{__tablename__}__acl_user', acl_id, user_id),
    )

class RoleACLEntry(Base):
    __tablename__ = 'role_acl'

    row_id = Column(Integer, Sequence('seq_role_acl_row_id'), primary_key=True)
    acl_id = Column(String, nullable=False)
    role_id = Column(String, nullable=False)
    details = Column(String)

    def __repr__(self) -> str:
        return f'<RoleACL(id=`{self.row_id}`, acl_id=`{self.acl_id}`, role_id=`{self.role_id}`, details=`{self.details}`)>'

    __table_args__ = (
        Index(f'idx_{__tablename__}__acl_group', acl_id, role_id),
    )

def check_acl_user(db_session: Session, acl_id: str, user_id: str) -> bool:
    return db_session.query(UserACLEntry).filter(
        UserACLEntry.acl_id == acl_id,
        UserACLEntry.user_id == user_id,
    ).count() > 0

def check_acl_roles(db_session: Session, acl_id: str, role_ids: [str]) -> [str]:
    return [entry.role_id for entry in db_session.query(RoleACLEntry).filter(
        RoleACLEntry.acl_id == acl_id,
        RoleACLEntry.role_id.in_(role_ids),
    )]
    
def create_acl_user(db_session: Session, acl_id: str, user_id: str, details: str = '') -> UserACLEntry:
    entry = UserACLEntry(acl_id=acl_id, user_id=user_id, details=details)
    db_session.add(entry)
    return entry

def create_acl_role(db_session: Session, acl_id: str, role_id: str, details: str = '') -> RoleACLEntry:
    entry = RoleACLEntry(acl_id=acl_id, role_id=role_id, details=details)
    db_session.add(entry)
    return entry

def delete_acl_user(db_session: Session, acl_id: str, user_id:str):
    entry = db_session.query(UserACLEntry).filter(
        UserACLEntry.acl_id == acl_id,
        UserACLEntry.user_id == user_id,
    ).one_or_none()
    if entry:
        db_session.delete(entry)

def delete_acl_role(db_session: Session, acl_id: str, role_id: str):
    entry = db_session.query(RoleACLEntry).filter(
        RoleACLEntry.acl_id == acl_id,
        RoleACLEntry.role_id == role_id,
    ).one_or_none()
    if entry:
        db_session.delete(entry)
