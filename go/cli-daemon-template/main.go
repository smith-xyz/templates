package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cli-daemon-template/service"
)

const (
	serviceName = "cli-daemon-templater"
	serviceDesc = "CLI Daemon Template - A template service"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	sm := service.NewServiceManager(serviceName, serviceDesc)

	switch command {
	case "start":
		if err := sm.Start(); err != nil {
			fmt.Printf("Error starting service: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Service started successfully")

	case "stop":
		if err := sm.Stop(); err != nil {
			fmt.Printf("Error stopping service: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Service stopped successfully")

	case "restart":
		if err := sm.Restart(); err != nil {
			fmt.Printf("Error restarting service: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Service restarted successfully")

	case "status":
		status, err := sm.Status()
		if err != nil {
			fmt.Printf("Error checking service status: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Service status: %s\n", status)

	case "install":
		if err := sm.Install(); err != nil {
			fmt.Printf("Error installing service: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Service installed successfully")

	case "uninstall":
		if err := sm.Uninstall(); err != nil {
			fmt.Printf("Error uninstalling service: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Service uninstalled successfully")

	case "run":
		// This is called by systemd to actually run the service
		runService()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf(`Usage: %s <command>

Commands:
  start      Start the service
  stop       Stop the service
  restart    Restart the service
  status     Check service status
  install    Install the service (requires sudo)
  uninstall  Uninstall the service (requires sudo)
  run        Run the service (used by systemd)

Example:
  sudo %s install
  sudo %s start
  %s status
  sudo %s stop
  sudo %s uninstall
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func runService() {
	fmt.Println("Service is starting...")

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a ticker for periodic work (example: every 30 seconds)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	fmt.Println("Service is running... (Press Ctrl+C to stop)")

	for {
		select {
		case <-ticker.C:
			// Do periodic work here
			fmt.Printf("[%s] Service is running and doing work...\n", time.Now().Format("2006-01-02 15:04:05"))

		case sig := <-sigChan:
			fmt.Printf("Received signal: %s. Shutting down gracefully...\n", sig)

			// Perform cleanup here
			cleanup()

			fmt.Println("Service stopped.")
			return
		}
	}
}

func cleanup() {
	// Add any cleanup logic here
	fmt.Println("Performing cleanup...")
	time.Sleep(1 * time.Second)
	fmt.Println("Cleanup completed.")
}
