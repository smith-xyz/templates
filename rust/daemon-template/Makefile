# Rust Daemon Template Makefile

# Project information
PROJECT_NAME = daemon-template
BINARY_NAME = daemon-template
TARGET_DIR = target
RELEASE_DIR = $(TARGET_DIR)/release
DEBUG_DIR = $(TARGET_DIR)/debug

# Colors for output
RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[1;33m
BLUE = \033[0;34m
NC = \033[0m # No Color

# Default target
.PHONY: all
all: build

# Help target
.PHONY: help
help:
	@echo "$(BLUE)Available targets:$(NC)"
	@echo "  $(GREEN)build$(NC)          - Build debug version"
	@echo "  $(GREEN)build-release$(NC)  - Build release version"
	@echo "  $(GREEN)run$(NC)            - Run debug version"
	@echo "  $(GREEN)run-release$(NC)    - Run release version"
	@echo "  $(GREEN)test$(NC)           - Run all tests"
	@echo "  $(GREEN)test-coverage$(NC)  - Run tests with coverage"
	@echo "  $(GREEN)bench$(NC)          - Run benchmarks"
	@echo "  $(GREEN)fmt$(NC)            - Format code"
	@echo "  $(GREEN)fmt-check$(NC)      - Check code formatting"
	@echo "  $(GREEN)clippy$(NC)         - Run clippy lints"
	@echo "  $(GREEN)clean$(NC)          - Clean build artifacts"
	@echo "  $(GREEN)deps$(NC)           - Install dependencies"
	@echo "  $(GREEN)update$(NC)         - Update dependencies"
	@echo "  $(GREEN)check$(NC)          - Run all checks (fmt, clippy, test)"
	@echo "  $(GREEN)install$(NC)        - Install binary to system"
	@echo "  $(GREEN)uninstall$(NC)      - Uninstall binary from system"
	@echo "  $(GREEN)docker-build$(NC)   - Build Docker image"
	@echo "  $(GREEN)docker-run$(NC)     - Run Docker container"
	@echo "  $(GREEN)logs$(NC)           - View daemon logs (systemd)"
	@echo "  $(GREEN)status$(NC)         - Check daemon status"

# Build targets
.PHONY: build
build:
	@echo "$(BLUE)Building debug version...$(NC)"
	cargo build

.PHONY: build-release
build-release:
	@echo "$(BLUE)Building release version...$(NC)"
	cargo build --release

# Run targets
.PHONY: run
run:
	@echo "$(BLUE)Running debug version...$(NC)"
	cargo run

.PHONY: run-release
run-release: build-release
	@echo "$(BLUE)Running release version...$(NC)"
	./$(RELEASE_DIR)/$(BINARY_NAME)

.PHONY: run-background
run-background: build-release
	@echo "$(BLUE)Running daemon in background...$(NC)"
	nohup ./$(RELEASE_DIR)/$(BINARY_NAME) > daemon.log 2>&1 &
	@echo "$(GREEN)Daemon started. Check daemon.log for output.$(NC)"

# Test targets
.PHONY: test
test:
	@echo "$(BLUE)Running tests...$(NC)"
	cargo test

.PHONY: test-coverage
test-coverage:
	@echo "$(BLUE)Running tests with coverage...$(NC)"
	cargo test --all-features --no-fail-fast
	@echo "$(YELLOW)Note: Install cargo-tarpaulin for detailed coverage: cargo install cargo-tarpaulin$(NC)"

.PHONY: bench
bench:
	@echo "$(BLUE)Running benchmarks...$(NC)"
	cargo bench

# Code quality targets
.PHONY: fmt
fmt:
	@echo "$(BLUE)Formatting code...$(NC)"
	cargo fmt

.PHONY: fmt-check
fmt-check:
	@echo "$(BLUE)Checking code formatting...$(NC)"
	cargo fmt --check

.PHONY: clippy
clippy:
	@echo "$(BLUE)Running clippy lints...$(NC)"
	cargo clippy --all-targets --all-features -- -D warnings

.PHONY: check
check: fmt-check clippy test
	@echo "$(GREEN)All checks passed!$(NC)"

# Dependency management
.PHONY: deps
deps:
	@echo "$(BLUE)Installing dependencies...$(NC)"
	cargo fetch

.PHONY: update
update:
	@echo "$(BLUE)Updating dependencies...$(NC)"
	cargo update

.PHONY: audit
audit:
	@echo "$(BLUE)Auditing dependencies...$(NC)"
	cargo audit

# Installation targets
.PHONY: install
install: build-release
	@echo "$(BLUE)Installing binary to system...$(NC)"
	cargo install --path .

.PHONY: uninstall
uninstall:
	@echo "$(BLUE)Uninstalling binary from system...$(NC)"
	cargo uninstall $(BINARY_NAME)

# Docker targets
.PHONY: docker-build
docker-build:
	@echo "$(BLUE)Building Docker image...$(NC)"
	docker build -t $(PROJECT_NAME) .

.PHONY: docker-run
docker-run: docker-build
	@echo "$(BLUE)Running Docker container...$(NC)"
	docker run --rm --name $(PROJECT_NAME) $(PROJECT_NAME)

.PHONY: docker-run-detached
docker-run-detached: docker-build
	@echo "$(BLUE)Running Docker container in detached mode...$(NC)"
	docker run -d --name $(PROJECT_NAME) $(PROJECT_NAME)

# System service targets (for systemd)
.PHONY: service-install
service-install: build-release
	@echo "$(BLUE)Installing systemd service...$(NC)"
	@echo "$(YELLOW)Note: This requires root privileges$(NC)"
	sudo cp daemon-template.service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl enable $(PROJECT_NAME)

.PHONY: service-start
service-start:
	@echo "$(BLUE)Starting systemd service...$(NC)"
	sudo systemctl start $(PROJECT_NAME)

.PHONY: service-stop
service-stop:
	@echo "$(BLUE)Stopping systemd service...$(NC)"
	sudo systemctl stop $(PROJECT_NAME)

.PHONY: service-restart
service-restart:
	@echo "$(BLUE)Restarting systemd service...$(NC)"
	sudo systemctl restart $(PROJECT_NAME)

.PHONY: status
status:
	@echo "$(BLUE)Checking daemon status...$(NC)"
	sudo systemctl status $(PROJECT_NAME)

.PHONY: logs
logs:
	@echo "$(BLUE)Viewing daemon logs...$(NC)"
	sudo journalctl -u $(PROJECT_NAME) -f

# Development targets
.PHONY: dev
dev:
	@echo "$(BLUE)Starting development server with auto-reload...$(NC)"
	cargo watch -x run

.PHONY: dev-test
dev-test:
	@echo "$(BLUE)Running tests with auto-reload...$(NC)"
	cargo watch -x test

# Cleanup targets
.PHONY: clean
clean:
	@echo "$(BLUE)Cleaning build artifacts...$(NC)"
	cargo clean

.PHONY: clean-all
clean-all: clean
	@echo "$(BLUE)Cleaning all artifacts including logs...$(NC)"
	rm -f daemon.log
	rm -f nohup.out

# Documentation targets
.PHONY: doc
doc:
	@echo "$(BLUE)Building documentation...$(NC)"
	cargo doc --open

.PHONY: doc-deps
doc-deps:
	@echo "$(BLUE)Building documentation with dependencies...$(NC)"
	cargo doc --document-private-items --open

# Release targets
.PHONY: release
release: check build-release
	@echo "$(GREEN)Release build completed successfully!$(NC)"
	@echo "$(BLUE)Binary location: $(RELEASE_DIR)/$(BINARY_NAME)$(NC)"

.PHONY: package
package: release
	@echo "$(BLUE)Creating release package...$(NC)"
	tar -czf $(PROJECT_NAME)-$(shell date +%Y%m%d-%H%M%S).tar.gz -C $(RELEASE_DIR) $(BINARY_NAME)
	@echo "$(GREEN)Package created successfully!$(NC)"

# Debug targets
.PHONY: debug
debug:
	@echo "$(BLUE)Running with debug output...$(NC)"
	RUST_LOG=debug cargo run

.PHONY: trace
trace:
	@echo "$(BLUE)Running with trace output...$(NC)"
	RUST_LOG=trace cargo run

# Stop running daemon
.PHONY: stop
stop:
	@echo "$(BLUE)Stopping daemon...$(NC)"
	@pkill -f $(BINARY_NAME) || echo "$(YELLOW)No running daemon found$(NC)"

# Show running processes
.PHONY: ps
ps:
	@echo "$(BLUE)Checking for running daemon processes...$(NC)"
	@pgrep -f $(BINARY_NAME) || echo "$(YELLOW)No daemon processes found$(NC)" 