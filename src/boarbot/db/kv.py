from sqlalchemy import Column, DateTime, String

from boarbot.db.common import Base

class KV(Base):
    __tablename__ = 'kv'

    key = Column(String, primary_key=True)
    value = Column(String)
    timestamp = Column(DateTime, nullable=False)

    def __repr__(self) -> str:
        return f'<KV(K=`{self.key}`, v=`{self.value}`, t=`{self.timestamp}`)>'