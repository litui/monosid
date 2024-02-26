package settings

import (
	"runtime"
)

var (
	Settings StorageDevice
)

func Task() {
	Settings.Init()

	for {
		Settings.Tick()
		runtime.Gosched()
	}
}
