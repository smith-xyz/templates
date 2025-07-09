package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// ServiceManager handles systemd service operations
type ServiceManager struct {
	Name        string
	Description string
	ServiceFile string
	BinaryPath  string
}

// NewServiceManager creates a new ServiceManager instance
func NewServiceManager(name, description string) *ServiceManager {
	execPath, _ := os.Executable()
	return &ServiceManager{
		Name:        name,
		Description: description,
		ServiceFile: fmt.Sprintf("/etc/systemd/system/%s.service", name),
		BinaryPath:  execPath,
	}
}

// Start starts the service
func (sm *ServiceManager) Start() error {
	return sm.systemctl("start")
}

// Stop stops the service
func (sm *ServiceManager) Stop() error {
	return sm.systemctl("stop")
}

// Restart restarts the service
func (sm *ServiceManager) Restart() error {
	return sm.systemctl("restart")
}

// Status returns the service status
func (sm *ServiceManager) Status() (string, error) {
	cmd := exec.Command("systemctl", "is-active", sm.Name)
	output, err := cmd.Output()
	if err != nil {
		// Check if service exists but is inactive
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 3 {
			return "inactive", nil
		}
		return "unknown", err
	}
	return strings.TrimSpace(string(output)), nil
}

// Install installs the service by creating the systemd service file
func (sm *ServiceManager) Install() error {
	// Check if running as root
	if os.Geteuid() != 0 {
		return fmt.Errorf("installation requires root privileges. Run with sudo")
	}

	// Create the service file
	serviceContent := sm.generateServiceFile()

	// Write the service file
	if err := os.WriteFile(sm.ServiceFile, []byte(serviceContent), 0644); err != nil {
		return fmt.Errorf("failed to create service file: %w", err)
	}

	// Reload systemd daemon
	if err := sm.systemctl("daemon-reload"); err != nil {
		return fmt.Errorf("failed to reload systemd daemon: %w", err)
	}

	// Enable the service
	if err := sm.systemctl("enable"); err != nil {
		return fmt.Errorf("failed to enable service: %w", err)
	}

	return nil
}

// Uninstall removes the service
func (sm *ServiceManager) Uninstall() error {
	// Check if running as root
	if os.Geteuid() != 0 {
		return fmt.Errorf("uninstallation requires root privileges. Run with sudo")
	}

	// Stop the service if running
	sm.Stop()

	// Disable the service
	sm.systemctl("disable")

	// Remove the service file
	if err := os.Remove(sm.ServiceFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove service file: %w", err)
	}

	// Reload systemd daemon
	if err := sm.systemctl("daemon-reload"); err != nil {
		return fmt.Errorf("failed to reload systemd daemon: %w", err)
	}

	return nil
}

// systemctl executes systemctl commands
func (sm *ServiceManager) systemctl(action string) error {
	var cmd *exec.Cmd

	switch action {
	case "daemon-reload":
		cmd = exec.Command("systemctl", "daemon-reload")
	default:
		cmd = exec.Command("systemctl", action, sm.Name)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl %s failed: %w\nOutput: %s", action, err, string(output))
	}

	return nil
}

// generateServiceFile creates the systemd service file content
func (sm *ServiceManager) generateServiceFile() string {
	serviceTemplate := `[Unit]
Description={{.Description}}
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=root
ExecStart={{.BinaryPath}} run
StandardOutput=journal
StandardError=journal
SyslogIdentifier={{.Name}}

[Install]
WantedBy=multi-user.target
`

	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return serviceTemplate // fallback to template string
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, sm)
	if err != nil {
		return serviceTemplate // fallback to template string
	}

	return buf.String()
}

// IsInstalled checks if the service is installed
func (sm *ServiceManager) IsInstalled() bool {
	_, err := os.Stat(sm.ServiceFile)
	return err == nil
}

// GetServiceFile returns the path to the service file
func (sm *ServiceManager) GetServiceFile() string {
	return sm.ServiceFile
}

// GetBinaryPath returns the path to the binary
func (sm *ServiceManager) GetBinaryPath() string {
	return sm.BinaryPath
}

// SetBinaryPath sets a custom binary path (useful for installation)
func (sm *ServiceManager) SetBinaryPath(path string) {
	sm.BinaryPath, _ = filepath.Abs(path)
}
