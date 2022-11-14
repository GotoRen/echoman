GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc
GOLANGCI=golangci-lint
DUMPCMD=tcpdump

INTERFACE_REAL=eth0
INTERFACE_VIRTUAL=echoman_tun



all: build run

build: ## go build
	$(GOCMD) ./cmd/main.go

build/pi: ## go build for armv7l
	GOOS=linux GOARCH=arm GOARM=7 $(GOBUILD) ./cmd/main.go

build/pi64: ## go build for aarch64
	GOOS=linux GOARCH=arm64 $(GOBUILD) ./cmd/main.go

run: ## go run
	$(GORUN) ./cmd/main.go

test: ## go test
	$(GOTEST) -v ./...

benchmark: ## go benchmark
	$(GOTEST) -bench -v ./...

lint: ## golang-ci lint
	$(GOLANGCI) run --config=.golangci.yaml ./...

doc: ## godoc http:6060
	$(GODOC) -http=:6060

coverage: ## coverage
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

dump-real: ## dump packet
	$(DUMPCMD) -i ${INTERFACE_REAL}

dump-virtual: ## dump packet
	$(DUMPCMD) -i ${INTERFACE_VIRTUAL}

dump-real/out: ## dump packet and output to file
	$(DUMPCMD) -i ${INTERFACE_REAL} -w ./debug/dumpfile-server-real.pcapng

dump-virtual/out: ## dump packet and output to file
	$(DUMPCMD) -i ${INTERFACE_VIRTUAL} -w ./debug/dumpfile-server-virtual.pcapng

listen: ## check listen port
	lsof -i -P



help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
