# CLI Service Manager

 ⚠️ Project generated by AI

A template Go application that can run as a Linux system service with systemd. This project provides a complete example of how to create a CLI application that can be installed, started, stopped, and managed as a system service.

## Features

- ✅ CLI interface with commands for service management
- ✅ Systemd service integration
- ✅ Graceful shutdown handling
- ✅ Automatic service installation/uninstallation
- ✅ Service status monitoring
- ✅ Logging to systemd journal
- ✅ Automatic restart on failure
- ✅ Easy to extend and customize

## Commands

```bash
cli-service-manager <command>

Commands:
  start      Start the service
  stop       Stop the service
  restart    Restart the service
  status     Check service status
  install    Install the service (requires sudo)
  uninstall  Uninstall the service (requires sudo)
  run        Run the service (used by systemd)
```

## Installation

### Prerequisites

- Go 1.19 or higher
- Linux with systemd
- sudo privileges (for service installation)

### Build and Install

1. Clone or download the project
2. Build the project:
   ```bash
   make build
   ```

3. Install the binary to the system:
   ```bash
   make install
   ```

### Quick Service Installation

To build, install, and start the service in one step:
```bash
make install-service
```

## Usage

### Manual Installation Steps

1. **Build the project:**
   ```bash
   make build
   ```

2. **Install the binary:**
   ```bash
   make install
   ```

3. **Install the service:**
   ```bash
   sudo cli-service-manager install
   ```

4. **Start the service:**
   ```bash
   sudo cli-service-manager start
   ```

5. **Check service status:**
   ```bash
   cli-service-manager status
   ```

### Service Management

- **Start the service:**
  ```bash
  sudo cli-service-manager start
  ```

- **Stop the service:**
  ```bash
  sudo cli-service-manager stop
  ```

- **Restart the service:**
  ```bash
  sudo cli-service-manager restart
  ```

- **Check service status:**
  ```bash
  cli-service-manager status
  ```

- **View service logs:**
  ```bash
  sudo journalctl -u cli-service-manager -f
  ```
  Or use the Makefile:
  ```bash
  make logs
  ```

### Uninstallation

To completely remove the service and binary:
```bash
make uninstall-service
make uninstall
```

Or step by step:
```bash
sudo cli-service-manager stop
sudo cli-service-manager uninstall
sudo rm /usr/local/bin/cli-service-manager
```

## Development

### Project Structure

```
cli-service-manager/
├── main.go              # Main CLI application
├── service/
│   └── service.go       # Service management logic
├── Makefile            # Build and installation scripts
├── README.md           # This file
└── go.mod              # Go module definition
```

### Building from Source

```bash
# Download dependencies
make deps

# Build the project
make build

# Run tests
make test

# Clean build artifacts
make clean
```

### Customization

To customize this service for your needs:

1. **Modify the service logic** in `main.go` in the `runService()` function
2. **Update service name and description** in the constants at the top of `main.go`
3. **Add configuration** by creating a config file and parsing it in the service
4. **Add more CLI commands** by extending the switch statement in `main.go`

### Example Service Logic

The default service runs every 30 seconds and logs a message. You can replace this with your own logic:

```go
func runService() {
    // Your service logic here
    // This could be:
    // - A web server
    // - A background worker
    // - A monitoring service
    // - A file watcher
    // - etc.
}
```

## Systemd Service File

The service automatically generates a systemd service file at `/etc/systemd/system/cli-service-manager.service`:

```ini
[Unit]
Description=CLI Service Manager - A template service
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=root
ExecStart=/usr/local/bin/cli-service-manager run
StandardOutput=journal
StandardError=journal
SyslogIdentifier=cli-service-manager

[Install]
WantedBy=multi-user.target
```

## Logging

The service logs to the systemd journal. You can view logs using:

```bash
# View recent logs
sudo journalctl -u cli-service-manager

# Follow logs in real-time
sudo journalctl -u cli-service-manager -f

# View logs since last boot
sudo journalctl -u cli-service-manager -b
```

## Troubleshooting

### Service won't start
- Check if the binary exists: `ls -la /usr/local/bin/cli-service-manager`
- Check service status: `systemctl status cli-service-manager`
- Check logs: `sudo journalctl -u cli-service-manager`

### Permission denied
- Make sure you're using `sudo` for service management commands
- Check binary permissions: `ls -la /usr/local/bin/cli-service-manager`

### Service file not found
- Reinstall the service: `sudo cli-service-manager install`
- Check if service file exists: `ls -la /etc/systemd/system/cli-service-manager.service`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is provided as-is for educational and development purposes. Modify and use as needed for your projects.

## Support

For issues and questions, please check the logs first using `make logs` or `sudo journalctl -u cli-service-manager`. 