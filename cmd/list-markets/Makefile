VERSION ?= $(shell git describe --tags --always --dirty)

../../list-markets : *.go ../../pkg/kalshi/*.go
	go build -ldflags="-X 'main.Version=$(VERSION)'" -o ../../list-markets *.go
