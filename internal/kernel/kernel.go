package kernel

import (
	"fmt"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
	"hamix-os/internal/syscall"
	"strconv"
)

type Kernel struct {
	processes map[int]*process.Process
	nextPID   int
	syscalls  chan syscall.Syscall
	program   map[string]func(*process.Process)
}

func New() *Kernel {
	logger.Kernel("Initializing kernel core")

	return &Kernel{
		processes: make(map[int]*process.Process),
		nextPID:   1,
		syscalls:  make(chan syscall.Syscall),
		program:   make(map[string]func(*process.Process)),
	}
}

func (k *Kernel) Spawn(programName string) (int, error) {
	entry, exist := k.program[programName]
	if !exist {
		return 0, fmt.Errorf("no such a program: %s", programName)
	}

	p := process.New(programName, entry)
	k.RegisterProcess(p)

	go p.Start()

	return p.PID, nil
}

func (k *Kernel) RegisterProgram(program string, entry func(*process.Process)) {
	k.program[program] = entry
	logger.Kernel("Register program: " + program)
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
	case "kill":
		k.handelKill(call)
	case "fork":
		k.handelFork(call)
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

func (k *Kernel) handelKill(call syscall.Syscall) {
	if len(call.ARGS) < 1 {
		call.Reply <- "Usage: Kill <PID>\n"
		return
	}

	pidStr := call.ARGS[0]
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		call.Reply <- fmt.Sprintf("Invalid PID: %s", pidStr)
		return
	}

	if pid == call.PID {
		call.Reply <- "Process cannot kill itself. \n"
	}

	proc, exist := k.processes[pid]
	if !exist {
		call.Reply <- "Process does not exist"
	}

	delete(k.processes, pid)

	call.Reply <- fmt.Sprintf("Process %d (%s) Terminated\n", pid, proc.Name)
}

func (k *Kernel) handelFork(call syscall.Syscall) {
	if len(call.ARGS) < 1 {
		call.Reply <- "Usage: Fork <program>\n"
		return
	}

	program := call.ARGS[0]

	pid, error := k.Spawn(program)
	if error != nil {
		call.Reply <- error.Error()
		return
	}

	call.Reply <- fmt.Sprintf("Spawned process %s with PID %d\"", program, pid)
}
