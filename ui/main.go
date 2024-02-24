package ui

import (
	"image/color"
	"machine"
	"runtime"

	"github.com/litui/monosid/config"
	"github.com/litui/monosid/log"
	"tinygo.org/x/drivers/ssd1306"
)

const (
	HEIGHT = 32
)

type Screen uint8

const (
	SCREEN_LOG Screen = iota
)

var (
	BLACK = color.RGBA{0, 0, 0, 255}
	WHITE = color.RGBA{1, 1, 1, 255}

	currentScreen Screen
	display       ssd1306.Device
)

func Task(i2c *machine.I2C) {
	initEncoders(i2c)

	display = ssd1306.NewI2C(i2c)
	display.Configure(ssd1306.Config{
		Width:   config.DISPLAY_WIDTH,
		Height:  config.DISPLAY_HEIGHT,
		Address: config.DISPLAY_I2C_ADDRESS,
	})

	currentScreen = SCREEN_LOG

	display.ClearDisplay()

	log.Logf("UI ready")

	// tinyfont.WriteLineRotated(&display, &tinyfont.TomThumb, 0, 8, "Test", WHITE, tinyfont.NO_ROTATION)

	for {
		tickEncoders()
		// log.Logf("Enc: %d, %d, %d, %d", Encoder[0].Value(), Encoder[1].Value(), Encoder[2].Value(), Encoder[3].Value())

		// Handle display

		display.ClearBuffer()

		switch currentScreen {
		case SCREEN_LOG:
			renderLog()
		}
		display.Display()

		runtime.Gosched()
	}
}
