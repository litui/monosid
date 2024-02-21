package graphics

import (
	"image/color"
	"machine"
	"runtime"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	BLACK = color.RGBA{0, 0, 0, 255}
	WHITE = color.RGBA{1, 1, 1, 255}
)

var display ssd1306.Device

func Task(i2c *machine.I2C) {
	display = ssd1306.NewI2C(i2c)
	display.Configure(ssd1306.Config{
		Width:   128,
		Height:  32,
		Address: 0x3c,
	})

	display.ClearBuffer()
	display.ClearDisplay()

	tinyfont.WriteLineRotated(&display, &freesans.Bold9pt7b, 0, 14, "Test", WHITE, tinyfont.NO_ROTATION)

	display.Display()

	for {
		runtime.Gosched()
	}
}
