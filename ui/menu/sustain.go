package menu

import (
	"strconv"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initSustainMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceSustain(0, voice)

		e.SetValue(int(oldValue))
	}
}

func processSustainMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceSustain(0, voice)

		if int(oldValue) != e.Value() {
			// log.Logf("Old: %d, New: %d", oldValue, e.Value())

			for c := 0; c < 2; c++ {
				voiceObj := sid.SID[c].Voice[voice]
				voiceObj.SetSustain(uint8(e.Value()))
				settings.Storage.SetVoiceSustain(shared.SidChip(c), voice, uint8(e.Value()))
			}
		}
	}
}

func renderSustainMenu(display *ssd1306.Device) {
	writeHeader(display, "Sustain Level (/15)")

	for i := 0; i < 3; i++ {
		sustain := uint64(settings.Storage.GetVoiceSustain(0, shared.VoiceIndex(i)))
		write3Box(display, uint8(i), strconv.FormatUint(sustain, 10))
	}
}
