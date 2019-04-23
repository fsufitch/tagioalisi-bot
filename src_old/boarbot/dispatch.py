import importlib
import traceback

from boarbot.common.config import LOAD_MODULES
from boarbot.db.common import BoarbotDatabase
from boarbot.common.events import EventType
from boarbot.common.log import LOGGER

MODULES = []

DB = BoarbotDatabase.get_instance()

def initialize_modules(client, extra_modules=[], reinit=False):
    if MODULES and not reinit:
        print("Not re-initializing modules")

    modules = LOAD_MODULES + extra_modules

    for module_config in modules:
        LOGGER.debug('Loading ' + module_config)
        module_name, class_name = module_config.split(':', 1)

        pymodule = importlib.import_module(module_name)
        module_class = getattr(pymodule, class_name)
        module = module_class(client)
        MODULES.append(module)

async def dispatch_event(event_type: EventType, *args):
    for module in MODULES:
        session = DB.get_session()
        try:
            await module.handle_event(session, event_type, args)
        except Exception:
            tb = traceback.format_exc().strip()
            LOGGER.error('```' + tb + '```')
        finally:
            session.rollback()
