package menu

import (
	"strconv"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initDetuneMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceDetune(0, voice)

		e.SetValue(int(oldValue))
	}
}

func processDetuneMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceDetune(0, voice)

		if int(oldValue) != e.Value() {
			for c := 0; c < 2; c++ {
				settings.Storage.SetVoiceDetune(shared.SidChip(c), voice, int8(e.Value()))
			}
		}
	}
}

func renderDetuneMenu(display *ssd1306.Device) {
	writeHeader(display, "Detune (cents)")

	for i := 0; i < 3; i++ {
		channel := int64(settings.Storage.GetVoiceDetune(0, shared.VoiceIndex(i)))
		write3Box(display, uint8(i), strconv.FormatInt(channel, 10))
	}
}
