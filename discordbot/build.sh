#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd );
cd $SCRIPT_DIR;

echo "Running in: ${SCRIPT_DIR}"

if [[ -z "$PLATFORM_NAME" ]]; then
    PLATFORM_NAME=$(echo $(uname -s)-$(uname -m) | tr '[:upper:]' '[:lower:]');
fi

EXT='';
if [[ "$PLATFORM_NAME" =~ "windows" ]]; then
    EXT=.exe
fi

echo "Platform: '${PLATFORM_NAME}' Extension: '${EXT}'"

PROTO_SRC_DIR=${SCRIPT_DIR}/../proto
PROTO_GEN_DIR=${SCRIPT_DIR}/proto

BIN_DIR=${SCRIPT_DIR}/bin
BIN_BOT=${BIN_DIR}/tagi-bot
BIN_MIGRATIONS=${BIN_DIR}/tagi-migrate

# Start build commands, echo them
set -x

# Generate protobuf
GOARCH= GOOS= protoc \
    --go_out=${PROTO_GEN_DIR} --go_opt=paths=source_relative \
    --go-grpc_out=${PROTO_GEN_DIR} --go-grpc_opt=paths=source_relative \
    -I ${PROTO_SRC_DIR} \
    ${PROTO_SRC_DIR}/*.proto

# Generate wired files
GOOS= GOARCH= go run github.com/google/wire/cmd/wire ./...    

# Build
go build -o ${BIN_BOT} ./cmd/tagi-bot
go build -o ${BIN_MIGRATIONS} ./cmd/tagi-migrate

# Done
set +x

echo Done
