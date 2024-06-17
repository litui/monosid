package menu

import (
	"strconv"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initVolumeMenuValues(subEncoder []*rotaryencoder.Device) {
	vol := settings.Storage.GetVolume(shared.SidChip(0))
	subEncoder[0].SetValue(int(vol))
}

func processVolumeMenuEncoders(subEncoder []*rotaryencoder.Device) {
	vol := settings.Storage.GetVolume(shared.SidChip(0))

	if vol != uint8(subEncoder[0].Value()) {
		for c := 0; c < 2; c++ {
			ci := shared.SidChip(c)
			settings.Storage.SetVolume(ci, uint8(subEncoder[0].Value()))
			sid.SID[c].SetVolume(uint8(subEncoder[0].Value()))
		}
	}
}

func renderVolumeMenu(display *ssd1306.Device) {
	writeHeader(display, "Main Volume")

	vol := settings.Storage.GetVolume(shared.SidChip(0))
	write3Box(display, 0, strconv.FormatUint(uint64(vol), 10))
}
