# Go Template Generator Makefile

BINARY_NAME=go-gen
MAIN_PATH=./cmd/main.go
BIN_DIR=./bin

# Build the generator
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✓ Build complete: $(BIN_DIR)/$(BINARY_NAME)"

# Install dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Run the generator directly
run:
	@go run $(MAIN_PATH) $(ARGS)

# Build and install globally
install: build
	@echo "Installing $(BINARY_NAME) globally..."
	@cp $(BIN_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "✓ $(BINARY_NAME) installed to /usr/local/bin/"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR)
	@echo "✓ Clean complete"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Setup development environment
setup: deps
	@echo "Development environment setup complete!"

.PHONY: build deps run install clean fmt test setup

# For manual development

.PHONY: dev
dev: 
	@echo "Starting development server with hot reload..."
	@air

