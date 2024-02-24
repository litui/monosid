package sid

import (
	"runtime"

	"github.com/litui/monosid/log"
)

func Task() {
	if initGpio() && initClock() {
		log.Logf("SID ready")
	}

	for {
		runtime.Gosched()
	}
}
