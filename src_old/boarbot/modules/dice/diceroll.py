import random, regex

def roll_dice(num_dice: int, num_sides: int) -> [int]:
    num_dice = int(num_dice)
    num_sides = int(num_sides)
    if num_dice < 1:
        raise ValueError("Number of dice less than 0", num_dice)
    if num_sides < 1:
        raise ValueError("Number of sides less than 0", num_sides)
    rolls = [random.randint(1, num_sides) for i in range(num_dice)]
    return rolls


DICE_REGEX = regex.compile("\A([0-9]+d[0-9]+)([+-][0-9]+(?:d[0-9]+)?)*\Z")

class DiceRoll(object):
    def __init__(self, rolldef: str, comment=""):
        #Clear whitespace
        self.rolldef_dirty = rolldef
        self.rolldef = regex.split('[ \r\n\t]', rolldef)
        self.rolldef = ''.join(self.rolldef).lower()

        self.comment = comment

        self.rolls = []
        self.roll_results = []
        self.roll_total = 0
        self._parse()
        self.do_roll()

    def _parse(self):
        match = DICE_REGEX.search(self.rolldef)
        if not match:
            raise ValueError("Invalid dice", self.rolldef)
        firstroll = '+' + match.captures(1)[0]
        self.rolls = [firstroll] + match.captures(2)

    @property
    def roll(self) -> (int, int, int, int):
        if not self.roll_results:
            self.do_roll()
        return self.roll_total, self.roll_min, self.roll_max, self.roll_results

    def do_roll(self, reroll=False):
        if not reroll and self.roll_results:
            return
        self.roll_results = []

        for roll_exp in self.rolls:
            roll_entry = {}
            roll_entry['roll_exp'] = roll_exp
            roll_sign = roll_exp[0]
            if "d" in roll_exp:
                num_dice, num_sides = roll_exp[1:].split("d")
                roll_entry['num_dice'] = int(num_dice)
                roll_entry['num_sides'] = int(num_sides)
                dice_rolls = roll_dice(num_dice, num_sides)
                roll_entry['rolls'] = dice_rolls
                roll_entry['min'] = roll_entry['num_dice']
                roll_entry['max'] = roll_entry['num_dice'] * roll_entry['num_sides']
            else:
                # Constant value
                roll_entry['num_dice'] = 0
                roll_entry['num_sides'] = 0
                roll_entry['rolls'] = []
                dice_rolls = [int(roll_exp[1:])]
                roll_entry['min'] = dice_rolls[0]
                roll_entry['max'] = dice_rolls[0]
            roll_entry['value'] = sum(dice_rolls) * {'+':1,'-':-1}[roll_sign]
            self.roll_results.append(roll_entry)

        self.roll_total = sum([entry['value'] for entry in self.roll_results])
        self.roll_min = sum([entry['min'] for entry in self.roll_results])
        self.roll_max = sum([entry['max'] for entry in self.roll_results])
