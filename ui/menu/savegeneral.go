package menu

import (
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initSaveGeneralMenuValues(subEncoder []*rotaryencoder.Device) {
	// if selectedPatch == -1 {
	// 	selectedPatch = int(settings.Storage.GetSelectedPatch())
	// }
	// subEncoder[0].SetValue(int(selectedPatch))
}

func processSaveGeneralMenuEncoders(subEncoder []*rotaryencoder.Device) {
	// if selectedPatch != subEncoder[0].Value() {
	// 	selectedPatch = subEncoder[0].Value()
	// }

	if subEncoder[1].SwitchWasClicked() && selectedPatch != -1 {
		settings.Storage.SaveGeneral()
		SaveLoadComplete = true
	}

	if subEncoder[2].SwitchWasClicked() {
		selectedPatch = -1
		SaveLoadComplete = true
	}
}

func renderSaveGeneralMenu(display *ssd1306.Device) {
	writeHeader(display, "Save General")

	write3Box(display, 0, "")
	write3Box(display, 1, "Y")
	write3Box(display, 2, "N")
}
