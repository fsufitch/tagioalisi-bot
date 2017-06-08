import logging

from boarbot.common.botmodule import BotModule
from boarbot.common.config import CONFIG
from boarbot.common.events import EventType
from boarbot.common.log import LOGGER, register_discord_handler

LOG_CHANNEL = CONFIG.get('discordLogChannel')

class BoarLogger(BotModule):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        register_discord_handler(self.emit_log)
        LOGGER.debug('Boar logger started and registered with main logger')

    async def handle_event(self, event_type, args):
        if event_type == EventType.READY:
            LOGGER.info('Discord bot started')

    async def emit_log(self, message: str):
        channel = self.client.get_channel(LOG_CHANNEL)
        if not channel:
            return
        await self.client.send_message(channel, message)
