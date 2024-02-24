package ui

import (
	"github.com/litui/monosid/log"
	"tinygo.org/x/tinyfont"
)

var (
	lastLogLine int
)

func renderLog() {
	j := 0
	for i, v := range log.LogLines {
		if i >= len(log.LogLines)-log.VisibleLogLines {
			tinyfont.WriteLineRotated(&display, &tinyfont.TomThumb, 0, int16(8*j)+8, v, WHITE, tinyfont.NO_ROTATION)
			j++
		}
	}
}
