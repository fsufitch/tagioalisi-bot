#!/usr/bin/env bash

export PATH="$(go env GOPATH)/bin:$PATH"

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
cd $SCRIPT_DIR

echo "Running in: ${SCRIPT_DIR}"

(
    set -x
    go version
)

if [[ -z "$PLATFORM_NAME" ]]; then
    PLATFORM_NAME=$(echo $(uname -s)-$(uname -m) | tr '[:upper:]' '[:lower:]')
fi

EXT=''
if [[ "$PLATFORM_NAME" =~ "windows" ]]; then
    EXT=.exe
fi

if [[ -n "$SKIP_PROTOC" ]]; then
    echo "NOTE: SKIP_PROTOC is set; protobuf stubs will not be regenerated"
fi

if [[ -n "$SKIP_WIRE" ]]; then
    echo "NOTE: SKIP_WIRE is set; wire dependency injection will not be regenerated"
fi

echo "Platform: '${PLATFORM_NAME}' Extension: '${EXT}'"

PROTO_SRC_DIR=$(
    cd ${SCRIPT_DIR}/../proto
    pwd
)
PROTO_GEN_DIR=${SCRIPT_DIR}/proto

BIN_DIR=${SCRIPT_DIR}/bin
BIN_BOT=${BIN_DIR}/tagi-bot
BIN_MIGRATIONS=${BIN_DIR}/tagi-migrate

set -e

# Generate protobuf
if [[ -z "$SKIP_PROTOC" ]]; then
    echo "Updating protobuf generated stubs..."
    (
        set -x
        GOARCH= GOOS= protoc \
            --go_out=${PROTO_GEN_DIR} --go_opt=paths=source_relative \
            --go-grpc_out=${PROTO_GEN_DIR} --go-grpc_opt=paths=source_relative \
            -I ${PROTO_SRC_DIR} \
            ${PROTO_SRC_DIR}/*.proto
    )
fi

# Generate wired files
if [[ -z "$SKIP_WIRE" ]]; then
    echo "Updating wire generated sources..."
    echo $PATH
    (
        set -x
        GOOS= GOARCH= wire ./...
    )
fi

# Build
echo "Building actual binaries..."
(
    set -x

    go build -o ${BIN_BOT} ./cmd/tagi-bot
    go build -o ${BIN_MIGRATIONS} ./cmd/tagi-migrate
)

# Done

echo Done
