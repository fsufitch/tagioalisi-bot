
Place `bot.tar.gz` and `web.tar.gz` here to build the Tagioalisi runtime. Do not check them into source control.

These each have a specific expected file layout:

## `bot.tar.gz`

    bot.tar.gz/
    ├── linux-x86_64/
    │   ├── tagioalisi-bot
    │   └── tagioalisi-migrate
    ├── linux-aarch64/
    │   ├── tagioalisi-bot
    │   └── tagioalisi-migrate
    └── <...>

Each file in the directories is a binary executable on the OS/processor as specified by the directory prefix. The prefixes are required for building a Docker image targeting a particular system. They must be in the format `$(uname -s)-$(uname -m)`, lowercased. Try:

    echo $(uname -s)-$(uname -m) | tr '[:upper:]' '[:lower:]'

Linux binaries must be compiled for `glibc`, not `uclibc` or `musl`.

The bot's API is HTTP-based and serves traffic on port 80. Expose it as you wish.

## `web.tar.gz`

    web.tar.gz/
    ├── index.html
    └── <...>

Everything will be served by the web server. The container's Nginx config is amended at runtime using the `TAGIOALISI_BOT_BASE_URL` environment variable.

The web server is HTTP-based and serves traffic (internally) on port 80. Expose it as you wish.