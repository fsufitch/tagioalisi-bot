import argparse

from boarbot.common.log import LOGGER

class YTPlayParserException(Exception): pass

class YTPlayParser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        LOGGER.warn('YTPlayParser.exit called with %s %s' % (status, message))

    def error(self, message: str):
        raise YTPlayParserException(message)

PLAY='play'
STOP='stop'
STATUS='status'

YTPLAY_PARSER = YTPlayParser(prog='!ytplay', description='Play YouTube audio over a Discord channel', add_help=False)
YTPLAY_PARSER.add_argument('action', choices=[PLAY, STATUS, STOP])
YTPLAY_PARSER.add_argument('channel', metavar='CHANNEL', type=str, nargs='?', help='channel to play audio to (use quotes for multi-word channel)')
YTPLAY_PARSER.add_argument('yturl', metavar='YOUTUBE_URL', type=str, nargs='?')
YTPLAY_PARSER.add_argument('-d', '--duration', type=int, help='number of seconds to play')
YTPLAY_PARSER.add_argument('-h', '--help', action='store_true', help='print help/usage instructions')
