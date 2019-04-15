from datetime import datetime
from sqlalchemy.orm.session import Session

from boarbot.common.config import DATABASE
from boarbot.common.log import LOGGER
from boarbot.db.acl import UserACLEntry
from boarbot.modules.dbdebug.module import DEBUG_USERS_ACL


def bootstrap_dbdebug_users(db_session: Session):
    debug_users = DATABASE['debugUsers']
    if not debug_users:
        LOGGER.warn('Unable to bootstrap debug_users due to missing config')
        return

    
    ids_known = [entry.user_id for entry in db_session.query(
        UserACLEntry).filter(UserACLEntry.acl_id == DEBUG_USERS_ACL)]

    count = 0
    for user_id in [user_id for user_id in debug_users if user_id not in ids_known]:
        acl = UserACLEntry(acl_id=DEBUG_USERS_ACL, user_id=user_id,
                           details=f'bootstrapped on {datetime.utcnow().isoformat()}')
        db_session.add(acl)
        count += 1

    LOGGER.info(f"Bootstrapped {count} users as dbdebug users")
