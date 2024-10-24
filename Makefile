.PHONY: build test run clean docker-build docker-run

# Build variables
BINARY_NAME=device-sec
SERVER_BINARY=server
AGENT_BINARY=agent

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Build flags
LDFLAGS=-ldflags "-w -s"

all: test build

build: 
	$(GOBUILD) -o bin/$(SERVER_BINARY) ./cmd/server
	$(GOBUILD) -o bin/$(AGENT_BINARY) ./cmd/agent

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f bin/$(SERVER_BINARY)
	rm -f bin/$(AGENT_BINARY)

run-server:
	./bin/$(SERVER_BINARY)

run-agent:
	./bin/$(AGENT_BINARY)

docker-build:
	docker build -t $(BINARY_NAME)-server -f Dockerfile.server .
	docker build -t $(BINARY_NAME)-agent -f Dockerfile.agent .

docker-run:
	docker-compose up -d

deps:
	$(GOGET) -v ./...