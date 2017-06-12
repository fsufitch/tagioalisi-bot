import argparse

from boarbot.common.log import LOGGER

class DiceParserException(Exception): pass

class DiceParser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        LOGGER.warn('DiceParser.exit called with %s %s' % (status, message))

    def error(self, message: str):
        raise DiceParserException(message)

DICE_PARSER = DiceParser(prog='!roll', description='Roll some dice', add_help=False)
DICE_PARSER.add_argument('dice', metavar='DICE', type=str, nargs='*',
                     help='the dice to roll in XdY format (e.g. "2d8 + 1d6 + 4")')
DICE_PARSER.add_argument('-v', '--verbose', action='store_true', help='print individual rolls')
DICE_PARSER.add_argument('-h', '--help', action='store_true', help='print help/usage instructions')
