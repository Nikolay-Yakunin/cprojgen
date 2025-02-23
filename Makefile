BINARY_NAME = cprojgen
BUILD_DIR = build
SRC = cmd/cprojgen/main.go

.PHONY: all build-linux build-windows build clean

all: build

build: build-linux build-windows

$(BUILD_DIR):
	@echo "Создаем папку $(BUILD_DIR)..."
	mkdir -p $(BUILD_DIR)

build-linux: $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

build-windows: $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME).exe main.go

fmt:
	go fmt ./...

test:
	go test ./...


clean:
	rm -rf $(BUILD_DIR)
