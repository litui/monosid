package settings

import (
	"machine"

	"tinygo.org/x/tinyfs"
)

var (
	currentDir = "/"

	blockDev tinyfs.BlockDevice
	fs       tinyfs.Filesystem

	flashDev = machine.Flash
)

func initFs() {
	if err := fs.Mount(); err != nil {
		return
	}
}
