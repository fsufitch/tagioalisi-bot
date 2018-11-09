import discord
import mimetypes
import pkg_resources
import random
import re
import types
import yaml
from sqlalchemy.orm.session import Session

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.chunks import chunk_lines
from boarbot.common.log import LOGGER
from boarbot.db.memes import get_meme, search_memes

from .cmd import MEME_LINK_PARSER, MemeLinkParserException

MEME_EDIT_ACL_ID = 'boarbot.modules.memelink::edit'
QUERY_RE = re.compile(r'[a-zA-Z0-9_-]+(?:\.[a-zA-Z0-9_-]+)+')

MEMES_COMMAND = '!memes'
ERROR_FORMAT = '`{error}`\nTry `!memes --help` to get usage instructions.'


class MemeLinkModule(BotModule):
    async def handle_event(self, db_session: Session, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return

        message = args[0] # type: discord.Message
        if message.author.bot:
            return # Ignore bots

        meme_query = self.extract_query(message.clean_content)
        if meme_query:
            # someone is trying to link a meme
            LOGGER.debug(f'Meme query: {message.clean_content} -> {meme_query}')
            name = meme_query.split('.', 1)[0]
            meme = get_meme(db_session, name)
            if not meme or not meme.urls:
                return

            url = random.choice(meme.urls)

            msg = '**%s**: %s' % (meme_query, url.url)
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

        if parsed_args.search:
            await self.list_memes(db_session, message, parsed_args.search)

    def extract_query(self, message_text: str) -> str:
        match = QUERY_RE.search(message_text)
        if not match:
            return None

        return match.group(0)

    async def list_memes(self, db_session: Session, message: discord.Message, search: str):
        output_lines = []
        for meme in search_memes(db_session, search):
            memelist = ', '.join([n.name for n in meme.names])
            output_lines.append(f'- [{meme.id}]: {memelist}')

        if output_lines:
            for message_chunk in chunk_lines(output_lines):
                reply = '```' + '\n'.join(message_chunk) + '```'
                await self.client.send_message(message.author, reply)
        else:
            await self.client.send_message(message.author, '`No memes found for "%s"`' % search)

        if message.server:
            await self.client.send_message(message.channel, '%s check your direct messages for your search result. You can also query me again there to not spam the channel!' % message.author.mention)
