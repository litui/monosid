package menu

import (
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

type SaveLoadMenu uint8

const (
	LOAD_MENU SaveLoadMenu = iota
	SAVE_MENU
	SAVE_GENERAL_MENU
	SAVE_LOAD_MENU_LENGTH
)

var (
	currentSaveLoadMenu SaveLoadMenu = LOAD_MENU
	selectedPatch       int          = -1
	SaveLoadComplete    bool         = false
)

func RenderSaveLoadMenu(display *ssd1306.Device, subEncoder []*rotaryencoder.Device) {
	display.ClearBuffer()

	switch currentSaveLoadMenu {
	case SAVE_MENU:
		processSaveMenuEncoders(subEncoder)
		renderSaveMenu(display)
	case LOAD_MENU:
		processLoadMenuEncoders(subEncoder)
		renderLoadMenu(display)
	case SAVE_GENERAL_MENU:
		processSaveGeneralMenuEncoders(subEncoder)
		renderSaveGeneralMenu(display)
	}

	display.Display()
}

func ChangeSaveLoadMenu(encoderZero *rotaryencoder.Device) bool {
	newMenu := SaveLoadMenu(encoderZero.Value())
	if newMenu != currentSaveLoadMenu {
		currentSaveLoadMenu = newMenu
		return true
	}
	return false
}

func SetupEncoderSaveLoadMenuRanges(subEncoder []*rotaryencoder.Device) {
	subEncoder[0].SetRange(0, 127)
	subEncoder[1].SetRange(0, 0)
	subEncoder[2].SetRange(0, 0)

	switch currentSaveLoadMenu {
	case SAVE_MENU:
		initSaveMenuValues(subEncoder)
	case LOAD_MENU:
		initLoadMenuValues(subEncoder)
	case SAVE_GENERAL_MENU:
		initSaveGeneralMenuValues(subEncoder)
	}
}
