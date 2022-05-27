GO_PACKAGE=github.com/fsufitch/tagioalisi-bot

PLATFORM_NAME=$(shell echo $$(uname -s)-$$(uname -m) | tr '[:upper:]' '[:lower:]')

EXT=
ifneq (, $(findstring windows,${PLATFORM_NAME}))
EXT=.exe
endif

BIN_BOT=bin/${PLATFORM_NAME}/tagioalisi-bot${EXT}
BIN_MIGRATIONS=bin/${PLATFORM_NAME}/tagioalisi-migrations${EXT}

SOURCES=$(shell find . -type f -a \
	-not -path './bin/*' -a \
	-not -path './.gitignore' -a \
	-not -path './Dockerfile')
SOURCES_MANUAL=$(shell for s in ${SOURCES}; do echo $$s | grep -v '_gen.go'; done)
SOURCES_GENERATED=$(shell for s in ${SOURCES}; do echo $$s | grep '_gen.go'; done)

# === Phonies

.PHONY: all build bot generate proto migrations tests
all: proto generate build

build: bot migrations tests

bot: ${BIN_BOT}

generate: ${SOURCES_GENERATED}

migrations: ${BIN_MIGRATIONS}

tests: ${BIN_TESTS}

# Use Protoc to create the Go Protobuf/GRPC stubs
proto:
	make -C proto


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

${SOURCES_GENERATED}: ${SOURCES_MANUAL}
	@echo Ignoring GOOS/GOARCH setting since code generation should be platform-independent
	GOOS= GOARCH= go run github.com/google/wire/cmd/wire@latest ./...    


# === Actions

.PHONY: test clean

test: ${SOURCES}
	go test ${BIN_TESTS} ./...

clean:
	rm -rf bin/