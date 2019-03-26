import argparse
import asyncio
import concurrent.futures as futures
import discord
import pkg_resources
import random
import shlex
import yaml
from sqlalchemy.orm.session import Session


from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.log import LOGGER

from .cmd import NOWPLAYING_COMMAND, NOWPLAYING_PARSER, NowPlayingParserException

NOWPLAYING_YAML = pkg_resources.resource_string('boarbot.modules.nowplaying', 'nowplaying.yaml').decode()
NOWPLAYING_MESSAGES = yaml.load(NOWPLAYING_YAML)['messages'] # type: [str]
ERROR_FORMAT = '`{error}`\nTry `!play --help` to get usage instructions.'
NOWPLAYING_TIMEOUT = 60.0

class NowPlayingModule(BotModule):
    def __init__(self, client: discord.Client):
        super().__init__(client)
        self.status_queue = asyncio.Queue(loop=self.client.loop)
        asyncio.ensure_future(self.handle_status_updates(), loop=self.client.loop)
        
    async def handle_event(self, db_session: Session, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return

        message = args[0]  # type: discord.Message
        if message.author.bot:
            return  # Ignore bots
        
        args = self.parse_command(NOWPLAYING_COMMAND, message)
        if args is None:
            return
        try:
            parsed_args = NOWPLAYING_PARSER.parse_args(args)
        except NowPlayingParserException as e:
            await self.client.send_message(message.channel, ERROR_FORMAT.format(error=e.args[0]))
            return

        playmessage = ' '.join(parsed_args.playmsg or [])
        if not playmessage:
            await self.client.send_message(message.channel, "You should probably tell me to actually play _something_.")
            return
        if len(playmessage) > 50:
            await self.client.send_message(message.channel, "That's too long of a thing to play.")
            return

        LOGGER.info("Now playing `%s` thanks to user `%s`", playmessage, message.author.display_name)
        await self.client.send_message(message.channel, "Now playing `{}`.".format(playmessage))
        await self.status_queue.put(playmessage)
        
    async def handle_status_updates(self):
        LOGGER.debug('Starting status update loop')
        await asyncio.sleep(5, loop=self.client.loop)
        await self.set_status(random.choice(NOWPLAYING_MESSAGES))
        while True:
            try:
                status = await asyncio.wait_for(self.status_queue.get(), NOWPLAYING_TIMEOUT, loop=self.client.loop)
                LOGGER.debug('Status loop received: ' + status)
            except asyncio.TimeoutError:
                status = random.choice(NOWPLAYING_MESSAGES)
                LOGGER.debug('Status loop timed out: ' + status)

            await self.set_status(status)

    async def set_status(self, status: str):
        await self.client.change_presence(game=discord.Game(name=status))
