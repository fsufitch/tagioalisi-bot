GO_PACKAGE=github.com/fsufitch/tagioalisi-bot
BIN_BOT=bin/tagioalisi-bot
BIN_MIGRATIONS=bin/tagioalisi-migrations
SOURCES=$(shell git ls-files)

# === Phonies

.PHONY: all build bot migrations tests
all: build

build: bot migrations tests

bot: ${BIN_BOT}

migrations: ${BIN_MIGRATIONS}

tests: ${BIN_TESTS}

# === Real targets


${BIN_BOT}: ${SOURCES}
	go build -x -o ${BIN_BOT} ${GO_PACKAGE}/cmd/tagi-bot

${BIN_MIGRATIONS}: ${SOURCES}
	go build -x -o ${BIN_MIGRATIONS} ${GO_PACKAGE}/cmd/tagi-migrate

# === Actions

.PHONY: test clean

test: ${SOURCES}
	go test ${BIN_TESTS} ./...

clean:
	rm -rf bin/