import asyncio
import logging, logging.handlers
import sys

from boarbot.common.config import CONFIG

LOG_LEVELS = {
    'DEBUG': logging.DEBUG,
    'INFO': logging.INFO,
    'WARNING': logging.WARNING,
    'ERROR': logging.ERROR,
    'CRITICAL': logging.CRITICAL,
}

LOGGER = logging.getLogger('boarbot')

class AsyncHandler(logging.Handler):
    def __init__(self, callback_async):
        super().__init__()
        self.callback_async = callback_async

    def emit(self, record: logging.LogRecord):
        message = self.format(record)
        asyncio.ensure_future(self.callback_async(message))

def register_discord_handler(callback_async):
    discord_level = _parse_log_level(CONFIG.get('discordLogLevel'))
    discord_formatter = logging.Formatter('[%(levelname)s] %(message)s')
    handler = AsyncHandler(callback_async)
    handler.setLevel(discord_level)
    handler.setFormatter(discord_formatter)
    LOGGER.addHandler(handler)

def _parse_log_level(level: str) -> int:
    if level not in LOG_LEVELS:
        raise KeyError("Log level does not exist %s" % level)
    return LOG_LEVELS[level]

def _setup_logging():
    LOGGER.setLevel(logging.DEBUG)

    cli_level = _parse_log_level(CONFIG.get('cliLogLevel'))
    cli_formatter = logging.Formatter('%(asctime)s [%(levelname)s] %(message)s')
    console = logging.StreamHandler(sys.stdout)
    console.setLevel(cli_level)
    console.setFormatter(cli_formatter)
    LOGGER.addHandler(console)

_setup_logging()
