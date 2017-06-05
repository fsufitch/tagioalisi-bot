import json
import pkg_resources

RAW = pkg_resources.resource_string('boarbot', 'config.json').decode()
CONFIG = json.loads(RAW)
