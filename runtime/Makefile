ROOT=$(shell echo $$(while [ ! -e .root.txt ]; do cd ..; done && pwd))
GO_PLATFORMS=linux/amd64 linux/arm64
WEBPACK_MODE=production

# Phony

.PHONY: check-dependencies clean default binpkg

default: check-dependencies binpkg

check-dependencies:
	go version
	@ldd $$(which go) | grep musl && (echo "WARNING: $$(which go) is linked against libmusl, which is incompatible with the supported Tagioalisi Docker setup. Check yourself!" >dev/stderr; sleep 3s) ||:
	yarn versions
	@docker-compose version || (echo "WARNING: docker-compose not found. You will need it to run the Tagioalisi Docker runtime." >/dev/stderr && sleep 3s) ||:

clean:
	rm binpkg/*.tar.gz ||:

binpkg: binpkg/bot.tar.gz binpkg/web.tar.gz



# Actual targets

binpkg/bot.tar.gz: $(shell find ${ROOT}/src/bot/bin -type f)
	make -C "${ROOT}/src/bot" clean ${GO_PLATFORMS}
	tar cfz binpkg/bot.tar.gz -C "${ROOT}/src/bot/bin" $$(cd "${ROOT}/src/bot/bin" && find . -type f)
	@echo "Created bot bundle in binpkg/bot.tar.gz"

binpkg/web.tar.gz: $(shell find ${ROOT}/src/web/dist -type f)
	cd "${ROOT}/src/web"; yarn; yarn run webpack --mode ${WEBPACK_MODE}
	tar cfz binpkg/web.tar.gz -C "${ROOT}/src/web/dist" $$(cd "${ROOT}/src/web/dist" && find . -type f)
	@echo "Created web bundle in binpkg/web.tar.gz"
