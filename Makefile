VERSION ?= dev
BINARY  := relx
PKG     := github.com/ademajagon/relx/cmd

build:
	go build -ldflags "-X $(PKG).version=$(VERSION)" -o $(BINARY) .

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f $(BINARY)

install: build
	mv $(BINARY) /usr/local/bin/$(BINARY)

.PHONY: build test lint clean install