package sid

import (
	"runtime"

	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid/chip"
	"github.com/litui/monosid/sid/gpio"
)

var (
	SID = [2]chip.SIDDevice{chip.New(0), chip.New(1)}

	ready bool = false
)

func SetupAfterLoad() {
	for si, s := range SID {
		c := shared.SidChip(si)

		s.SetVolume(settings.Storage.GetVolume(c))
		s.SetFilterMode(
			settings.Storage.GetFilterLP(c),
			settings.Storage.GetFilterBP(c),
			settings.Storage.GetFilterHP(c),
		)
		s.SetFilterCutoff(settings.Storage.GetFilterCutoff(c))
		s.SetFilterRes(settings.Storage.GetFilterRes(c))

		for vcnt, v := range s.Voice {
			vi := shared.VoiceIndex(vcnt)

			s.SetFilterEn(vi, settings.Storage.GetFilterEn(c, vi))
			v.SetEnvelope(
				settings.Storage.GetVoiceAttack(c, vi),
				settings.Storage.GetVoiceDecay(c, vi),
				settings.Storage.GetVoiceSustain(c, vi),
				settings.Storage.GetVoiceRelease(c, vi),
			)
			v.SetRawPulseWidth(settings.Storage.GetVoicePW(c, vi))

			// Don't need to set waveform and related settings which are loaded dynamically
		}
	}
}

func Task() {
	gpio.Init()

	// Sensible audio defaults until we get settings in
	for _, s := range SID {
		s.SetVolume(4)
		s.SetFilterCutoff(1024)
		s.SetFilterRes(0)
		for vi, v := range s.Voice {
			s.SetFilterEn(shared.VoiceIndex(vi), false)
			v.SetEnvelope(1, 0, 6, 8)
			v.SetWaveform(false, true, false, false)
			v.SetPulseWidth(0.5)
		}
	}

	ready = true

	for {
		runtime.Gosched()
	}
}

func IsReady() bool {
	return ready
}
