# discord-boar-bot
Custom bot for the Sociologic Planning Boar

## Development instructions

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
