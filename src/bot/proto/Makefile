ROOT := $(shell \
	while [ ! -f .root.txt ]; do \
		cd ..; \
		if [ $$PWD = '/' ]; then \
			echo '.root.txt not found' > /dev/stderr; \
			exit 1; \
		fi; \
	done; \
	echo $$PWD;)
PROTO_SRC_DIR := ${ROOT}/src/proto

# Include GOPATH/bin in PATH in order to get protoc-gen-go and friends.
# If missing, install them with:
#   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
PATH := $(shell echo $$PATH:$$(go env GOPATH)/bin)

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -f *.pb.go

build: *.pb.go
.PHONY: build

*.pb.go: ${PROTO_SRC_DIR}/*.proto
	@echo Ignoring GOOS/GOARCH setting since proto generation is platform-independent
	GOARCH= GOOS= protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			-I ${PROTO_SRC_DIR} \
			${PROTO_SRC_DIR}/*.proto
