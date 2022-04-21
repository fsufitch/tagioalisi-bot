# Tagioalisi

> Formerly _discord-boar-bot_

Custom bot for the Sociologic Planning Boar server on Discord. Now in Go!

                 ,           _______________
            `-.   \    .-'  < Hello, world! >
    ,-"`````""-\__ |  /      ---------------
     '-.._    _.-'` '-o,   /
         _>--:{{<   ) |)  /
     .-''      '-.__.-o`
    '-._____..-/`  |  \
         ,-'   /    `-.

## Requirements

* Go 1.14+
* Docker and Docker Compose (optional)

## Build and run instructions

The bot provides two executable entry points, which can be built with these commands:

    go build ./cmd/tagi-bot
    go build ./cmd/tagi-migrate
    
The executables produced by these packages are ready to be run anywhere. They are configured
using environment variables. An example configuration for local development can be found
in [.env.template](./.env.template).

Remember that the bot depends on connecting to Discord itself, and there is no "fake" local
Discord server available. Also take not that both the bot and its migrations rely on a
Postgres database. To simplify setting that up, see the Docker instructions below.

## Using Docker

The bot comes with a `Dockerfile` and `docker-compose.yml` to ease setup. To set up a full
running bot, follow these steps.

1. Copy `.env.template` into `.env` and edit it so `DISCORD_TOKEN` contains a valid Discord bot token 
2. Run `docker-compose run bot ./tagi-migrate` to set up the database
3. Run `docker up` (with `--build` after every time you make code changes)

## Configuration

The runtime is configured exclusively through environment variables. Note the below default behaviors
do differ from the values in `.env.template`.

| Variable | Default (unset) | Purpose |
| --- | --- | --- |
| `DEBUG` | false | toggles detailed debug logging; also toggles info-level logging by the bot in a discord channel |
| `DATABASE_URL` | **ERROR** | [Postgres connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) for accessing the database |
| `DISCORD_TOKEN` | **ERROR** | (bot only) used for authenticating the bot to Discord |
| `BLACKLIST_BOT_MODULES` | (bot only) comma-separated bot modules to not load; see below for details |
| `DISCORD_LOG_CHANNEL` | | (bot only) ID of channel the bot should use for logging; no channel logs if empty |
| `GROUP_PREFIX` | `g-` | (bot only) the prefix to use for any bot-managed roles, as part of the `groups` module |
| `AZURE_NEWS_SEARCH_KEY` | **ERROR** | (bot only) the API key to use to contact Azure News Search API |
| `OAUTH_*` | **ERROR** | (bot only) configuration for allowing users to identify themselves using the [Discord OAuth2 API](https://discord.com/developers/docs/topics/oauth2) |
| `JWT_HMAC_SECRET` | `supersecret` | (bot only) the secret salt used for signing JWTs; *do not use the default value* |
| `AES_KEY_B64` | `zvFXk6LbvG5jTG7JlGBQ57MdqUwjdS/wH3gVWXkSOUU=` | (bot only) base-64 encoded random 256-bit AES string to use as an AES encryption key; *do not use the default value* | 
| `WEB_ENABLED` | false | (bot only) enable the HTTP web API for runtime control of the bot |
| `PORT` | 80 | (bot only) the TCP port that the bot HTTP web API listens on; when running inside Docker, this should always be 80 |
| `MIGRATION_DIR` | **ERROR** | (migration only) the directory where migration SQL files can be found; see [`./db/migration`](./db/migration) |

## Modular design

The bot functionality is split up into mostly independent modules that each act as a controller for a smaller set of behavior. They are:

| Name | Implemented/Ported | Purpose |
| --- | --- | --- |
| `log` | ✅ | writes logs of warning+ level into a channel; with `DEBUG` set, also writes info-level logs |
| `memelink` | ✅ | registers and responds with meme link content in response to certain "filenames" (e.g. `facepalm.jpg`)  |
| `ping` | ✅ | responds to `!ping` with `!pong` for sanity checking |
| `sockpuppet` | ✅ | hooks allowing for custom messages to be sent via the web UI |
| `groups` | ✅ | manage a server groups system using special prefixed roles |
| `wiki` | ✅ | search a variety of wikis |
| `dice` | ✅ | roll dice (e.g. `!roll 1d20+3`) |
| `news` | ✅ | use the Bing News Search API to look for recent news about some keywords |
| `welcome` | ⬜️ | welcomes new users and points them useful places|
| `remindme` | ⬜️ | basic reminder system for reminding yourself and others of stuff |
| `ytplay` | ⬜️ | pipe audio from a YouTube video into an audio channel; as annoying as possible |

These can be individually turned on/off using the `BLACKLIST_BOT_MODULES` environment variable. 
