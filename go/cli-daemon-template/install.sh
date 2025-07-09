#!/bin/bash

# CLI Service Manager Installation Script
# This script builds and installs the CLI service manager

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BINARY_NAME="cli-service-manager"
INSTALL_DIR="/usr/local/bin"
SERVICE_NAME="cli-service-manager"

# Functions
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

check_root() {
    if [[ $EUID -ne 0 ]]; then
        print_error "This script requires root privileges. Please run with sudo."
        exit 1
    fi
}

check_dependencies() {
    print_info "Checking dependencies..."
    
    # Check if Go is installed
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go 1.19 or higher."
        exit 1
    fi
    
    # Check Go version
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_info "Found Go version: $GO_VERSION"
    
    # Check if systemctl is available
    if ! command -v systemctl &> /dev/null; then
        print_error "systemctl is not available. This script requires systemd."
        exit 1
    fi
    
    print_success "All dependencies are satisfied"
}

build_service() {
    print_info "Building the service..."
    
    # Clean previous builds
    if [ -d "build" ]; then
        rm -rf build
    fi
    
    # Build the service
    make build
    
    if [ ! -f "build/$BINARY_NAME" ]; then
        print_error "Build failed. Binary not found."
        exit 1
    fi
    
    print_success "Service built successfully"
}

install_binary() {
    print_info "Installing binary to $INSTALL_DIR..."
    
    # Copy binary to install directory
    cp "build/$BINARY_NAME" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
    
    # Verify installation
    if [ ! -f "$INSTALL_DIR/$BINARY_NAME" ]; then
        print_error "Binary installation failed"
        exit 1
    fi
    
    print_success "Binary installed successfully"
}

install_service() {
    print_info "Installing systemd service..."
    
    # Install the service
    "$INSTALL_DIR/$BINARY_NAME" install
    
    print_success "Service installed successfully"
}

start_service() {
    print_info "Starting the service..."
    
    # Start the service
    "$INSTALL_DIR/$BINARY_NAME" start
    
    # Check if service is running
    sleep 2
    STATUS=$("$INSTALL_DIR/$BINARY_NAME" status)
    
    if [ "$STATUS" = "active" ]; then
        print_success "Service started successfully"
    else
        print_warning "Service may not have started properly. Status: $STATUS"
    fi
}

show_status() {
    print_info "Service status:"
    "$INSTALL_DIR/$BINARY_NAME" status
    
    print_info "You can check logs with:"
    echo "  sudo journalctl -u $SERVICE_NAME -f"
    
    print_info "You can manage the service with:"
    echo "  sudo $BINARY_NAME start"
    echo "  sudo $BINARY_NAME stop"
    echo "  sudo $BINARY_NAME restart"
    echo "  $BINARY_NAME status"
}

uninstall_service() {
    print_info "Uninstalling service..."
    
    # Stop and uninstall service
    if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
        "$INSTALL_DIR/$BINARY_NAME" stop 2>/dev/null || true
        "$INSTALL_DIR/$BINARY_NAME" uninstall 2>/dev/null || true
        rm -f "$INSTALL_DIR/$BINARY_NAME"
        print_success "Service uninstalled successfully"
    else
        print_warning "Service binary not found"
    fi
}

usage() {
    echo "Usage: $0 [install|uninstall|build|status]"
    echo ""
    echo "Commands:"
    echo "  install    - Build and install the service (default)"
    echo "  uninstall  - Uninstall the service"
    echo "  build      - Build the service only"
    echo "  status     - Show service status"
    echo ""
    echo "Examples:"
    echo "  sudo $0 install"
    echo "  sudo $0 uninstall"
    echo "  $0 build"
    echo "  $0 status"
    exit 1
}

main() {
    local command=${1:-install}
    
    case $command in
        install)
            check_root
            check_dependencies
            build_service
            install_binary
            install_service
            start_service
            show_status
            ;;
        uninstall)
            check_root
            uninstall_service
            ;;
        build)
            check_dependencies
            build_service
            print_success "Build complete. Use 'sudo $0 install' to install."
            ;;
        status)
            if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
                show_status
            else
                print_error "Service not installed. Run 'sudo $0 install' first."
                exit 1
            fi
            ;;
        *)
            usage
            ;;
    esac
}

# Run main function with all arguments
main "$@"