package settings

import "runtime"

func Task() {
	initFs()

	for {
		runtime.Gosched()
	}
}
