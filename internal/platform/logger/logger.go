package logger

import "fmt"

func Boot(message string) {
	fmt.Println("[BOOT]", message)
}

func Kernel(message string) {
	fmt.Println("[Kernel]", message)
}

func Process(pid int, message string) {
	fmt.Printf("[Process %d] %s\n", pid, message)
}
