package shell

import (
	"bufio"
	"fmt"
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
	"os"
)

func New() *process.Process {
	return process.New("shell", func() {
		logger.Process(0, "Shell started")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Hamix> ")
			cmd, _ := reader.ReadString('\n')
			fmt.Println("You typed: ", cmd)
		}
	})
}
