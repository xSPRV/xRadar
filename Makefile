.PHONY: build run clean test

APP_NAME = xradar
BUILD_DIR = bin

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/xradar

run: build
	@./$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

test:
	@go test -v ./...
