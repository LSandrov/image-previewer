ifeq ($(OS),Windows_NT)
    GOOS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        GOOS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        GOOS := darwin
    endif
endif
GOARCH = amd64
CGO_ENABLED = 0
GO = $(shell which go)
GO_BUILD = GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build
SHA_COMMIT = $(shell git rev-parse --short HEAD)

################################################################################################ build

.PHONY: build
build:
	$(GO_BUILD) -ldflags="-X 'main.shaCommit=$(SHA_COMMIT)'" -trimpath -o ./bin/app ./cmd

################################################################################################ clear

.PHONY: clear
clear:
	rm -rf bin/*

################################################################################################ lint

.PHONY: lint
lint:
	golangci-lint run -v

################################################################################################ test

.PHONY: test
test:
	$(GO) test -race -cover -short -v \
				-coverprofile profile.cov.tmp -p 2 \
				./pkg/... ./internal/...
	cat profile.cov.tmp | grep -v "_gen.go" > profile.cov
	$(MAKE) cover

.PHONY: cover
cover:
	$(GO) tool cover -func profile.cov

.PHONY: test-e2e
test-e2e:
	$(GO) clean -testcache
	$(GO) test -short ./e2e/... -v
################################################################################################ docker

.PHONY: run
run:
	env UID=${UID} GID=${GID} docker-compose up -d

.PHONY: stop
stop:
	docker-compose down