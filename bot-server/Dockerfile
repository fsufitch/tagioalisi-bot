# ===== Image for building tagi binaries
FROM golang:1.18 AS bot-builder

WORKDIR /bot
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
# Copied too much!
RUN rm -rf terraform webui

RUN go build -x -o ./bin/tagi-bot ./cmd/tagi-bot
RUN go build -x -o ./bin/tagi-migrate ./cmd/tagi-migrate


# ===== Image for running tagi database migrations
FROM busybox:glibc as bot-migrations

WORKDIR /bot
COPY db/migration migration
COPY --from=bot-builder /bot/bin/tagi-migrate .

ENV DEBUG=false \
    DATABASE_URL=postgres://tagi:secret@db:5432/tagi?sslmode=disable \
    MIGRATION_DIR=/bot/migration

CMD [ "/bot/tagi-migrate" ]


# ===== Image for running tagi
FROM busybox:glibc AS bot-runner

WORKDIR /bot
COPY --from=bot-builder /bot/bin/tagi-bot .

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

CMD [ "/bot/tagi-bot" ]
