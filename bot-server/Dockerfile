# ===== Image for building tagi binaries
FROM golang:1.18 AS bot-builder

WORKDIR /bot
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

RUN make


# ===== Image for running tagi database migrations
# TODO: eliminate the need on the .sql migration files (bundle?)
# FROM busybox:glibc as bot-migrations

FROM bot-builder as bot-migrations

WORKDIR /bot
COPY db/migration migration

ENV DEBUG=true \
    DATABASE_URL=postgres://tagi:secret@db:5432/tagi?sslmode=disable \
    MIGRATION_DIR=/bot/migration

CMD [ "/bot/bin/tagioalisi-migrations" ]


# ===== Image for running tagi
FROM golang:1.18 AS bot-runtime

WORKDIR /bot
COPY --from=bot-builder /bot/bin/tagioalisi-bot .

ENV DEBUG=false \
    PORT=80 \
    WEB_ENABLED=true \
    WEB_SECRET= \
    DISCORD_TOKEN= \
    DISCORD_LOG_CHANNEL=327526752203177984 \
    DATABASE_URL=postgres://tagi:secret@db:5432/tagi?sslmode=disable \
    BLACKLIST_BOT_MODULES= \
    OAUTH_CLIENT_ID= \
    OAUTH_CLIENT_SECRET= \
    OAUTH_AUTH_ENDPOINT=https://discordapp.com/api/oauth2/authorize \
    OAUTH_TOKEN_ENDPOINT=https://discordapp.com/api/oauth2/token \
    OAUTH_REDIRECT_URL= \
    JWT_HMAC_SECRET=supersecret \
    AES_KEY_B64=zvFXk6LbvG5jTG7JlGBQ57MdqUwjdS/wH3gVWXkSOUU= \
    AZURE_NEWS_SEARCH_KEY= \
    MERRIAM_WEBSTER_DICTIONARY_KEY= 

CMD [ "/bot/tagioalisi-bot" ]
