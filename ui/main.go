package ui

import (
	"machine"
	"runtime"

	"github.com/litui/monosid/config"
	"github.com/litui/monosid/log"
	"github.com/litui/monosid/ui/menu"
	"tinygo.org/x/drivers/ssd1306"
)

const (
	HEIGHT = 32
)

var (
	display ssd1306.Device
)

func Task(i2c *machine.I2C) {
	initEncoders(i2c)

	display = ssd1306.NewI2C(i2c)
	display.Configure(ssd1306.Config{
		Width:   config.DISPLAY_WIDTH,
		Height:  config.DISPLAY_HEIGHT,
		Address: config.DISPLAY_I2C_ADDRESS,
	})

	display.ClearDisplay()

	log.Logf("UI ready")

	// tinyfont.WriteLineRotated(&display, &tinyfont.TomThumb, 0, 8, "Test", WHITE, tinyfont.NO_ROTATION)

	for {
		tickEncoders()

		menu.RenderMainMenu(&display, Encoder)

		runtime.Gosched()
	}
}
