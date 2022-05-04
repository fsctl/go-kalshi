.PHONY: clean all

GOFLAGS ?=

all: swagger
	$(MAKE) -C cmd/kalshi-tool

swagger:
	mkdir swagger
	curl -X POST https://generator3.swagger.io/api/generate -H 'content-type: application/json' -d '{"specURL" : "https://goelzer.org/kalshi-swagger3.json", "lang" : "go", "type" : "CLIENT", "codegenVersion" : "V3"}' --output swagger/outfile.zip
	unzip -d swagger swagger/outfile.zip
	rm swagger/outfile.zip

clean:
	@rm -rf kalshi-tool
	@rm -rf swagger