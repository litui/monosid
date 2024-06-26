package menu

import (
	"strconv"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initSaveMenuValues(subEncoder []*rotaryencoder.Device) {
	if selectedPatch == -1 {
		selectedPatch = int(settings.Storage.GetSelectedPatch())
	}
	subEncoder[0].SetValue(int(selectedPatch))
}

func processSaveMenuEncoders(subEncoder []*rotaryencoder.Device) {
	if selectedPatch != subEncoder[0].Value() {
		selectedPatch = subEncoder[0].Value()
	}

	if subEncoder[1].SwitchWasClicked() && selectedPatch != -1 {
		settings.Storage.Save(uint8(selectedPatch))
		selectedPatch = selectedPatch
		SaveLoadComplete = true
	}

	if subEncoder[2].SwitchWasClicked() {
		selectedPatch = -1
		SaveLoadComplete = true
	}
}

func renderSaveMenu(display *ssd1306.Device) {
	writeHeader(display, "Save to Slot")

	write3Box(display, 0, strconv.FormatUint(uint64(selectedPatch)+1, 10))
	write3Box(display, 1, "Y")
	write3Box(display, 2, "N")
}
