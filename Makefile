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
clear: ## Clear the working area and the project
	rm -rf bin/*

################################################################################################ lint

.PHONY: lint
lint:
	golangci-lint run -v

################################################################################################ test

.PHONY: test
test: ## Run test ./pkg/... ./internal/...
	$(GO) test -race -cover -short -v \
				-coverprofile profile.cov.tmp -p 2 \
				./pkg/... ./internal/...
	cat profile.cov.tmp | grep -v "_gen.go" > profile.cov
	$(MAKE) cover

.PHONY: cover
cover:
	$(GO) tool cover -func profile.cov

################################################################################################ docker

.PHONY: docker-up
docker-up: ## Start containers
	make build && docker-compose up -d

.PHONY: docker-down
docker-down: ## Stop containers
	docker-compose down