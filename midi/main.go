package midi

import (
	"machine"
	"runtime"

	"github.com/litui/monosid/sid"
)

func Task(uart *machine.UART) {
	for !sid.IsReady() {
		runtime.Gosched()
	}

	// log.Logf("MIDI ready")

	for {
		processBuffer(uart)

		runtime.Gosched()
	}
}
