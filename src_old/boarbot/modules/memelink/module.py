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
from boarbot.db.memes import get_meme, search_memes, new_meme, MemeAlreadyExistsException, add_url, add_alias

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
                await self.client.send_message(message.channel, '```' + ADD_MEME_SUBPARSER.format_help() + '```')

        if parsed_args.subcommand == 'alias':
            if parsed_args.alias_name and parsed_args.alias_new:
                await self.add_alias(db_session, message, parsed_args.alias_name, parsed_args.alias_new)
            else:
                await self.client.send_message(message.channel, '```' + ADD_MEME_SUBPARSER.format_help() + '```')


    def extract_query(self, message_text: str) -> str:
        match = QUERY_RE.search(message_text)
        if not match:
            return None

        return match.group(0)

    def _is_meme_editor(self, db_session: Session, user: discord.User) -> bool:
        return (
            check_acl_user(db_session, MEME_EDIT_ACL_ID, user.id) or
            (
                type(user) is discord.Member and 
                check_acl_roles(db_session, MEME_EDIT_ACL_ID, [r.id for r in user.roles])
            )
        )

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

    async def add_meme(self, db_session: Session, message: discord.Message, name: str, url: str):
        if not self._is_meme_editor(db_session, message.author):
            await self.client.send_message(message.channel, f'You are not a meme editor and may not do that.')
            return

        if not name.isalnum():
            await self.client.send_message(message.channel, 'The name must be a single alphanumeric series of characters.')
            return
        
        author = f'{message.author.name}#{message.author.discriminator}'
        try:
            new_meme(db_session, name, url, author)
            await self.client.send_message(message.channel, f'Meme `{name}` -> `{url}` added.')
        except MemeAlreadyExistsException:
            add_url(db_session, name, url, author)
            meme = get_meme(db_session, name)
            also_affected = ', '.join([n.name for n in meme.names if n.name != name.lower()])
            reply = f'Added new URL to existing meme `{name}`.'
            if also_affected:
                reply += f' Also affected: `{also_affected}`'
            await self.client.send_message(message.channel, reply)
        
        db_session.commit()

    async def add_alias(self, db_session: Session, message: discord.Message, name: str, alias: str):
        if not self._is_meme_editor(db_session, message.author):
            await self.client.send_message(message.channel, f'You are not a meme editor and may not do that.')
            return

        if not name.isalnum() or not alias.isalnum():
            await self.client.send_message(message.channel, 'The name must be a single alphanumeric series of characters.')
            return

        if get_meme(db_session, alias):
            await self.client.send_message(message.channel, 'A meme already exists with that alias!')
            return

        alias = alias.lower()
        author = f'{message.author.name}#{message.author.discriminator}'
        try:
            add_alias(db_session, name, alias, author)
            await self.client.send_message(message.channel, f'Alias `{alias}` added to `{name}`.')
        except KeyError:
            await self.client.send_message(message.channel, 'A meme does not exist with that name in the first place.')
            return
        except Exception as e:
            await self.client.send_message(message.channel, f'Exception adding alias: {e}')
            return

        db_session.commit()
