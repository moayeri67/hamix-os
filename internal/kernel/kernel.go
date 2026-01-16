package kernel

import (
	"fmt"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
	"hamix-os/internal/syscall"
)

type Kernel struct {
	processes map[int]*process.Process
	nextPID   int
	syscalls  chan syscall.Syscall
}

func New() *Kernel {
	logger.Kernel("Initializing kernel core")

	return &Kernel{
		processes: make(map[int]*process.Process),
		nextPID:   1,
		syscalls:  make(chan syscall.Syscall),
	}
}

func (k *Kernel) RegisterProcess(p *process.Process) {
	pid := k.nextPID
	k.nextPID++

	p.PID = pid
	p.SyscallChan = k.syscalls
	k.processes[pid] = p

	logger.Kernel("Register process: " + p.Name)
}

func (k *Kernel) Run() {
	fmt.Println("Starting scheduler")

	for _, p := range k.processes {
		go p.Start()
	}

	for {
		call := <-k.syscalls
		k.handelSyscall(call)
	}
}

func (k *Kernel) handelSyscall(call syscall.Syscall) {
	switch call.Name {
	case "ps":
		k.handelPS(call)
	default:
		call.Reply <- "Unknown syscall: " + call.Name
	}
}

func (k *Kernel) handelPS(call syscall.Syscall) {
	result := "PID\tNAME\n"

	for _, p := range k.processes {
		result += fmt.Sprintf("%d\t%s\n", p.PID, p.Name)
	}

	call.Reply <- result
}
