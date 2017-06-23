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
        #LOGGER.debug('Parsing content `{}` for command `{}` and user {}'.format(message.content.strip(), command, mention))
        if not content.startswith(mention):
            return None

        content = content[len(mention):].strip() # type: str

        try:
            parts = shlex.split(content) # type: [str]
        except Exception as e:
            LOGGER.debug('Failed shlex.split on ' + str(content))
            return None

        if not parts or parts[0] != command:
            return None

        return parts[1:]

    def load_opus(self):
        if not discord.opus.is_loaded():
            from ctypes.util import find_library
            LOGGER.debug('libopus not automatically loaded')
            opus_attempts = [
                find_library('opus'),
                'libopus.so.0',
                'opus.dll',
                'opus',
            ]
            for opus in opus_attempts:
                if not opus:
                    continue
                try:
                    LOGGER.debug('Loading Opus codec from ' + opus)
                    discord.opus.load_opus(opus)
                    LOGGER.debug('Success!')
                    break

                except OSError as e:
                    LOGGER.debug('Failed! ' + str(e))
            else:
                LOGGER.error('Could not load Opus codec! Voice support is not available.')
        return discord.opus.is_loaded()
