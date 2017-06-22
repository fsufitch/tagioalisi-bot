import asyncio
import discord

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.log import LOGGER

from .cmd import YTPLAY_PARSER, YTPlayParserException, PLAY, STATUS, STOP

YTPLAY_COMMAND = '!yt'
ERROR_FORMAT = '`{error}`\nTry `!ytplay --help` to get usage instructions.'
CHANNEL_NOT_FOUND = 'Error -- Channel not found: "{channel}"'
NOW_PLAYING = 'Playing in channel "{channel}": {url}'

class YTPlayModule(BotModule):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        if not self.load_opus():
            LOGGER.error('Opus not loaded successfully! YTPlayModule will not work.')
        self.players = {}
        LOGGER.debug('YTPlayModule initialized')

    async def handle_event(self, event_type, args):
        if event_type != EventType.MESSAGE:
            return
        message = args[0] # type: discord.Message
        if message.author.bot or not message.server:
            return # Ignore bots and private messages

        if len(message.mentions) > 1 or message.channel_mentions or message.role_mentions:
            return # Only work when I'm mentioned specifically

        args = self.parse_command(YTPLAY_COMMAND, message)
        if args is None:
            return

        if '-h' in args or '--help' in args:
            await self.client.send_message(message.channel, '```' + YTPLAY_PARSER.format_help() + '```')
            return

        try:
            parsed_args = YTPLAY_PARSER.parse_args(args)
        except YTPlayParserException as e:
            await self.client.send_message(message.channel, ERROR_FORMAT.format(error=e.args[0]))
            return

        if parsed_args.action == PLAY:
            await self.start_audio(message, parsed_args.yturl, parsed_args.channel)
        elif parsed_args.action == STOP:
            await self.stop_audio(message)
        elif parsed_args.action == STATUS:
            await self.status(message)


    async def start_audio(self, message: discord.Message, yturl: str, channel_name: str, duration=0):
        if await self._is_playing(message):
            await self.client.send_message(message.channel, "I am already playing audio! Try `!yt status` to ask what's going on.")
            return

        for channel in message.server.channels: # type: discord.Channel
            if channel.type == discord.ChannelType.voice and channel.name == channel_name:
                break
        else:
            await self.client.send_message(message.channel, CHANNEL_NOT_FOUND.format(channel=channel_name))
            return

        voice_client = await self.client.join_voice_channel(channel) # type: discord.VoiceClient
        player = await voice_client.create_ytdl_player(
            yturl, ytdl_options={'quiet': True},
            after=lambda: self.audio_done(message, voice_client),
        )

        await self.client.send_message(message.channel, 'Playing audio')
        player.start()
        self.players[voice_client.session_id] = player

    def audio_done(self, message: discord.Message, voice_client: discord.VoiceClient):
        async def _done():
            await self.client.send_message(message.channel, 'Audio done')
            await voice_client.disconnect()
        coro = _done()
        fut = asyncio.run_coroutine_threadsafe(coro, self.client.loop)
        try:
            fut.result()
        except Exception as e:
            LOGGER.exception(e)

    async def stop_audio(self, message: discord.Message):
        voice_client = message.server.voice_client # type: discord.VoiceClient
        if not voice_client:
            await self.client.send_message(message.channel, 'No audio to stop!')
            return
        await self.client.send_message(message.channel, 'Stopping audio')
        self.players[voice_client.session_id].stop()
        await voice_client.disconnect()
        del self.players[voice_client.session_id]

    async def status(self, message: discord.Message):
        if not await self._is_playing(message):
            await self.client.send_message(message.channel, 'Not playing anything')
            return
        voice_client = message.server.voice_client # type: discord.VoiceClient
        player = self.players[voice_client.session_id]
        await self.client.send_message(message.channel,
            NOW_PLAYING.format(channel=voice_client.channel.name, url=player.url)
        )

    async def _is_playing(self, message: discord.Message):
        voice_client = message.server.voice_client # type: discord.VoiceClient
        if not voice_client:
            return False
        player = self.players[voice_client.session_id]
        if not player.is_playing():
            LOGGER.warn('Bad state: voice client in channel but not playing audio')
            await self.stop_audio(message)
            return False
        return True
