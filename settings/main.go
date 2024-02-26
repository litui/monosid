package settings

import (
	"runtime"

	"github.com/litui/monosid/settings/storage"
)

var (
	Storage storage.StorageDevice
)

func Task() {
	Storage.Init()

	for {
		Storage.Tick()
		runtime.Gosched()
	}
}
