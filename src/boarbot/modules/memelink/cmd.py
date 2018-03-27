import argparse

from boarbot.common.log import LOGGER

class MemeLinkParserException(Exception): pass

class MemeLinkParser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        LOGGER.warn('MemeLinkParser.exit called with %s %s' % (status, message))

    def error(self, message: str):
        raise MemeLinkParserException(message)

MEME_LINK_PARSER = MemeLinkParser(prog='!memes', description='List available meme image shortcuts', add_help=False)
MEME_LINK_PARSER.add_argument('search', metavar='SEARCH', type=str, nargs='?', default='', help='search string')
MEME_LINK_PARSER.add_argument('-h', '--help', action='store_true', help='print help/usage instructions')
