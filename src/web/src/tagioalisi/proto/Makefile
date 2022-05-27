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

# OUTPUTS=_out/%
OUTPUTS = ./%.ts %.js

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -rf *.js *.ts

build: ${OUTPUTS}
.PHONY: build


# Use improbable-eng's implementation of grpc-web to generate proto stubs https://github.com/improbable-eng/ts-protoc-gen
#
# However, the output is CJS, while Tagioalisi is ESM (and ts-protoc-gen does not support ESM yet).
# They must be converted, but the cjstoesm converter (understandably) can't understand `goog.object.extend`.
# The goog_to_cjs_export.go script must first be applied so cjstoesm can understand the CJS.
#
# May TC39 have mercy on my soul.
${OUTPUTS}: ${PROTO_SRC_DIR}/*.proto
	mkdir -p _cjs _esm
	protoc \
	    --plugin="protoc-gen-ts=$(shell yarn bin)/protoc-gen-ts" \
		--js_out="import_style=commonjs:_cjs" \
		--ts_out="service=grpc-web:_cjs" \
		-I=${PROTO_SRC_DIR} ${PROTO_SRC_DIR}/*.proto
	go run goog_to_cjs_export.go _cjs/*.js
	set -x; for commonjs_src in *.js; do npx cjstoesm $$commonjs_src _esm; done
	mv _esm/* .
	rm -r _cjs _esm
