package menu

import (
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
)

type Waveform uint8

const (
	NONE                          Waveform = iota // 0b0000
	TRIANGLE                                      // 0b0001
	SAWTOOTH                                      // 0b0010
	SAWTOOTH_TRIANGLE                             // 0b0011
	PULSE                                         // 0b0100
	PULSE_TRIANGLE                                // 0b0101
	PULSE_SAWTOOTH                                // 0b0110
	PULSE_SAWTOOTH_TRIANGLE                       // 0b0111
	NOISE                                         // 0b1000
	NOISE_TRIANGLE                                // 0b1001
	NOISE_SAWTOOTH                                // 0b1010
	NOISE_SAWTOOTH_TRIANGLE                       // 0b1011
	NOISE_PULSE                                   // 0b1100
	NOISE_PULSE_TRIANGLE                          // 0b1101
	NOISE_PULSE_SAWTOOTH                          // 0b1110
	NOISE_PULSE_SAWTOOTH_TRIANGLE                 // 0b1111
)

var (
	WaveformNames = []string{
		"", "T", "S", "ST", "P", "PT", "PS", "PST", "N", "NT", "NS", "NST", "NP", "NPT", "NPS", "All",
	}
)

func getOldWaveformValue(voice shared.VoiceIndex) Waveform {
	oldTri := shared.BToI(settings.Storage.GetVoiceTriangle(0, voice))
	oldSaw := shared.BToI(settings.Storage.GetVoiceSawtooth(0, voice))
	oldPulse := shared.BToI(settings.Storage.GetVoicePulse(0, voice))
	oldNoise := shared.BToI(settings.Storage.GetVoiceNoise(0, voice))
	return Waveform((oldNoise << 3) | (oldPulse << 2) | (oldSaw << 1) | oldTri)
}

func initWaveformMenuValues(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := getOldWaveformValue(voice)

		e.SetValue(int(oldValue))
	}
}

func processWaveformMenuEncoders(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		voice := shared.VoiceIndex(ei)
		oldValue := getOldWaveformValue(voice)

		if int(oldValue) != e.Value() && e.Value() > 0 {
			newValue := uint8(e.Value())
			for c := 0; c < 2; c++ {
				settings.Storage.SetVoiceTriangle(shared.SidChip(c), voice, shared.IToB(newValue&1))
				settings.Storage.SetVoiceSawtooth(shared.SidChip(c), voice, shared.IToB((newValue>>1)&1))
				settings.Storage.SetVoicePulse(shared.SidChip(c), voice, shared.IToB((newValue>>2)&1))
				settings.Storage.SetVoiceNoise(shared.SidChip(c), voice, shared.IToB((newValue>>3)&1))
			}
		}
	}
}

func renderWaveformMenu(display *ssd1306.Device) {
	writeHeader(display, "Waveform")

	for i := 0; i < 3; i++ {
		waveform := getOldWaveformValue(shared.VoiceIndex(i))
		wName := WaveformNames[waveform]
		write3Box(display, uint8(i), wName)
	}
}
