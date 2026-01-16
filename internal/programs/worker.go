package programs

import (
	"hamix-os/internal/platform/logger"
	"hamix-os/internal/process"
	"time"
)

func Worker(p *process.Process) {
	logger.Process(p.PID, "Worker Started")

	for i := 0; i <= 5; i++ {
		logger.Process(p.PID, "Working... step"+string(rune('0'+i)))
		time.Sleep(1 * time.Second)
	}

	logger.Process(p.PID, "Worker Finished")
}
