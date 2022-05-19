GO_PACKAGE=github.com/fsufitch/tagioalisi-bot
SOURCES=$(shell find . -type f -a \
	-not -path './bin/*' -a \
	-not -path './.gitignore' -a \
	-not -path './Dockerfile')

PLATFORM_NAME=$(shell echo $$(uname -s)-$$(uname -m) | tr '[:upper:]' '[:lower:]')

EXT=
ifneq (, $(findstring windows,${PLATFORM_NAME}))
EXT=.exe
endif

BIN_BOT=bin/${PLATFORM_NAME}/tagioalisi-bot${EXT}
BIN_MIGRATIONS=bin/${PLATFORM_NAME}/tagioalisi-migrations${EXT}


# === Phonies

.PHONY: all build bot migrations tests
all: build

build: bot migrations tests

bot: ${BIN_BOT}

migrations: ${BIN_MIGRATIONS}

tests: ${BIN_TESTS}

# === Platform phonies

.PHONY: $(shell go tool dist list)

darwin/amd64:
	make GOOS=darwin GOARCH=amd64 PLATFORM_NAME=darwin-x86_64 build

linux/amd64: 
	make GOOS=linux GOARCH=amd64 PLATFORM_NAME=linux-x86_64 build

linux/arm64:
	make GOOS=linux GOARCH=arm64 PLATFORM_NAME=linux-aarch64 build

windows/amd64: 
	make GOOS=windows GOARCH=amd64 PLATFORM_NAME=windows-x86_64 build



# === Real targets


${BIN_BOT}: ${SOURCES}
	go build -o ${BIN_BOT} ${GO_PACKAGE}/cmd/tagi-bot

${BIN_MIGRATIONS}: ${SOURCES}
	go build -o ${BIN_MIGRATIONS} ${GO_PACKAGE}/cmd/tagi-migrate

# === Actions

.PHONY: test clean

test: ${SOURCES}
	go test ${BIN_TESTS} ./...

clean:
	rm -rf bin/