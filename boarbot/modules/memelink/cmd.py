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
_subparsers = MEME_LINK_PARSER.add_subparsers(dest='subcommand', help='meme subcommand')

SEARCH_SUBPARSER = _subparsers.add_parser('search', help='search for memes')
SEARCH_SUBPARSER.add_argument('search', type=str, help='search query string')

ADD_MEME_SUBPARSER = _subparsers.add_parser('add', help='add a new meme')
ADD_MEME_SUBPARSER.add_argument('add_name', metavar='name', type=str, help='a short name for the meme, with no spaces')
ADD_MEME_SUBPARSER.add_argument('add_url', metavar='url', type=str, help='URL the meme should display')

ADD_MEME_SUBPARSER = _subparsers.add_parser('alias', help='add a new name alias for a meme')
ADD_MEME_SUBPARSER.add_argument('alias_new', metavar='alias', type=str, help='new alias for the meme')
ADD_MEME_SUBPARSER.add_argument('alias_name', metavar='name', type=str, help='the name of an existing meme')
