.PHONY: clean all

GOFLAGS ?=

all:
	$(MAKE) -C pkg/kalshi
	$(MAKE) -C cmd/kalshi-tool

clean:
	@rm -rf kalshi-tool
	@rm -rf pkg/kalshi/swagger