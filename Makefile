.PHONY: clean all

VERSION ?= $(shell git describe --tags --always --dirty)

GOFLAGS ?=

all:
	$(MAKE) -C pkg/kalshi
	$(MAKE) -C cmd

clean:
	@rm list-markets
	@rm -rf pkg/kalshi/swagger
