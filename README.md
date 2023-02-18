# Tagioalisi

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


## Modular design

The bot functionality is split up into mostly independent modules that each act as a controller for a smaller set of behavior. They are:

| Name |  Purpose |
| --- |  --- |
| `log` | writes logs of warning+ level into a channel; with `DEBUG` set, also writes info-level logs |
| `memelink` |  registers and responds with meme link content in response to certain "filenames" (e.g. `facepalm.jpg`)  |
| `ping` |  responds to `!ping` with `!pong` for sanity checking |
| `sockpuppet` | hooks allowing for custom messages to be sent via the web UI |
| `groups` | manage a server groups system using special prefixed roles |
| `wiki` | search a variety of wikis |
| `dice` | roll dice (e.g. `!roll 1d20+3`) |
| `news` | use the Bing News Search API to look for recent news about some keywords |

These can be individually turned on/off using the `BLACKLIST_BOT_MODULES` environment variable.

## Docker-Compose Deployment

Tagioalisi Bot is created to easily run via a Docker Compose stack. 

### 1. Configure

In the repository root, edit `.env`. You should change nearly all the values to correspond to your setup.
Only some placeholders for development are included.

> Note: there are other environment variables that don't need to change. Those are in `docker-compose.yml` itself.

A guide to the environment variables:

| Name | Description |
| ---- | ----------- |
| `DISCORD_TOKEN` | Your bot's token, from the "Bot" tab of its [developer page](https://discord.com/developers/applications); this is required for the bot to work at all |
| `DISCORD_LOG_CHANNEL` | The ID of the Discord channel the bot should use for INFO+ logging; right click on the channel in your Discord client to copy the ID |
| `BLACKLIST_BOT_MODULES` | Comma-separated list of bot modules (see above) to not load |
| `OAUTH_CLIENT_ID` / `OAUTH_CLIENT_SECRET` | Your bot's OAuth2 credentials, from the "OAuth2" tab of its [developer page](https://discord.com/developers/applications); these are required for the "Login" feature of the webui to work |
| `OAUTH_REDIRECT_URL` | The URL that Discord should redirect back to as part of the login process; example, for development: `http://localhost:8081/login/redirect` |
| `JWT_HMAC_SECRET` | String to use for signing JWTs; **change in production** |
| `AES_KEY_B64` | 32-byte AES key to encrypt session details with; **change in production**; to generate a new one, run `dd if=/dev/random of=/dev/stdout bs=1 count=32 \| base64` |
| `AZURE_NEWS_SEARCH_KEY` | API key for use with the [Bing News Search API](https://www.microsoft.com/en-us/bing/apis/bing-news-search-api) |
| `MERRIAM_WEBSTER_DICTIONARY_KEY` | API key for use with the [Merriam-Webster Dictionary API]([https://www.microsoft.com/en-us/bing/apis/bing-news-search-api](https://dictionaryapi.com/)) |

### 2. Get the Docker images

To run the latest versions hosted on Docker Hub:

    docker-compose pull

Alternatively, build local images.

    ./build-docker-images.sh
    
The stack relies on three images:

* `docker.io/fsufitch/tagioalisi-discordbot`
* `docker.io/fsufitch/tagioalisi-webapp`
* `docker.io/fsufitch/tagioalisi-grpcwebproxy`

### 3. Run it!

    docker-compose up -d

The stack exposes the following ports:

| Port | Protocol | Description |
| ---- | -------- | ----------- |
| 8090 | HTTPS    | The webapp UI |
| 8091 | HTTP (not S!) | The bot's HTTP API; note this is not secured with TLS |
| 8092 | HTTPS  | The bot's gRPC-web endpoint |

> **TLS Note:** TLS uses self-signed certificates; you should gate _all_ these endpoints behind a proxy that gives them proper certs.

By default, the Postgres database's data is stored in `./var/db`, so it can survive `docker-compose down`. 
If this is displeasing, remove the volume mount from `docker-compose.yml`.

## Docker-Compose Development

Tagioalisi is designed to be developed through the use of [Development containers](https://containers.dev/).
Both the `discordbot/` and `webapp/` directories feature a `.devcontainer.json` which contains a configuration
for launching a devcontainer intended to develop that particular piece of the application. 

The devcontainers use the same `.env` configuration as the production runtime (see above), but a different
Docker Compose file (`docker-compose.dev.yml`) and a different set of images (which include all sorts of tools/goodies). 

**Use your favorite IDE to launch the devcontainers individually.** Visual Studio Code is recommended. 

Alternatively, build and start the dev stack without an IDE by using `docker-compose -f docker-compose.dev.yml up -d`.

The dev stack exposes the following ports:

| Port | Protocol | Description |
| ---- | -------- | ----------- |
| 8080 | HTTPS    | The webapp UI |
| 8081 | HTTP (not S!) | The bot's HTTP API; note this is not secured with TLS |
| 8082 | HTTPS    | The bot's gRPC-web endpoint |
| 9001 | gRPC     | The bot's raw gRPC endpoint |
| 5432 | PostgreSQL | The bot's database |

The dev stack has some notable differences from the production runtime:

* **The bot and webapp processes do not start automatically**. 
  You can use `docker-compose -f docker-compose.dev.yml exec dev-discordbot bash` or 
  `docker-compose -f docker-compose.dev.yml exec dev-webapp bash` to "SSH in". You then need to build
  and run the two processes manually (see below).
* The user is `developer`. It is not root, but has `sudo` within the containers.
* The repository is volume-mounted as `/home/developer/tagioalisi-bot`.
* Your _host user's_ home directory is volume-mounted as `/home/developer/host-home`, for your convenience.

## Developing the bot

These instructions are intended to be used within the devcontainer. Development is possible
without the container, but extra configuration (environment variables) may be necessary.
The bot's code *is cross-platform compatible*, if you wish to develop it in a non-Linux environment.

##### Sources:

Source code relevant to the bot are in the `discordbot/` and `proto/` directories.

##### Build:

Within the `discordbot/` directory:

    ./build.sh

This does three things:

1. Compiles the `proto/*.proto` Protobuf sources into `discordbot/proto/*.pb.go`.
2. Uses [Wire](https://github.com/google/wire) to perform compile-time dependency injection. 
   This generates the somewhat onerous code required to hook together the components of the
   application:
   * `discordbot/cmd/tagi-bot/wire_gen.go`
   * `discordbot/cmd/tagi-migrate/wire_gen.go`
3. Calls `go build` to actually generate `discordbot/bin/tagi-bot` and `discordbot/bin/tagi-migrate`
   (with `.exe` extensions if built on/for Windows).
   
If you want to build the binaries for a different platform, specify the 
GOOS/GOARCH environment variables. To see all combinations available to your current system, run
`go tool dist list`.

##### Run:

Simply run `tagi-bot` for the main bot, and `tagi-migrate` for database migrations. 

##### Test:

Within the `discordbot/` directory, run:

    go test ./...

See `go help test` for relevant arguments.

## Developing the webapp

These instructions are intended to be used within the devcontainer. Development is possible
without the container, but extra configuration (environment variables) may be necessary.
The bot's code *is cross-platform compatible*, if you wish to develop it in a non-Linux environment.

##### Sources:

Source code relevant to the webapp are in the `webapp/` and `proto/` directories.

For initial setup, within the `webapp` directory, run:

    npm install

##### Live server:

Within the `webapp` directory:

    npm run dev

This does two things:

1. Compiles the `proto/*.proto` Protobuf sources into `webapp/src/proto/*.ts`.
2. Starts a [Vite](https://vitejs.dev/) server which compiles and serves the application.
   
The Vite server has HMR (hot module reload) enabled, so the served webpage *should reflect any
changes as soon as you save the source file*! Note that if you make changes to the proto sources,
you will need to re-run the command.

##### Build and bundle static artifacts:

Within the `webapp/` directory:

    npm run build
   
This does two things:

1. Compiles the `proto/*.proto` Protobuf sources into `webapp/src/proto/*.ts`.
2. Uses [Vite](https://vitejs.dev/) to compile the sources and dependencies into static Javascript/etc files.

The output files are found in `dist/`. They can be used with your webserver of choice (nginx, etc)
to serve the webapp.

