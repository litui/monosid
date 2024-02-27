package menu

import (
	"strconv"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initChannelMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetMidiChannel(0, voice)

		e.SetValue(int(oldValue))
	}
}

func processChannelMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetMidiChannel(0, voice)

		if int(oldValue) != e.Value() {
			// log.Logf("Old: %d, New: %d", oldValue, e.Value())

			for c := 0; c < 2; c++ {
				settings.Storage.SetMidiChannel(shared.SidChip(c), voice, uint8(e.Value()))
			}
		}
	}
}

func renderChannelMenu(display *ssd1306.Device) {
	writeHeader(display, "MIDI Channel")

	for i := 0; i < 3; i++ {
		channel := uint64(settings.Storage.GetMidiChannel(0, shared.VoiceIndex(i)) + 1)
		write3Box(display, uint8(i), strconv.FormatUint(channel, 10))
	}
}
