package service

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewServiceManager(t *testing.T) {
	name := "test-service"
	description := "Test service description"

	sm := NewServiceManager(name, description)

	if sm.Name != name {
		t.Errorf("Expected name %s, got %s", name, sm.Name)
	}

	if sm.Description != description {
		t.Errorf("Expected description %s, got %s", description, sm.Description)
	}

	expectedServiceFile := "/etc/systemd/system/test-service.service"
	if sm.ServiceFile != expectedServiceFile {
		t.Errorf("Expected service file %s, got %s", expectedServiceFile, sm.ServiceFile)
	}

	if sm.BinaryPath == "" {
		t.Error("Expected binary path to be set")
	}
}

func TestGenerateServiceFile(t *testing.T) {
	sm := NewServiceManager("test-service", "Test Description")
	sm.BinaryPath = "/usr/local/bin/test-service"

	content := sm.generateServiceFile()

	// Check that the content contains expected sections
	expectedSections := []string{
		"[Unit]",
		"[Service]",
		"[Install]",
		"Description=Test Description",
		"ExecStart=/usr/local/bin/test-service run",
		"WantedBy=multi-user.target",
	}

	for _, section := range expectedSections {
		if !containsString(content, section) {
			t.Errorf("Service file content missing expected section: %s", section)
		}
	}
}

func TestSetBinaryPath(t *testing.T) {
	sm := NewServiceManager("test-service", "Test Description")

	testPath := "/custom/path/to/binary"
	sm.SetBinaryPath(testPath)

	absPath, _ := filepath.Abs(testPath)
	if sm.BinaryPath != absPath {
		t.Errorf("Expected binary path %s, got %s", absPath, sm.BinaryPath)
	}
}

func TestIsInstalled(t *testing.T) {
	sm := NewServiceManager("test-service", "Test Description")

	// For a non-existent service file, should return false
	if sm.IsInstalled() {
		t.Error("Expected service to not be installed")
	}

	// Create a temporary service file to test positive case
	tempDir := t.TempDir()
	tempServiceFile := filepath.Join(tempDir, "test-service.service")

	// Create the file
	file, err := os.Create(tempServiceFile)
	if err != nil {
		t.Fatalf("Failed to create temp service file: %v", err)
	}
	file.Close()

	// Update the service file path to point to our temp file
	sm.ServiceFile = tempServiceFile

	// Now it should be installed
	if !sm.IsInstalled() {
		t.Error("Expected service to be installed")
	}
}

func TestGetServiceFile(t *testing.T) {
	sm := NewServiceManager("test-service", "Test Description")

	expected := "/etc/systemd/system/test-service.service"
	if sm.GetServiceFile() != expected {
		t.Errorf("Expected service file %s, got %s", expected, sm.GetServiceFile())
	}
}

func TestGetBinaryPath(t *testing.T) {
	sm := NewServiceManager("test-service", "Test Description")

	if sm.GetBinaryPath() == "" {
		t.Error("Expected binary path to be set")
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		len(s) > len(substr) && containsString(s[1:], substr)
}

// Alternative implementation using a simple search
func containsString2(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark test to compare containsString implementations
func BenchmarkContainsString(b *testing.B) {
	text := "This is a test string with some content to search through"
	search := "test"

	b.Run("Implementation1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			containsString(text, search)
		}
	})

	b.Run("Implementation2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			containsString2(text, search)
		}
	})
}
