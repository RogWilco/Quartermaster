.PHONY: all build clean run test lint tidy help

DIR_BUILD ?= bin
DIR_DIST ?= dist

# Binary name
BINARY_NAME=qm

all: clean build

build:
	@echo "Building..."
	@mkdir -p $(DIR_BUILD)
	@go build -o $(DIR_BUILD)/$(BINARY_NAME) cmd/qm/main.go

clean:
	@echo "Cleaning..."
	@rm -rf $(DIR_BUILD)

run:
	@go run cmd/qm/main.go

test:
	@echo "Running tests..."
	@go test ./...

lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed"; \
		echo "Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

help:
	@echo "Available commands:"
	@echo "  make build    - Build the application"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make lint     - Run linter"
	@echo "  make tidy     - Tidy dependencies"
	@echo "  make help     - Show this help message"