import os
import yaml
import pkg_resources

RAW = yaml.load(pkg_resources.resource_stream('boarbot.common', 'config.json'))

LOGGING = {
    'cliLevel': os.getenv('CLI_LOG_LEVEL') or RAW.get('CLI_LOG_LEVEL'),
    'discordLevel': os.getenv('DISCORD_LOG_LEVEL') or RAW.get('DISCORD_LOG_LEVEL'),
    'channel': os.getenv('DISCORD_LOG_CHANNEL') or str(RAW.get('DISCORD_LOG_CHANNEL') or 0),
}

WELCOME = {
    'channel': os.getenv('WELCOME_CHANNEL') or str(RAW.get('WELCOME_CHANNEL') or 0),
    'rulesChannel': os.getenv('WELCOME_RULES_CHANNEL') or str(RAW.get('WELCOME_RULES_CHANNEL') or 0),
}

GROUPS = {
    'managers': ','.split(os.getenv('GROUPS_MANAGERS') or RAW.get('GROUPS_MANAGERS')),
}

LOAD_MODULES = ','.split(os.getenv('LOAD_MODULES') or RAW.get('LOAD_MODULES'))
