package syscall

type Syscall struct {
	PID   int
	Name  string
	ARGS  []string
	Reply chan string
}
