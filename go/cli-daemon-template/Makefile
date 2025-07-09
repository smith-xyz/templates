BINARY_NAME=cli-service-manager
VERSION=1.0.0
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

.PHONY: all build clean test install uninstall help

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@sudo chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete. You can now run: $(BINARY_NAME)"

uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete."

install-service: install
	@echo "Installing and starting service..."
	@sudo $(INSTALL_DIR)/$(BINARY_NAME) install
	@sudo $(INSTALL_DIR)/$(BINARY_NAME) start
	@echo "Service installed and started."

uninstall-service:
	@echo "Stopping and uninstalling service..."
	@sudo $(INSTALL_DIR)/$(BINARY_NAME) stop || true
	@sudo $(INSTALL_DIR)/$(BINARY_NAME) uninstall || true
	@echo "Service uninstalled."

status:
	@$(INSTALL_DIR)/$(BINARY_NAME) status

logs:
	@echo "Showing service logs..."
	@sudo journalctl -u $(BINARY_NAME) -f

help:
	@echo "Available targets:"
	@echo "  build           - Build the binary"
	@echo "  clean           - Clean build artifacts"
	@echo "  test            - Run tests"
	@echo "  deps            - Download and tidy dependencies"
	@echo "  install         - Install binary to system"
	@echo "  uninstall       - Remove binary from system"
	@echo "  install-service - Install and start the service"
	@echo "  uninstall-service - Stop and uninstall the service"
	@echo "  status          - Check service status"
	@echo "  logs            - Show service logs"
	@echo "  help            - Show this help message" 