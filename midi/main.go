package midi

import (
	"machine"
	"runtime"

	"github.com/litui/monosid/log"
)

func Task(uart *machine.UART) {
	log.Logf("MIDI ready")

	for {
		processBuffer(uart)

		runtime.Gosched()
	}
}
