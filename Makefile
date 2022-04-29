.PHONY: clean all

GOFLAGS ?=

all:
	$(MAKE) -C pkg/kalshi

clean:
	@rm -rf pkg/kalshi/swagger
