VERSION ?= "dev"
LDFLAGS=-ldflags "-s -w -X github.com/axllent/wiregod/cmd.Version=${VERSION}"
BINARY=wiregod
UPX := $(shell which upx)

build = echo "\n\nBuilding $(1)-$(2)" && GOOS=$(1) GOARCH=$(2) go build ${LDFLAGS} -o dist/${BINARY}_${VERSION}_$(1)_$(2) \
	&& if [ "${UPX}" != "" ]; then ${UPX} -q -9 dist/${BINARY}_${VERSION}_$(1)_$(2); fi \
	&& bzip2 dist/${BINARY}_${VERSION}_$(1)_$(2) \
	&& if [ $(1) = "windows" ]; then mv dist/${BINARY}_${VERSION}_$(1)_$(2).bz2 dist/${BINARY}_${VERSION}_$(1)_$(2).exe.bz2; fi

build: *.go go.*
	go build ${LDFLAGS} -o ${BINARY}
	@if [ "${UPX}" != "" ]; then ${UPX} -q -9 ${BINARY}; fi
	rm -rf /tmp/go-*

clean:
	rm -f ${BINARY}

release:
	mkdir -p dist
	rm -f dist/${BINARY}_${VERSION}_*
	$(call build,linux,amd64,)
	$(call build,linux,386,)
	# $(call build,darwin,amd64,)
