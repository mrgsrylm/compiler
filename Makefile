.PHONY: default
default: test

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build .