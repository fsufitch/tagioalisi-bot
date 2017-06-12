import argparse
import discord

from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType

from .cmd import DICE_PARSER, DiceParserException
from .diceroll import DiceRoll

DICE_COMMAND = '!roll'
ERROR_FORMAT = '`{error}`\nTry `!roll --help` to get usage instructions.'

class DiceRollModule(BotModule):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    async def handle_event(self, event_type: EventType, args):
        if event_type != EventType.MESSAGE:
            return
        message = args[0]

        args = self.parse_command(DICE_COMMAND)
        if args is None:
            return

        try:
            parsed_args = DICE_PARSER.parse_args(args)
        except DiceParserException as e:
            await self.send_message(message.channel, ERROR_FORMAT.format(error=e.args[0]))
            return

        if parsed_args.help:
            await self.send_message(message.channel, '```' + DICE_PARSER.format_help() + '```')
            return

        if not parsed_args.dice:
            await self.send_message(message.channel, ERROR_FORMAT.format(error='error: no dice specified'))
            return

        try:
            roll = DiceRoll(' '.join(parsed_args.dice))
        except Exception as e:
            await self.send_message(message.channel, ERROR_FORMAT.format(error=str(e)))

        if parsed_args.verbose:
            msg = self.reply_verbose(message.channel, roll)
        else:
            msg = self.reply(message.channel.roll)

        await self.send_message(message.channel, msg)

    def reply(self, roll: DiceRoll) -> str:
        return '`{}` => `{}`'.format(roll.rolldef, roll.roll[0])

    def reply_verbose(self, roll: DiceRoll):
        msg = self.reply(roll) + '\n**Details:**\n'
        total, min_roll, max_roll, roll_results = roll.roll
        for row in roll_results:
            expr = row['roll_exp']
            dice = [str(val) for val in row['rolls']]
            value = row['value']
            if dice:
                line = '`{}` => `[{}]` => `{}`\n'.format(expr, ', '.join(dice), value)
            else:
                line = '`{}` => `{}`\n'.format(expr, value)
            msg += line
        return msg
