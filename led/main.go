package led

import (
	"machine"
	"runtime"
)

var (
	led   = machine.LED
	state bool
)

// Change LED state so it flashes once
func Flash() {
	state = true
}

// Monitors state of LED and flashes it if state changes
func Task() {
	state = false

	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led.Low()

	for {
		if state {
			led.High()
			for i := 0; i < 1000; i++ {
				runtime.Gosched()
			}

			led.Low()
			for i := 0; i < 1000; i++ {
				runtime.Gosched()
			}
			state = false
		}

		runtime.Gosched()
	}
}
