import discord
from sqlalchemy.orm.session import Session

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType

class EchoModule(BotModule):
    async def handle_event(self, db_session: Session, event_type: EventType, args):
        if event_type == EventType.MESSAGE:
            await self.echo(args[0])

    async def echo(self, message: discord.Message):
        if not self.client.user.mentioned_in(message):
            return # Gotta mention me

        if '!echo' not in message.clean_content:
            return # Need !echo

        echo = message.clean_content.split('!echo', 1)[1]
        await self.client.send_message(message.channel, echo)
