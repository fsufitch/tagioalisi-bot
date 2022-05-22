# Tagioalisi

[![DockerHub CI](https://github.com/fsufitch/tagioalisi-bot/actions/workflows/docker-image.yml/badge.svg)](https://github.com/fsufitch/tagioalisi-bot/actions/workflows/docker-image.yml)
[![Generic badge](https://img.shields.io/badge/DockerHub-latest-blue.svg)](https://hub.docker.com/r/fsufitch/tagioalisi-bot)

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![TypeScript](https://badgen.net/badge/icon/typescript?icon=typescript&label)](https://typescriptlang.org)
[![Visual Studio](https://badgen.net/badge/icon/visualstudio?icon=visualstudio&label=devcontainer)](https://code.visualstudio.com/docs/remote/containers)


Custom bot for the Sociologic Planning Boar server on Discord. Now in Go!

                 ,           _______________
            `-.   \    .-'  < Hello, world! >
    ,-"`````""-\__ |  /      ---------------
     '-.._    _.-'` '-o,   /
         _>--:{{<   ) |)  /
     .-''      '-.__.-o`
    '-._____..-/`  |  \
         ,-'   /    `-.


For non-Docker build instructions, consult the docs in [`src/bot`](./src/bot) and [`src/web`](./src/web). (TODO) 

## Docker-Compose Stack

Tagioalisi Bot is created to easily run via a Docker Compose stack. Follow these steps.

### 1. Create the configuration

In the repository root, run:

    cp default.env .env

Edit `.env` and follow the instructions there.
Then, `cd runtime`. The remainder of the instructions depend on that.

### 2. Get the Docker images

To run the latest versions hosted on Docker Hub:

    docker-compose pull

Alternatively, build local images.

    make  # Build archives containing precompiled files
    docker-compose build  # Build the actual runtime images

### 3. Initialize the database

    docker-compose run migrations tagioalisi-migrations

Re-run this command anytime database migrations need to be applied.

### 4. Run the servers

    docker-compose up -d

Note: both bot API server and web server expose their traffic via plain HTTP. Add a separate HTTPS layer at your leisure. 

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
