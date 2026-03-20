.PHONY: build install uninstall

GOOS ?= linux
GOARCH ?= amd64
CGO_ENABLED ?= 0

build:
	@mkdir -p bin
	@for dir in cmd/*/; do \
		name=$$(basename $$dir); \
		CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$$name ./$$dir; \
	done

PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

install: build
	@mkdir -p $(BINDIR)
	@for file in bin/*; do \
		install -m 755 $$file $(BINDIR)/; \
	done

uninstall:
	@for file in bin/*; do \
		rm -f $(BINDIR)/$$(basename $$file); \
	done