import discord
import mimetypes
import pkg_resources
import random
import re
import types
import yaml

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.chunks import chunk_lines

from .cmd import MEME_LINK_PARSER, MemeLinkParserException

QUERY_RE = re.compile(r'[a-zA-Z0-9_-]+(?:\.[a-zA-Z0-9_-]+)+')
MEMES_YAML = pkg_resources.resource_string('boarbot.modules.memelink', 'memes.yaml').decode()
MEMES = yaml.load(MEMES_YAML)

MEMES_COMMAND = '!memes'
ERROR_FORMAT = '`{error}`\nTry `!memes --help` to get usage instructions.'


class MemeLinkModule(BotModule):
    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return

        message = args[0] # type: discord.Message
        if message.author.bot:
            return # Ignore bots

        meme_query = self.extract_query(message.clean_content)
        if meme_query:
            # someone is trying to link a meme
            url = self.get_meme(meme_query)
            if not url:
                return

            msg = '**%s**: %s' % (meme_query, url)
            await self.client.send_message(message.channel, msg)
            return

        # otherwise, someone might be trying to use !memes
        args = self.parse_command(MEMES_COMMAND, message)
        if args is None:
            return
        try:
            parsed_args = MEME_LINK_PARSER.parse_args(args)
        except MemeLinkParserException as e:
            await self.client.send_message(message.channel, ERROR_FORMAT.format(error=e.args[0]))
            return

        if parsed_args.help:
            await self.client.send_message(message.channel, '```' + MEME_LINK_PARSER.format_help() + '```')
            return

        search = parsed_args.search

        await self.list_memes(message, search)

    def extract_query(self, message_text: str) -> str:
        match = QUERY_RE.search(message_text)
        if not match:
            return None

        return match.group(0)

    def get_meme(self, query: str) -> str:
        query = query.lower()
        meme_name, _ = query.split('.', 1)
        mimetype = mimetypes.guess_type(query)[0] or ''

        for meme in MEMES:
            if meme.get('type') and meme['type'] not in mimetype:
                continue

            if meme_name in meme['names']:
                return random.choice(meme['urls'])

        return None

    async def list_memes(self, message: discord.Message, search: str):
        output_lines = []
        for meme in MEMES:
            for name in meme['names']:
                if search in name:
                    break
            else: # query not in any names, skip this meme
                continue

            output = '- '
            if meme.get('type'):
                output = '- (%s) ' % meme['type']
            output += ', '.join(meme['names'])
            output_lines.append(output)

        if output_lines:
            for message_chunk in chunk_lines(output_lines):
                reply = '```' + '\n'.join(message_chunk) + '```'
                await self.client.send_message(message.author, reply)
        else:
            await self.client.send_message(message.author, '`No memes found for "%s"`' % search)

        if message.server:
            await self.client.send_message(message.channel, '%s check your direct messages for your search result. You can also query me again there to not spam the channel!' % message.author.mention)
