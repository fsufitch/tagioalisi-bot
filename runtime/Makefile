ROOT=$(shell echo $$(while [ ! -e .root.txt ]; do cd ..; done && pwd))
GO_PLATFORMS=linux/amd64 linux/arm64
WEBPACK_MODE=production

BOT_BINARIES=$(shell find ${ROOT}/src/bot/bin -type f)
WEB_BINARIES=$(shell find ${ROOT}/src/web/dist -type f)

# Phony

.PHONY: bin-bot bin-web check-dependencies clean default binpkg

default: check-dependencies binpkg

check-dependencies:
	go version
	@ldd $$(which go) | grep musl && (echo "WARNING: $$(which go) is linked against libmusl, which is incompatible with the supported Tagioalisi Docker setup. Check yourself!" >dev/stderr; sleep 3s) ||:
	yarn versions
	@docker-compose version || (echo "WARNING: docker-compose not found. You will need it to run the Tagioalisi Docker runtime." >/dev/stderr && sleep 3s) ||:

clean:
	rm binpkg/*.tar.gz ||:

binpkg: binpkg/bot.tar.gz binpkg/web.tar.gz

bin-bot:
	make -C "${ROOT}/src/bot" ${GO_PLATFORMS}

bin-web:
	make -C "${ROOT}/src/web"

# Actual targets

binpkg/bot.tar.gz: bin-bot
	tar cfz binpkg/bot.tar.gz -C "${ROOT}/src/bot/bin" $$(cd "${ROOT}/src/bot/bin" && find . -type f)
	@echo "Created bot bundle in binpkg/bot.tar.gz"

binpkg/web.tar.gz: bin-web
	tar cfz binpkg/web.tar.gz -C "${ROOT}/src/web/dist" $$(cd "${ROOT}/src/web/dist" && find . -type f)
	@echo "Created web bundle in binpkg/web.tar.gz"
