package kernel

import (
	"fmt"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
)

type Kernel struct {
	processes map[int]*process.Process
	nextPID   int
}

func New() *Kernel {
	logger.Kernel("Initializing kernel core")

	return &Kernel{
		processes: make(map[int]*process.Process),
		nextPID:   1,
	}
}

func (k *Kernel) RegisterProcess(p *process.Process) {
	pid := k.nextPID
	k.nextPID++

	p.PID = pid
	k.processes[pid] = p

	logger.Kernel("Register process: " + p.Name)
}

func (k *Kernel) Run() {
	fmt.Println("Starting scheduler")

	for _, p := range k.processes {
		go p.Start()
	}

	// Keep kernel alive
	select {}
}
