package menu

import (
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initReleaseMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceRelease(0, voice)

		e.SetValue(int(oldValue))
	}
}

func processReleaseMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceRelease(0, voice)

		if int(oldValue) != e.Value() {
			// log.Logf("Old: %d, New: %d", oldValue, e.Value())

			for c := 0; c < 2; c++ {
				voiceObj := sid.SID[c].Voice[voice]
				voiceObj.SetRelease(shared.ReleaseRate(e.Value()))
				settings.Storage.SetVoiceRelease(shared.SidChip(c), voice, shared.ReleaseRate(e.Value()))
			}
		}
	}
}

func renderReleaseMenu(display *ssd1306.Device) {
	writeHeader(display, "Release Rate")

	for i := 0; i < 3; i++ {
		release := settings.Storage.GetVoiceRelease(0, shared.VoiceIndex(i))
		releaseText := shared.ReleaseText[release]
		write3Box(display, uint8(i), releaseText)
	}
}
