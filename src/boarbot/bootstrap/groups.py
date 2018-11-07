from datetime import datetime
from sqlalchemy.orm.session import Session

from boarbot.common.config import GROUPS
from boarbot.common.log import LOGGER
from boarbot.db.acl import UserACLEntry
from boarbot.modules.groups.module import get_server_group_manager_acl_id

BOOTSTRAP_SERVER_ID = '228286778832846848' # Sociologic Planning Boar

def bootstrap_group_managers(db_session: Session):
    groups_managers = GROUPS['managers']
    if not groups_managers:
        LOGGER.warn('Unable to bootstrap managers due to missing config')
        return

    acl_id = get_server_group_manager_acl_id(BOOTSTRAP_SERVER_ID)

    ids_known = [entry.user_id for entry in db_session.query(UserACLEntry).filter(UserACLEntry.acl_id==acl_id)]

    count = 0
    for user_id in [user_id for user_id in groups_managers if user_id not in ids_known]:
        acl = UserACLEntry(acl_id=acl_id, user_id=user_id, details=f'bootstrapped on {datetime.utcnow().isoformat()}')
        db_session.add(acl)
        count += 1
    
    LOGGER.info(f"Bootstrapped {count} users as group managers")

    

