.PHONY: fmt test

fmt:
	go fmt ./...

test:
	go test -cover ./...

all: fmt test
