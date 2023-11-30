.PHONY: default
default: test

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build .

.PHONY: build-bench
build-bench:
	go build -o benchmark-fib ./benchmark
