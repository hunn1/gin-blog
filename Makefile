# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=morningo
BINARY_UNIX=$(BINARY_NAME)_unix

all: run

build:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)
	rm -f ./build/$(BINARY_UNIX)

run:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./
	./build/$(BINARY_NAME)

restart:
	kill -INT $$(cat pid)
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./
	./build/$(BINARY_NAME)

mod:
	go mod tidy
	go mod vendor

cross:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./

fmt:
	go fmt ./app/...
	go fmt ./bootstrap/...
	go fmt ./config/...
	go fmt ./Exceptions/...
	go fmt ./helpers/...
	go fmt ./library/...
	go fmt ./routes/...