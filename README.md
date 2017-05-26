# discord-boar-bot
Custom bot for the Sociologic Planning Boar

## Development instructions

**Requirements**

* Python 3.5+
* Opus Codec (for voice support)

**Environment setup**

> This setup is intended for a Linux system. You're on your own for a different
system.

1. Install `virtualenv`

    $ pip3 install virtualenv --upgrade

2. Set up the virtual Environment

    $ python -m virtualenv ./env

3. Using the Python binaries from the new environment,
   install the package.

    $ env/bin/pip install -e .

**Running the bot**

The binary executable for the bot resides in `env/bin/boarbot`. 
