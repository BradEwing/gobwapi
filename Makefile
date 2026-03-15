# Makefile for gobwapi — Go bindings for BWAPI (StarCraft: Brood War AI)

SHELL := cmd.exe

BINARY_DIR := bin
EXAMPLEBOT := $(BINARY_DIR)\examplebot.exe

# Default: build, vet, and test everything.
.PHONY: all
all: build vet test

# Build all packages and the example bot binary.
.PHONY: build
build:
	@if not exist $(BINARY_DIR) mkdir $(BINARY_DIR)
	go build ./...
	go build -o $(EXAMPLEBOT) ./cmd/examplebot

# Run all tests.
.PHONY: test
test:
	go test ./...

# Run Go vet for static analysis.
.PHONY: vet
vet:
	go vet ./...

# Format all Go source files.
.PHONY: fmt
fmt:
	go fmt ./...

# CI-like check: vet + test + verify formatting.
.PHONY: check
check: vet test
	@echo Checking formatting...
	@gofmt -l .

# Remove build artifacts and Go caches.
.PHONY: clean
clean:
	go clean -cache
	@if exist $(BINARY_DIR) rmdir /s /q $(BINARY_DIR)
