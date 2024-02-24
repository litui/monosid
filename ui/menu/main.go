package menu

import (
	"image/color"

	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

type Menu int

const (
	DEBUG_LOG Menu = iota
	// INCOMING_NOTE
	// CHANNELS
	// PITCH
	// WAVEFORM
	// PULSEWIDTH
	// ATTACK
	// DECAY
	// SUSTAIN
	// RELEASE
	MENU_LENGTH
)

var (
	BLACK = color.RGBA{0, 0, 0, 255}
	WHITE = color.RGBA{1, 1, 1, 255}

	currentMenu Menu
)

func RenderMainMenu(display *ssd1306.Device, encoder []*rotaryencoder.Device) {
	currentMenu = Menu(encoder[0].Value())

	display.ClearBuffer()

	switch currentMenu {
	case DEBUG_LOG:
		renderLogMenu(display)
	}

	display.Display()
}
