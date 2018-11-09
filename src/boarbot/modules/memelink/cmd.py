import argparse

from boarbot.common.log import LOGGER

class MemeLinkParserException(Exception): pass

class MemeLinkParser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        LOGGER.warn('MemeLinkParser.exit called with %s %s' % (status, message))

    def error(self, message: str):
        raise MemeLinkParserException(message)

MEME_LINK_PARSER = MemeLinkParser(prog='!memes', description='Search/edit memes', add_help=False)
MEME_LINK_PARSER.add_argument('-h', '--help', action='store_true', help='print help/usage instructions')
_subparsers = MEME_LINK_PARSER.add_subparsers(help='meme sub-commands')

_search_subparser = _subparsers.add_parser('search', help='search for memes')
_search_subparser.add_argument('search', type=str, help='search query string')
