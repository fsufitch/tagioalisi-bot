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
from boarbot.db.acl import check_acl_roles, check_acl_user
from boarbot.db.memes import get_meme, search_memes, new_meme, MemeAlreadyExistsException

from .cmd import MEME_LINK_PARSER, SEARCH_SUBPARSER, ADD_MEME_SUBPARSER, MemeLinkParserException

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
        if meme_query and not message.clean_content.strip().lower().startswith('!memes'):
            # someone is trying to link a meme
            LOGGER.debug(f'Meme query: {message.clean_content} -> {meme_query}')
            name = meme_query.split('.', 1)[0]
            meme = get_meme(db_session, name)
            if meme and meme.urls:
                url = random.choice(meme.urls)
                msg = '**%s**: %s' % (meme_query, url.url)
                await self.client.send_message(message.channel, msg)

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

        if parsed_args.subcommand == 'search':
            if parsed_args.search:
                await self.list_memes(db_session, message, parsed_args.search)
            else:
                await self.client.send_message(message.channel, '```' + SEARCH_SUBPARSER.format_help() + '```')

        if parsed_args.subcommand == 'add':
            if parsed_args.add_name and parsed_args.add_url:
                await self.add_meme(db_session, message, parsed_args.add_name, parsed_args.add_url)
            else:
                await self.client.send_message(message.channel, '```' + SEARCH_SUBPARSER.format_help() + '```')

    def extract_query(self, message_text: str) -> str:
        match = QUERY_RE.search(message_text)
        if not match:
            return None

        return match.group(0)

    def _is_meme_editor(self, db_session: Session, member: discord.Member) -> bool:
        return any([
            check_acl_user(db_session, MEME_EDIT_ACL_ID, member.id),
            check_acl_roles(db_session, MEME_EDIT_ACL_ID, [r.id for r in member.roles]),
        ])

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

    async def add_meme(self, db_session, message: discord.Message, name: str, url: str):
        if not self._is_meme_editor(db_session, message.author):
            await self.client.send_message(message.channel, f'You are not a meme editor and may not do that.')
            return

        author = f'{message.author.name}#{message.author.discriminator}'
        try:
            new_meme(db_session, name, url, author)
        except MemeAlreadyExistsException:
            await self.client.send_message(message.channel, f'A meme using the name `{name}` already exists!')
            return
        
        db_session.commit()
        await self.client.send_message(message.channel, f'Meme `{name}` -> `{url}` added.')
