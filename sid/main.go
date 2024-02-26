package sid

import (
	"runtime"

	"github.com/litui/monosid/sid/chip"
	"github.com/litui/monosid/sid/gpio"
)

var (
	SID = [2]chip.SIDDevice{chip.New(0), chip.New(1)}

	ready bool = false
)

func Task() {
	gpio.Init()

	// Sensible audio defaults until we get settings in
	for _, s := range SID {
		s.SetVolume(4)
		s.SetFilterCutoff(1024)
		s.SetFilterRes(0)
		for vi, v := range s.Voice {
			s.SetFilterEn(uint8(vi), false)
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
