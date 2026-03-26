.PHONY: build test clean install lint

BINARY_NAME=clighl
VERSION?=0.1.0

build:
	go build -ldflags "-X github.com/altusmusic/clighl/cmd.version=$(VERSION)" -o $(BINARY_NAME) .

test:
	go test ./... -v

clean:
	rm -f $(BINARY_NAME)
	rm -rf dist/

install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME) 2>/dev/null || mv $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

lint:
	golangci-lint run ./...
