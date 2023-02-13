GO_PACKAGE=github.com/fsufitch/tagioalisi-bot

PLATFORM_NAME=$(shell echo $$(uname -s)-$$(uname -m) | tr '[:upper:]' '[:lower:]')

EXT=
ifneq (, $(findstring windows,${PLATFORM_NAME}))
EXT=.exe
endif


SOURCES=$(shell find . -type f -a \
	-not -path './bin/*' -a \
	-not -path './.gitignore' -a \
	-not -path './Dockerfile')

# SOURCES_MANUAL=$(shell for s in ${SOURCES}; do echo $$s | grep -v '_gen.go'; done)
# SOURCES_GENERATED=$}GENERATED_TIMESTAMP} $(shell for s in ${SOURCES}; do echo $$s | grep '_gen.go'; done)

PROTO_SRC=$(shell find ../proto-src}
WIRE_OUTPUT=$(shell find proto)
WIRE_OUTPUT=$(shell find . -name wire_gen.go)

MANUAL_SRC=$(shell find . -type f -a \
	-not -path './bin/*' -a \
	-not -path './.gitignore' -a \
	-not -name wire_gen.go -a \
	-not -path './proto/*')

BOT_OUTPUT=bin/${PLATFORM_NAME}/tagioalisi-bot${EXT} bin/${PLATFORM_NAME}/tagioalisi-migrations${EXT}

# === Phonies

.PHONY: all
all: proto build test

.PHONY: build
build: ${BIN_BOT} ${BIG_MIGRATIONS}

.PHONY: protos
protos: ${SOURCES_GENERATED}

.PHONY: migrations
migrations: ${BIN_MIGRATIONS}

.PHONY: test
test: ${BIN_TESTS}

# Use Protoc to create the Go Protobuf/GRPC stubs
${SOURCES_GENERATED}: 
	*.pb.go: ${PROTO_SRC_DIR}/*.proto


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
	GOOS= GOARCH= go get -u github.com/google/wire/cmd/wire@latest
	GOOS= GOARCH= go run github.com/google/wire/cmd/wire ./...    


	@echo Ignoring GOOS/GOARCH setting since proto generation is platform-independent
	GOARCH= GOOS= protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			-I ${PROTO_SRC_DIR} \
			${PROTO_SRC_DIR}/*.proto
		echo 

# === Actions

.PHONY: test clean

test: ${SOURCES}
	go test ${BIN_TESTS} ./...

clean:
	rm -rf bin/