default: build

SOURCES=.babelrc .eslintrc.js .yarnrc.yml package.json \
	tsconfig.json webpack.config.ts yarn.lock \
	$(shell find ./src/ -type f -a \
	-not -path './dist/*' -a \
	-not -path './node_modules/*')

BINARIES=./dist $(shell find ./dist/ -type f)

WEBPACK_MODE=production

##### Phonies

.PHONY: proto build clean
build: ${BINARIES}

proto:
	make -C src/tagioalisi/proto

clean:
	rm -rf dist
	mkdir -p dist

##### Real targets

${BINARIES}: ${SOURCES}
	yarn run webpack:esm --mode ${WEBPACK_MODE}
	find dist | xargs touch