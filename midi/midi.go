package midi

import (
	"machine"
	"runtime"
)

func Task(uart *machine.UART) {
	for {
		processBuffer(uart)

		runtime.Gosched()
	}
}
