package sid

import (
	"runtime"
)

var (
	SID = [2]SIDDevice{NewSID(0), NewSID(1)}
)

func Task() {
	if initGpio() && initClock() {
		// log.Logf("SID ready")
	}

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

	for {
		runtime.Gosched()
	}
}

func IsReady() bool {
	return clockReady
}
