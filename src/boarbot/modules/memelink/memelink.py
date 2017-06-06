import discord
import mimetypes
import pkg_resources
import random
import re
import yaml

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType

QUERY_RE = re.compile('[a-zA-Z0-9_-]+(?:\.[a-zA-Z0-9_-]+)+')
MEMES_YAML = pkg_resources.resource_string('boarbot.modules.memelink', 'memes.yaml').decode()
MEMES = yaml.load(MEMES_YAML)

class MemeLinkModule(BotModule):
    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return

        message = args[0] # type: discord.Message
        if message.author.bot:
            return # Ignore bots

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
        mimetype = mimetypes.guess_type(query)[0]

        for meme in MEMES:
            if meme.get('type') and meme['type'] not in mimetype:
                continue

            if meme_name in meme['names']:
                return random.choice(meme['urls'])

        return None
