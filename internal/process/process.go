package process

import "hamix-os/internal/platform/logger"

type Process struct {
	PID  int
	Name string
	Run  func()
}

func New(name string, runFunc func()) *Process {
	return &Process{
		Name: name,
		Run:  runFunc,
	}
}

func (p *Process) Start() {
	logger.Process(p.PID, "Starting process")
	p.Run()
}
