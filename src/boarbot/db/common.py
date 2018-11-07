from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm.session import Session
from sqlalchemy.ext.declarative import declarative_base

from boarbot.common.log import LOGGER
from boarbot.common.config import DATABASE

Base = declarative_base()

class BoarbotDatabase:
    def __init__(self, engine):
        self.engine = engine
        self.sessionmaker = sessionmaker(bind=self.engine)

    def get_session(self) -> Session:
        return self.sessionmaker()

    __instance = None  # type: BoarbotDatabase
    @staticmethod
    def get_instance():
        if not BoarbotDatabase.__instance:
            BoarbotDatabase.initialize()
        return BoarbotDatabase.__instance

    @staticmethod
    def initialize():
        LOGGER.info('Database lazy initialization...')
        dburl = DATABASE['url']
        if dburl.startswith('postgres://'):
            dburl = 'postgres+psycopg2' + dburl[8:]

        engine = create_engine(dburl, echo=DATABASE['echo'])
        Base.metadata.create_all(engine)
        BoarbotDatabase.__instance = BoarbotDatabase(engine)
        LOGGER.info('Database lazy initialization complete.s')
