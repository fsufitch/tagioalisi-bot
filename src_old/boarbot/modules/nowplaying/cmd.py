import argparse

from boarbot.common.log import LOGGER

NOWPLAYING_COMMAND = '!play'

class NowPlayingParserException(Exception):
    pass


class NowPlayingParser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        LOGGER.warn('NowPlayingParser.exit called with %s %s' %
                    (status, message))

    def error(self, message: str):
        raise NowPlayingParserException(message)


NOWPLAYING_PARSER = NowPlayingParser(prog=NOWPLAYING_COMMAND, description='"Play" something', add_help=False)
NOWPLAYING_PARSER.add_argument('playmsg', metavar='MESSAGE', type=str, nargs='+', default='', help='thing to "play"')
NOWPLAYING_PARSER.add_argument('-h', '--help', action='store_true', help='print help/usage instructions')
