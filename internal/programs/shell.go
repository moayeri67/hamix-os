package programs

import (
	"bufio"
	"fmt"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
	"hamix-os/internal/syscall"
	"os"
	"strings"
)

func Shell(p *process.Process) {
	logger.Process(p.PID, "Shell started")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Hamix> ")
		cmd, _ := reader.ReadString('\n')
		line := strings.TrimSpace(cmd)
		params := strings.Split(line, " ")
		command := params[0]

		args := []string{}
		if len(params) > 1 {
			args = params[1:]
		}

		reply := make(chan string)

		call := syscall.Syscall{
			PID:   p.PID,
			Name:  command,
			ARGS:  args,
			Reply: reply,
		}

		p.SyscallChan <- call
		result := <-reply

		fmt.Println(result)
	}
}
