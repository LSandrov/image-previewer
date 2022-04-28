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
	$(GO_BUILD) -ldflags="-X 'main.shaCommit=$(SHA_COMMIT)'" -trimpath -o ./bin/app ./cmd
################################################################################################ lint

.PHONY: lint
lint:
	golangci-lint run -v