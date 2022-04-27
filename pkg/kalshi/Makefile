swagger:
	rm -rf swagger
	mkdir swagger
	curl -X POST https://generator3.swagger.io/api/generate -H 'content-type: application/json' -d '{"specURL" : "http://goelzer.org/kalshi-swagger3.json", "lang" : "go", "type" : "CLIENT", "codegenVersion" : "V3"}' --output swagger/outfile.zip
	unzip -d swagger swagger/outfile.zip
	rm swagger/outfile.zip
