from datetime import datetime
from sqlalchemy.orm.session import Session

from boarbot.common.config import BOOTSTRAP_OVERRIDE
from boarbot.common.log import LOGGER
from boarbot.db.common import BoarbotDatabase
from boarbot.db.kv import KV

from .dbdebug import bootstrap_dbdebug_users
from .groups import bootstrap_group_managers

BOOTSTRAP_KEY = "__BOOTSTRAP__"
BOOTSTRAP_VERSION = "1"

def _check_bootstrap_needed(db_session: Session):
    if BOOTSTRAP_OVERRIDE:
        return True
    kv = db_session.query(KV).filter(KV.key==BOOTSTRAP_KEY).one_or_none()
    return (not kv) or kv.value != BOOTSTRAP_VERSION
        
def _save_bootstrap_version(db_session: Session):
    kv = db_session.query(KV).filter(KV.key == BOOTSTRAP_KEY).one_or_none()
    if not kv:
        kv = KV(key=BOOTSTRAP_KEY, value=BOOTSTRAP_VERSION, timestamp=datetime.utcnow())
        db_session.add(kv)
    else:
        kv.value = BOOTSTRAP_VERSION
        kv.timestamp = datetime.utcnow()

def run_bootstrap(db_session: Session):
    if not _check_bootstrap_needed(db_session):
        LOGGER.info("Bootstrap not needed, skipping")
        return
    LOGGER.info("Bootstrapping legacy data into database")


    bootstrap_dbdebug_users(db_session)
    bootstrap_group_managers(db_session)

    _save_bootstrap_version(db_session)
    db_session.commit()
    LOGGER.info("Bootstraping legacy data complete")
