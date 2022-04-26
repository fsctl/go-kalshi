.PHONY: clean all

VERSION ?= $(shell git describe --tags --always --dirty)

GOFLAGS ?=

all:
	$(MAKE) -C cmd

clean:
	@rm list-markets
