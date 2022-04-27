.PHONY: clean all

VERSION ?= $(shell git describe --tags --always --dirty)

GOFLAGS ?=

all: #generate
	$(MAKE) -C cmd

#generate:
#	rm -rf

clean:
	@rm list-markets
