.PHONY: clean all update-pkg-cache

GOFLAGS ?=

GH_VERSION=$(shell GOPROXY=direct go list -m -json github.com/fsctl/go-kalshi@latest | jq --raw-output .Version)

all: swagger
	$(MAKE) -C cmd/kalshi-tool

swagger:
	mkdir swagger
	curl -X POST https://generator3.swagger.io/api/generate -H 'content-type: application/json' -d '{"specURL" : "https://goelzer.org/kalshi-swagger3.json", "lang" : "go", "type" : "CLIENT", "codegenVersion" : "V3"}' --output swagger/outfile.zip
	unzip -d swagger swagger/outfile.zip
	rm swagger/outfile.zip

update-pkg-cache:
	# TODO: not sure `go mod download` actually updates pkg.go.dev
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go mod download github.com/fsctl/go-kalshi@$(GH_VERSION)

clean:
	@rm -rf kalshi-tool
	@rm -rf swagger
