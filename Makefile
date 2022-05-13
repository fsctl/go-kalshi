.PHONY: clean all

GOFLAGS ?=

VERSION="$(shell go list -m -json github.com/fsctl/go-kalshi@latest | jq --raw-output .Version)"

all: swagger
	$(MAKE) -C cmd/kalshi-tool

swagger:
	mkdir swagger
	curl -X POST https://generator3.swagger.io/api/generate -H 'content-type: application/json' -d '{"specURL" : "https://goelzer.org/kalshi-swagger3.json", "lang" : "go", "type" : "CLIENT", "codegenVersion" : "V3"}' --output swagger/outfile.zip
	unzip -d swagger swagger/outfile.zip
	rm swagger/outfile.zip

update-pkg-cache:
	echo $(VERSION)

clean:
	@rm -rf kalshi-tool
	@rm -rf swagger
