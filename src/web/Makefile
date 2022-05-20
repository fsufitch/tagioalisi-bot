default: build

SOURCES=.babelrc .eslintrc.js .yarnrc.yml package.json \
	tsconfig.json webpack.config.ts yarn.lock \
	$(shell find ./src/ -type f -a \
	-not -path './dist/*' -a \
	-not -path './node_modules/*')

BINARIES=$(shell find ./dist/ -type f)

WEBPACK_MODE=production

##### Phonies

.PHONY: build clean
build: ${BINARIES}

clean:
	rm -rf dist
	mkdir -p dist

##### Real targets

${BINARIES}: ${SOURCES}
	yarn run webpack --mode ${WEBPACK_MODE}
	find dist | xargs touch