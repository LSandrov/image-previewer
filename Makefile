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

.PHONY: build
build:
	$(GO_BUILD) -ldflags="-X 'main.shaCommit=$(SHA_COMMIT)'" -mod vendor -trimpath -o ./bin/app ./cmd/server

.PHONY: tests
test-functional:
	$(GO) test -race -cover -tags musl -short -v \
				-coverprofile profile.cov.tmp -p 100 \
				./pkg/... ./internal/...
	cat profile.cov.tmp | grep -v "_gen.go" > profile.cov
	$(MAKE) cover

.PHONY: cover
cover:
	$(GO) tool cover -func profile.cov
################################################################################################ lint

.PHONY: lint
lint:
	golangci-lint run -v