package main

import (
	"machine"
	"time"
)

func uartInit() *machine.UART {
	uart := machine.UART0
	err := uart.Configure(machine.UARTConfig{
		BaudRate: 31250,
		TX:       machine.UART0_TX_PIN,
		RX:       machine.UART0_RX_PIN,
	})
	if err != nil {
		for {
			// loop forever
			time.Sleep(time.Millisecond)
		}
	}

	return uart
}
