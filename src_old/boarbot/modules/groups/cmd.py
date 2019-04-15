import argparse

from boarbot.common.log import LOGGER

class GroupsParserException(Exception): pass

class GroupsParser(argparse.ArgumentParser):
    def exit(self, status=0, message=None):
        LOGGER.warn('GroupsParser.exit called with %s %s' % (status, message))

    def error(self, message: str):
        raise GroupsParserException(message)

GROUPS_PARSER = GroupsParser(prog='!groups', description='Manage server "ping" groups', add_help=False)
GROUPS_PARSER.add_argument('command', metavar='CMD', type=str, nargs='?', default='<no command>', help='group command (list | members | create | delete | join | leave)')
GROUPS_PARSER.add_argument('group', metavar='GROUP', type=str, nargs='?', default=None, help='the group name to operate on')
GROUPS_PARSER.add_argument('users', metavar='USER', type=str, nargs='*', help='@mentions of the target users (available only for managers)')
GROUPS_PARSER.add_argument('-h', '--help', action='store_true', help='print help/usage instructions')
