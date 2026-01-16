package shell

import (
	"bufio"
	"fmt"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
	"hamix-os/internal/syscall"
	"os"
	"strings"
)

func New() *process.Process {
	return process.New("shell", func(p *process.Process) {
		logger.Process(p.PID, "Shell started")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Hamix> ")
			cmd, _ := reader.ReadString('\n')

			reply := make(chan string)

			call := syscall.Syscall{
				PID:   p.PID,
				Name:  strings.TrimSpace(cmd),
				ARGS:  []string{},
				Reply: reply,
			}

			p.SyscallChan <- call
			result := <-reply

			fmt.Println(result)
		}
	})
}
