package main

import (
	"machine"
	"time"

	"github.com/litui/monosid/config"
)

func i2cInit() *machine.I2C {
	i2c := config.MAIN_I2C
	err := i2c.Configure(machine.I2CConfig{
		Frequency: 800000,
		SCL:       config.PIN_I2C_SCL,
		SDA:       config.PIN_I2C_SDA,
	})
	if err != nil {
		for {
			// loop forever
			time.Sleep(time.Millisecond)
		}
	}

	return i2c
}
