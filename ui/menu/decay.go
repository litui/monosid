package menu

import (
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initDecayMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceDecay(0, voice)

		e.SetValue(int(oldValue))
	}
}

func processDecayMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceDecay(0, voice)

		if int(oldValue) != e.Value() {
			// log.Logf("Old: %d, New: %d", oldValue, e.Value())

			for c := 0; c < 2; c++ {
				voiceObj := sid.SID[c].Voice[voice]
				voiceObj.SetDecay(shared.DecayRate(e.Value()))
				settings.Storage.SetVoiceDecay(shared.SidChip(c), voice, shared.DecayRate(e.Value()))
			}
		}
	}
}

func renderDecayMenu(display *ssd1306.Device) {
	writeHeader(display, "Decay Rate")

	for i := 0; i < 3; i++ {
		decay := settings.Storage.GetVoiceDecay(0, shared.VoiceIndex(i))
		decayText := shared.DecayText[decay]
		write3Box(display, uint8(i), decayText)
	}
}
