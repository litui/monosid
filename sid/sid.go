package sid

import "runtime"

func Task() {
	if initGpio() {
		initClock()
	}

	for {
		runtime.Gosched()
	}
}
