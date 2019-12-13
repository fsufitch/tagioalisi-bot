#!/bin/sh

export LOG_LEVEL=debug

export PORT=9999
export WEB_ENABLED=true
export WEB_SECRET=secret123
export DISCORD_TOKEN=
export DISCORD_LOG_LEVEL=info
export DISCORD_LOG_CHANNEL=327526752203177984
export DATABASE_URL=postgres://boarbot:boarbot@localhost:5432/boarbot
export BLACKLIST_BOT_MODULES=
export MIGRATION_DIR=./db/migration