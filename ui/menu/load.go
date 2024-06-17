package menu

import (
	"strconv"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initLoadMenuValues(subEncoder []*rotaryencoder.Device) {
	if selectedPatch == -1 {
		selectedPatch = int(settings.Storage.GetSelectedPatch())
	}
	subEncoder[0].SetValue(int(selectedPatch))
}

func processLoadMenuEncoders(subEncoder []*rotaryencoder.Device) {
	if selectedPatch != subEncoder[0].Value() {
		selectedPatch = subEncoder[0].Value()
	}

	if subEncoder[1].SwitchWasClicked() && selectedPatch != -1 {
		settings.Storage.Load(uint8(selectedPatch))
		sid.SetupAfterLoad()

		selectedPatch = selectedPatch
		SaveLoadComplete = true
	}

	if subEncoder[2].SwitchWasClicked() {
		selectedPatch = -1
		SaveLoadComplete = true
	}
}

func renderLoadMenu(display *ssd1306.Device) {
	writeHeader(display, "Load from Slot")

	write3Box(display, 0, strconv.FormatUint(uint64(selectedPatch)+1, 10))
	write3Box(display, 1, "Y")
	write3Box(display, 2, "N")
}
