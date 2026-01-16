package main

import (
	"hamix-os/internal/kernel"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/shell"
)

func main() {
	//Power on
	logger.Boot("Hamix OS v0.1 is booting...")

	// Initialize kernel
	k := kernel.New()

	// Load first user process (shell)
	s := shell.New()

	// Register process in kernel
	k.RegisterProcess(s)

	// Start kernel
	k.Run()
}
