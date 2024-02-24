package menu

import (
	"github.com/litui/monosid/log"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
)

var (
	lastLogLine int
)

func renderLogMenu(display *ssd1306.Device) {
	j := 0
	for i, v := range log.LogLines {
		if i >= len(log.LogLines)-log.VisibleLogLines {
			tinyfont.WriteLineRotated(display, &tinyfont.TomThumb, 0, int16(8*j)+8, v, WHITE, tinyfont.NO_ROTATION)
			j++
		}
	}
}
