.DEFAULT_GOAL := build-linux

pwd := $(shell pwd)

pkg := cfnout

build.dir := $(pwd)/.build
dist.dir  := $(pwd)/.dist

linux.bin   := cfnout
windows.bin := cfnout.exe

build.windows.file    := $(build.dir)/windows/$(windows.bin)
build.linux.file      := $(build.dir)/linux/$(linux.bin)
build.darwin.file     := $(build.dir)/darwin/$(linux.bin)

dist.windows.filename := $(pkg)-$(VERSION)-windows-amd64.zip
dist.linux.filename   := $(pkg)-$(VERSION)-linux-amd64.zip
dist.darwin.filename  := $(pkg)-$(VERSION)-darwin-amd64.zip

.PHONY: build
build: build-linux build-windows build-darwin

.PHONY: format
format:
	go fmt ./...

.PHONY: dist
dist: dist-linux dist-windows dist-darwin

.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(build.windows.file) $(pkg)

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(build.linux.file) $(pkg)

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(build.darwin.file) $(pkg)

.PHONY: clean
clean:
	rm -rf $(build.dir) $(dist.dir)

.PHONY: dist-windows
dist-windows: _get-version
	@mkdir -p $(dist.dir)
	cd $(dist.dir) && zip $(dist.windows.filename) $(build.windows.file)

.PHONY: dist-linux
dist-linux: _get-version
	@mkdir -p $(dist.dir)
	cd $(dist.dir) && zip $(dist.linux.filename) $(build.linux.file)

.PHONY: dist-darwin
dist-darwin: _get-version
	@mkdir -p $(dist.dir)
	cd $(dist.dir) && zip $(dist.darwin.filename) $(build.darwin.file)

.PHONY: _get-version
_get-version:
ifndef VERSION
	@echo "Provide a VERSION to continue."; exit 1
endif
