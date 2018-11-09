import pkg_resources
import yaml
from datetime import datetime
from sqlalchemy.orm.session import Session

from boarbot.common.config import GROUPS
from boarbot.common.log import LOGGER
from boarbot.db.acl import UserACLEntry, RoleACLEntry
from boarbot.db.memes import new_meme, MemeName, MemeURL
from boarbot.modules.memelink.module import MEME_EDIT_ACL_ID

def bootstrap_memes(db_session: Session):
    _bootstrap_memes_data(db_session)
    _bootstrap_memes_acl(db_session)

def _bootstrap_memes_data(db_session: Session):
    raw_memes = yaml.load(pkg_resources.resource_stream(
        'boarbot.bootstrap', 'memes.yaml'))

    LOGGER.info(f'Bootstrapping {len(raw_memes)} memes...')
    for meme in raw_memes:
        names = meme['names'] 
        urls = meme['urls']

        name1, names = names[0], names[1:]
        url1, urls = urls[0], urls[1:]

        try:
            meme = new_meme(db_session, name1, url1, 'database bootstrap')
        except Exception as e:
            LOGGER.warn(f'... Skipped `{name1}` due to exception: {e}')
            continue

        for name in names:
            meme.names.append(MemeName(
                name=name,
                timestamp=datetime.utcnow(),
                author='database bootstrap',
            ))   
        for url in urls:
            meme.urls.append(MemeURL(
                url=url,
                timestamp=datetime.utcnow(),
                author='database bootstrap',
            ))

    LOGGER.info('Meme bootstrap complete.')


BOOTSTRAP_MEME_USERS = [
    '203684963864805376',  # Blackshell
]
BOOTSTRAP_MEME_ROLES = [
    # TODO
]
def _bootstrap_memes_acl(db_session: Session):
    LOGGER.info("Bootstrapping users and roles for meme editing")
    ids_known = [entry.user_id for entry in db_session.query(
        UserACLEntry).filter(UserACLEntry.acl_id == MEME_EDIT_ACL_ID)]
    
    count = 0
    for user_id in [i for i in BOOTSTRAP_MEME_USERS if i not in ids_known]:
        acl = UserACLEntry(acl_id=MEME_EDIT_ACL_ID, user_id=user_id,
                        details=f'bootstrapped on {datetime.utcnow().isoformat()}')
        db_session.add(acl)
        count += 1

    LOGGER.info(f"Bootstrapped {count} users as meme editors")

    ids_known = [entry.role_id for entry in db_session.query(
        RoleACLEntry).filter(RoleACLEntry.acl_id == MEME_EDIT_ACL_ID)]
    
    count = 0
    for role_id in [i for i in BOOTSTRAP_MEME_ROLES if i not in ids_known]:
        acl = RoleACLEntry(acl_id=MEME_EDIT_ACL_ID, role_id=role_id,
                        details=f'bootstrapped on {datetime.utcnow().isoformat()}')
        db_session.add(acl)
        count += 1
    
    LOGGER.info(f"Bootstrapped {count} roles as meme editors")
