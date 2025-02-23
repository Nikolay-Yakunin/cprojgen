.PHONY: all build-linux build-windows clean

MODULE_NAME=cprojgen
BINARY_NAME=cprojgen
SRC_DIR=./cmd/$(MODULE_NAME)
BUILD_DIR=./build

GOOS_LINUX=linux
GOARCH_AMD64=amd64
GOOS_WINDOWS=windows
GOARCH_AMD64_WINDOWS=amd64

all: format check-style check-dependencies build-linux build-windows

format:
	@echo "Formatting code..."
	@gofmt -s -w $(shell find . -type f -name '*.go' ! -path "./vendor/*")

check-style:
	@echo "Checking code style..."
	@gofmt -s -d $(shell find . -type f -name '*.go' ! -path "./vendor/*") | tee /dev/tty | (! read)
	@if [ $$? -eq 0 ]; then \
	    echo "Code style is correct."; \
	else \
	    echo "Code style issues found. Please run 'gofmt -s -w .' to fix them."; \
	    exit 1; \
	fi

check-dependencies:
	@echo "Checking dependencies..."
	@go mod tidy
	@go mod verify
	@if [ -n "$$(git status --porcelain go.mod go.sum)" ]; then \
	    echo "Dependencies are not up-to-date. Please run 'go mod tidy' and commit the changes."; \
	    git diff -- go.mod go.sum; \
	    exit 1; \
	fi

build-linux:
	@mkdir -p $(BUILD_DIR)/$(GOOS_LINUX)
	GOOS=$(GOOS_LINUX) GOARCH=$(GOARCH_AMD64) go build -o $(BUILD_DIR)/$(GOOS_LINUX)/$(BINARY_NAME) ./cmd/$(MODULE_NAME)

build-windows:
	@mkdir -p $(BUILD_DIR)/$(GOOS_WINDOWS)
	GOOS=$(GOOS_WINDOWS) GOARCH=$(GOARCH_AMD64_WINDOWS) go build -o $(BUILD_DIR)/$(GOOS_WINDOWS)/$(BINARY_NAME).exe ./cmd/$(MODULE_NAME)


test:
	go test ./...


clean:
	rm -rf $(BUILD_DIR)
