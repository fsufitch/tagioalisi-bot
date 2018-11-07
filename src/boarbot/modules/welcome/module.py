import argparse
import discord

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.config import WELCOME
from boarbot.common.log import LOGGER

WELCOME_CHANNEL = WELCOME['channel']
RULES_CHANNEL = WELCOME['rulesChannel']

WELCOME_MESSAGE = '''Welcome to the Discord server, {mention}!

I am the Boar Bot, and I work very hard to make everyone's life more fun.
You may want to check out the server rules here: {rulesChannel}

Have fun!
'''

class WelcomeModule(BotModule):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MEMBER_JOIN:
            return

        member = args[0] # type: discord.Member
        channel = self.client.get_channel(WELCOME_CHANNEL)
        if not channel:
            log.error("Could not find welcome channel")
            return

        rules_channel = self.client.get_channel(RULES_CHANNEL)
        if not rules_channel:
            log.error("Could not find rules channel")
            return

        message_text = WELCOME_MESSAGE.format(mention=member.mention, rulesChannel=rules_channel.mention)
        await self.client.send_message(channel, message_text)
