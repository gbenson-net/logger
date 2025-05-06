all: build

.PHONY: build check test

check: test

test:
	gofmt -w .
	go vet .
	go test -v -coverprofile=coverage.out .
