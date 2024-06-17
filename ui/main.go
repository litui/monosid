package ui

import (
	"machine"
	"runtime"
	"time"

	"github.com/litui/monosid/config"
	"github.com/litui/monosid/ui/menu"
	"tinygo.org/x/drivers/ssd1306"
)

const (
	HEIGHT = 32
)

var (
	display           ssd1306.Device
	lastDisplayUpdate = time.Time{}
	saveLoadMode      = false
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

	for {
		tickEncoders()
		currentTime := time.Now()

		if Encoder[0].SwitchWasClicked() {
			saveLoadMode = true
			Encoder[0].SetValue(0)
		}

		if saveLoadMode {
			if menu.ChangeSaveLoadMenu(Encoder[0]) {
				menu.SetupEncoderSaveLoadMenuRanges(Encoder[1:])
			}

			if currentTime.Compare(lastDisplayUpdate.Add(time.Millisecond*20)) > 0 {
				lastDisplayUpdate = currentTime
				menu.RenderSaveLoadMenu(&display, Encoder[1:])
			}

			if menu.SaveLoadComplete {
				menu.SaveLoadComplete = false
				saveLoadMode = false
				Encoder[0].SetValue(0)
			}
		} else {
			// Change menu - returns true if changed
			if menu.ChangeMainMenu(Encoder[0]) {
				menu.SetupEncoderMenuRanges(Encoder[1:])
			}

			if currentTime.Compare(lastDisplayUpdate.Add(time.Millisecond*20)) > 0 {
				lastDisplayUpdate = currentTime
				menu.RenderMainMenu(&display, Encoder[1:])
			}
		}

		runtime.Gosched()
	}
}
