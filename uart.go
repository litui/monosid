package main

import (
	"machine"
	"time"

	"github.com/litui/monosid/config"
)

func uartInit() *machine.UART {
	uart := config.MIDI_UART
	err := uart.Configure(machine.UARTConfig{
		BaudRate: 31250,
		TX:       config.PIN_MIDI_TX,
		RX:       config.PIN_MIDI_RX,
	})
	if err != nil {
		for {
			// loop forever
			time.Sleep(time.Millisecond)
		}
	}

	return uart
}
