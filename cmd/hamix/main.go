package main

import (
	"hamix-os/internal/kernel"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/programs"
)

func main() {
	//Power on
	logger.Boot("Hamix OS v0.1 is booting...")

	// Initialize kernel
	k := kernel.New()

	k.RegisterProgram("shell", programs.Shell)
	k.RegisterProgram("worker", programs.Worker)

	k.Spawn("shell")

	// Start kernel
	k.Run()
}
