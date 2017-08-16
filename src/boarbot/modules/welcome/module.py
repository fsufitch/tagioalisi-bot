import argparse
import discord

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.config import CONFIG

WELCOME_CHANNEL = CONFIG.get('welcome', {}).get('channel')
SERVER_RULES = CONFIG.get('welcome', {}).get('rules')

WELCOME_MESSAGE = '''Welcome to the Discord server, {mention}!

I am the Boar Bot, and I work very hard to make everyone's life more fun.
You may want to check out the server guide here: {rules}

Have fun!
'''

class WelcomeModule(BotModule):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MEMBER_JOIN:
            return

        member = args[0] # type: discord.Member
        channel = self.client.get_channel(LOG_CHANNEL)
        if not channel:
            return

        message_text = WELCOME_MESSAGE.format(mention=member.mention, rules=SERVER_RULES)
        await self.client.send_message(channel, message_text)
