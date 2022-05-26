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

# Include yarn bin, to access protoc-gen-ts_proto
PATH := $(shell echo $$PATH:$$(yarn bin))

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -f *.ts

build: *.ts
.PHONY: build

*.ts: ${PROTO_SRC_DIR}/*.proto
	protoc --ts_proto_out=. \
			-I ${PROTO_SRC_DIR} \
			${PROTO_SRC_DIR}/*.proto