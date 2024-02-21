package main

import (
	"machine"
	"time"
)

func i2cInit() *machine.I2C {
	i2c := machine.I2C1
	err := i2c.Configure(machine.I2CConfig{
		Frequency: 800000,
		SCL:       machine.GP27,
		SDA:       machine.GP26,
	})
	if err != nil {
		for {
			// loop forever
			time.Sleep(time.Millisecond)
		}
	}

	return i2c
}
