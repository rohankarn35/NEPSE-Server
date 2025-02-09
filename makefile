.PHONY: all run build clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
BINARY_NAME=server
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME)_windows.exe
BINARY_MAC=$(BINARY_NAME)_mac

all: run

run:
	$(GORUN) server.go

build: build-linux build-windows build-mac

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_UNIX) server.go

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_WINDOWS) server.go

build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_MAC) server.go

clean:
	rm -f bin/$(BINARY_UNIX) bin/$(BINARY_WINDOWS) bin/$(BINARY_MAC)