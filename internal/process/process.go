package process

import (
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/syscall"
)

type Process struct {
	PID         int
	Name        string
	Run         func(*Process)
	SyscallChan chan syscall.Syscall
}

func New(name string, runFunc func(*Process)) *Process {
	return &Process{
		Name: name,
		Run:  runFunc,
	}
}

func (p *Process) Start() {
	logger.Process(p.PID, "Starting process")
	p.Run(p)
}
