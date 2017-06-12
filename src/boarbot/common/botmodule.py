import discord
import shlex
from abc import ABCMeta, abstractmethod

from boarbot.common.events import EventType
from boarbot.common.log import LOGGER

class BotModule(metaclass=ABCMeta):
    def __init__(self, client: discord.Client):
        self.client = client

    @abstractmethod
    async def handle_event(self, event_type: EventType, args):
        ...

    '''
    Parse a command in the format `@<bot> <command> <command arguments>`.
    Returns either a list of string arguments, or None if this is not a valid
    call to the given command.
    '''
    def parse_command(self, command: str, message: discord.Message, ignore_bots=True) -> [str]:
        if ignore_bots and message.author.bot:
            return None

        content = message.content.strip() # type: str
        mention = self.client.user.mention # type: str
        LOGGER.debug('Parsing content `{}` for command `{}` and user {}'.format(message.content.strip(), command, mention))
        if not content.startswith(mention):
            return None

        content = content[len(mention):].strip() # type: str
        parts = shlex.split(content) # type: [str]
        if not parts or parts[0] != command:
            return None

        return parts[1:]
