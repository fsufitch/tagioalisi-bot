import discord
import mimetypes
import pkg_resources
import random
import re
import types
import yaml

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType

QUERY_RE = re.compile('[a-zA-Z0-9_-]+(?:\.[a-zA-Z0-9_-]+)+')
MEMES_YAML = pkg_resources.resource_string('boarbot.modules.memelink', 'memes.yaml').decode()
MEMES = yaml.load(MEMES_YAML)

LIST_MEMES_COMMAND = 'list memes'

class MemeLinkModule(BotModule):
    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return

        message = args[0] # type: discord.Message
        if message.author.bot:
            return # Ignore bots

        if self.client.user.mentioned_in(message) and LIST_MEMES_COMMAND in message.clean_content:
            await self.list_memes(message)
            return

        query = self.extract_query(message.clean_content)
        if not query:
            return
        url = self.get_meme(query)
        if not url:
            return

        msg = '**%s**: %s' % (query, url)
        await self.client.send_message(message.channel, msg)

    def extract_query(self, message_text: str) -> str:
        match = QUERY_RE.search(message_text)
        if not match:
            return None

        return match.group(0)

    def get_meme(self, query: str) -> str:
        query = query.lower()
        meme_name, ext = query.split('.', 1)
        mimetype = mimetypes.guess_type(query)[0] or ''

        for meme in MEMES:
            if meme.get('type') and meme['type'] not in mimetype:
                continue

            if meme_name in meme['names']:
                return random.choice(meme['urls'])

        return None

    async def list_memes(self, message: discord.Message):
        query = message.clean_content.split(LIST_MEMES_COMMAND, 1)[1].strip() # type: str
        output_lines = []
        for meme in MEMES:
            for name in meme['names']:
                if query in name:
                    break
            else: # query not in any names, skip this meme
                continue

            output = '- '
            if meme.get('type'):
                output = '- (%s) ' % meme['type']
            output += ', '.join(meme['names'])
            output_lines.append(output)

        if output_lines:
            for message_chunk in self.chunk_output_lines(lines):
                reply = '```' + '\n'.join(message_chunk) + '```'
                await self.client.send_message(message.author, reply)
        else:
            await self.client.send_message(message.author, '`No memes found for "%s"`' % query)

    def chunk_output_lines(self, lines: [str], max_chars=1900) -> [[str]]:
        chunks = []
        current_chunk = []
        current_chunk_len = 0
        for line in lines:
            if current_chunk_len + len(line) > max_chars:
                chunks.append(current_chunk)
                current_chunk = []
                current_chunk_len = 0
            current_chunk.append(line)
            current_chunk_len += len(line) + 1 # account for '\n'
        chunks.append(current_chunk)
        return chunks
