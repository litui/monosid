package main

import (
	"time"

	"github.com/litui/monosid/led"
	"github.com/litui/monosid/log"
	"github.com/litui/monosid/midi"
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui"
)

func main() {
	log.InitLog()

	i2c := i2cInit()
	uart := uartInit()

	go led.Task()
	go ui.Task(i2c)

	go settings.Task()

	go sid.Task()
	go midi.Task(uart)

	sid.SetupAfterLoad()

	for {
		time.Sleep(time.Second)
	}
}
