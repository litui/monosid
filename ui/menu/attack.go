package menu

import (
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

func initAttackMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceAttack(0, voice)

		e.SetValue(int(oldValue))
	}
}

func processAttackMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := settings.Storage.GetVoiceAttack(0, voice)

		if int(oldValue) != e.Value() {
			// log.Logf("Old: %d, New: %d", oldValue, e.Value())

			for c := 0; c < 2; c++ {
				voiceObj := sid.SID[c].Voice[voice]
				voiceObj.SetAttack(shared.AttackRate(e.Value()))
				settings.Storage.SetVoiceAttack(shared.SidChip(c), voice, shared.AttackRate(e.Value()))
			}
		}
	}
}

func renderAttackMenu(display *ssd1306.Device) {
	writeHeader(display, "Attack Rate")

	for i := 0; i < 3; i++ {
		attack := settings.Storage.GetVoiceAttack(0, shared.VoiceIndex(i))
		attackText := shared.AttackText[attack]
		write3Box(display, uint8(i), attackText)
	}
}
