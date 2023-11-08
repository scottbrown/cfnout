.DEFAULT_GOAL := build-linux

pwd := $(shell pwd)

pkg := cfnout
repo := github.com/scottbrown/$(pkg)

build.dir := $(pwd)/.build
dist.dir  := $(pwd)/.dist

linux.bin   := cfnout
windows.bin := cfnout.exe

build.windows.file  := $(build.dir)/windows/$(windows.bin)
build.linux.dir     := $(build.dir)/linux
build.linux.file    := $(build.linux.dir)/$(linux.bin)
build.darwin.dir    := $(build.dir)/darwin
build.darwin.file   := $(build.darwin.dir)/$(linux.bin)

dist.windows.filename := $(pkg)-$(VERSION)-windows-amd64.zip
dist.linux.filename   := $(pkg)-$(VERSION)-linux-amd64.tar.gz
dist.linux.file       := $(dist.dir)/$(dist.linux.filename)
dist.darwin.filename  := $(pkg)-$(VERSION)-darwin-amd64.tar.gz
dist.darwin.file      := $(dist.dir)/$(dist.darwin.filename)

.PHONY: build
build: build-linux build-windows build-darwin

.PHONY: test
test:
	go test ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: dist
dist: dist-linux dist-windows dist-darwin

.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(build.windows.file) $(repo)

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(build.linux.file) $(repo)

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(build.darwin.file) $(repo)

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
	tar cfz $(dist.linux.file) -C $(build.linux.dir) $(linux.bin)

.PHONY: dist-darwin
dist-darwin: _get-version
	@mkdir -p $(dist.dir)
	tar cfz $(dist.darwin.file) -C $(build.darwin.dir) $(linux.bin)

.PHONY: _get-version
_get-version:
ifndef VERSION
	@echo "Provide a VERSION to continue."; exit 1
endif
