# Rust Daemon Template

A robust template for creating daemon services in Rust with async runtime, graceful shutdown, and proper logging.

## Features

- **Async Runtime**: Built with Tokio for high-performance async operations
- **Graceful Shutdown**: Handles SIGTERM and other signals for clean termination
- **Structured Logging**: Uses `tracing` for structured, configurable logging
- **Periodic Tasks**: Executes work at regular intervals
- **Error Handling**: Comprehensive error handling throughout the daemon lifecycle
- **Signal Handling**: Proper Unix signal handling for production deployment

## Quick Start

### Prerequisites

- Rust 1.70+ (2024 edition)
- Cargo

### Installation

```bash
# Clone or copy this template
cd daemon-template

# Build the project
cargo build

# Run the daemon
cargo run
```

### Running in Production

```bash
# Build optimized release version
cargo build --release

# Run the daemon
./target/release/daemon-template

# Or use the Makefile
make run-release
```

## Usage

The daemon will:

1. Start up and initialize logging
2. Begin periodic work every 10 seconds
3. Log all activities with timestamps
4. Handle shutdown signals gracefully

### Stopping the Daemon

- **Interactive**: Press `Ctrl+C`
- **Signal**: `kill -TERM <pid>`
- **Systemd**: `systemctl stop daemon-template`

## Configuration

### Log Levels

Set the log level using the `RUST_LOG` environment variable:

```bash
# Debug level
RUST_LOG=debug cargo run

# Info level (default)
RUST_LOG=info cargo run

# Warning level only
RUST_LOG=warn cargo run
```

### Customization

The daemon is designed to be easily customizable:

1. **Work Interval**: Modify `Duration::from_secs(10)` in the main loop
2. **Work Logic**: Implement your business logic in the `perform_work` function
3. **Additional Signals**: Add more signal handlers in `handle_signals`

## Development

### Building

```bash
# Debug build
make build

# Release build
make build-release

# Or with cargo directly
cargo build
cargo build --release
```

### Testing

```bash
# Run tests
make test

# Run with coverage
make test-coverage
```

### Linting

```bash
# Check code formatting
make fmt-check

# Format code
make fmt

# Run clippy lints
make clippy
```

## Deployment

### Systemd Service

Create a systemd service file at `/etc/systemd/system/daemon-template.service`:

```ini
[Unit]
Description=Rust Daemon Template
After=network.target

[Service]
Type=simple
User=daemon-user
Group=daemon-group
WorkingDirectory=/opt/daemon-template
ExecStart=/opt/daemon-template/target/release/daemon-template
Restart=always
RestartSec=10
Environment=RUST_LOG=info

[Install]
WantedBy=multi-user.target
```

Then enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable daemon-template
sudo systemctl start daemon-template
```

### Docker

```bash
# Build Docker image
docker build -t daemon-template .

# Run container
docker run -d --name daemon-template daemon-template
```

## Architecture

The daemon follows a clean architecture with:

- **Main Function**: Orchestrates startup, signal handling, and shutdown
- **Signal Handler**: Manages Unix signals for graceful termination
- **Work Loop**: Performs periodic business logic
- **Error Handling**: Comprehensive error management throughout

## Dependencies

- `tokio`: Async runtime and utilities
- `tracing`: Structured logging framework
- `tracing-subscriber`: Log formatting and output
- `signal-hook`: Unix signal handling
- `signal-hook-tokio`: Async signal handling integration

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run `make check` to verify all checks pass
6. Submit a pull request

## License

This template is provided as-is for educational and development purposes.

## Troubleshooting

### Common Issues

1. **Permission Denied**: Ensure the binary has execute permissions
2. **Port Already in Use**: Check if another instance is running
3. **Signal Not Handled**: Verify signal handling is working with `kill -TERM`

### Debugging

Enable debug logging:

```bash
RUST_LOG=debug cargo run
```

Check system logs:

```bash
journalctl -u daemon-template -f
```

## Examples

### Custom Work Function

```rust
async fn perform_work(iteration: u64) -> Result<(), Box<dyn std::error::Error>> {
    // Your custom logic here
    match iteration % 3 {
        0 => process_queue().await?,
        1 => cleanup_old_files().await?,
        2 => send_heartbeat().await?,
        _ => unreachable!(),
    }
    Ok(())
}
```

### Adding Configuration

```rust
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
struct Config {
    interval_seconds: u64,
    log_level: String,
    work_dir: String,
}
``` 