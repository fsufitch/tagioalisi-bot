# discord-boar-bot
Custom bot for the Sociologic Planning Boar

## Development setup instructions

**Requirements**

* Python 3.5+
* Opus Codec (for voice support)

**Environment setup**

> This setup is intended for a Linux system. You're on your own for a different
system.

1\. Install `virtualenv`

    $ pip3 install virtualenv --upgrade

2\. Set up the virtual Environment

    $ python -m virtualenv ./env

3\. Using the Python binaries from the new environment,
   install the package.

    $ env/bin/pip install -e .

**Running the bot**

The binary executable for the bot resides in `env/bin/boarbot`. Running it
requires a Discord bot token. To deploy your own version of this bot, follow
these steps:

- Create your own application for it here https://discordapp.com/developers/applications/me

- Create a bot user for the application

- Run the bot using `env/bin/boarbot <DISCORD_BOT_TOKEN>`

- Insert your bot's client ID (from the application page above) into this URL
  and visit it in order to add the bot to your server(s):

  `https://discordapp.com/api/oauth2/authorize?client_id={{APPLICATION_CLIENT_ID}}&scope=bot&permissions=0`

**Command line options**

    usage: boarbot [-h] [-m import.path:ClassName] TOKEN

    positional arguments:
    TOKEN                 Discord API token identifying your bot

    optional arguments:
    -h, --help            show this help message and exit
    -m import.path:ClassName, --module import.path:ClassName
                        import paths and class names of modules to load; can
                        be specified multiple times

## Extending the bot

### Installing new modules

**Default modules:** The default modules (e.g. logging module) that the bot
loads are configured as part of the (config.json)[src/boarbot/config.json] file.
To change these, you must edit the file and restart the bot (and reinstall, if you
installed it without the `-e` option).

**Runtime modules:** To dynamically include a module when the bot starts,
you can use the `-m` or `--module` command line option to specify an import
path and class name. This can even be an external library, such as
`fsufitch.mybots.awesome:MyAwesomeModule`. As long as the description can be
imported as a Python class, it should work.

> Modules should never import `boarbot.cli`, as that causes an import loop.

### Writing new modules

Module classes _must_ extend the `boarbot.common.BotModule` abstract class.
In doing so, they must also implement the `handle_event` [subroutine](https://docs.python.org/3.5/library/asyncio.html#module-asyncio), and may
wish to implement initialization logic in the constructor. You can use the
[boarbot.modules.echo:EchoModule](src/boarbot/modules/echo.py) as a sample, or
the skeleton below:

```python
from boarbot.common.botmodule import BotModule
from boarbot.common.events import EventType
from boarbot.common.log import LOGGER

class MyAwesomeModule(BotModule):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.foo = "Look ma, I'm initializing things!"

    async def handle_event(self, event_type: EventType, args: list):
        if event_type == EventType.READY:
            LOGGER.info('MyAwesomeModule is ready!')
```

The `BotModule` class makes the Discord client available as `self.client`.
For documentation on what it can do and how it should be used, see its docs
[here](http://discordpy.readthedocs.io/en/latest/api.html#client).

The signature of the `handle_event` method includes an `EventType` value and
a list of arguments for that event (which varies based on the event). The
currently supported list of events and arguments are:

| Event Type  | Arguments | Description |
| ------------- | ------------- | --- |
| `READY` | None | The client is done preparing the data received from Discord. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_ready) |
| `RESUMED` | None | The client has resumed a session. |
| `MESSAGE` | `[ discord.Message ]` | A new message has been created. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_message) |
| `MESSAGE_DELETE` | `[ discord.Message ]` | A message has been deleted. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_message_delete) |
| `MESSAGE_EDIT` | `[ discord.Message, discord.Message ]` | A message has been edited. The "before" and "after" states are supplied as arguments. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_message_edit) |
| `REACTION_ADD` | `[ discord.Reaction, discord.User ]` | A user added a reaction. To get the message being reacted, access it via `Reaction.message`. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_reaction_add) |
| `REACTION_REMOVE` | `[ discord.Reaction, discord.User ]` | A user removed a reaction. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_reaction_remove) |
| `REACTION_CLEAR` | `[ discord.Message, [ discord.Reaction, ... ] ]` | A message was stripped of all its reactions. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_reaction_clear) |
| `MEMBER_JOIN` | `[ discord.Member ]` | A member joined a server. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_member_join) |
| `MEMBER_REMOVE` | `[ discord.Member ]` | A member left a server. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_member_remove) |
| `MEMBER_UPDATE` | `[ discord.Member, discord.Member ]` | A member's profile changed (status, game playing, avatar, nickname, roles). The "before" and "after" states are supplied as arguments. [More...](http://discordpy.readthedocs.io/en/latest/api.html#discord.on_member_update) |

### Library tools

In addition to the entire [discord.py](http://discordpy.readthedocs.io/en/latest/)
library, the following tools are available:

* `boarbot.common.config.CONFIG`: a `dict` containing settings from the `config.json`
  file, allowing for more easily having global configurations.
* `boarbot.common.log.LOGGER`: a [`logging.Logger`](https://docs.python.org/3/library/logging.html#logger-objects)
  that will log things according to the configuration in `config.json`. It
  is also hooked in to `boarbot.modules.logger.BoarLogger` module.
