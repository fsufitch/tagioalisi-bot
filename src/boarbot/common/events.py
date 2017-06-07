from enum import Enum

class EventType(Enum):
    READY = 1
    RESUMED = 2
    MESSAGE = 3
    MESSAGE_DELETE = 4
    MESSAGE_EDIT = 5
    REACTION_ADD = 6
    REACTION_REMOVE = 7
    REACTION_CLEAR = 8
    MEMBER_JOIN = 9
    MEMBER_REMOVE = 10
    MEMBER_UPDATE = 11
    # Add events as necessary from here:
    # http://discordpy.readthedocs.io/en/latest/api.html#event-reference
