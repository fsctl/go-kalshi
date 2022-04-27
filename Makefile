.PHONY: clean all

GOFLAGS ?=

all:
	$(MAKE) -C pkg/kalshi
	$(MAKE) -C cmd

clean:
	@rm list-markets
	@rm -rf pkg/kalshi/swagger
