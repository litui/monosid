package main

import (
	"time"

	"github.com/litui/monosid/graphics"
	"github.com/litui/monosid/midi"
	"github.com/litui/monosid/sid"
)

func main() {
	i2c := i2cInit()
	uart := uartInit()

	go graphics.Task(i2c)
	go midi.Task(uart)
	go sid.Task()

	for {
		time.Sleep(time.Second)
	}
}
