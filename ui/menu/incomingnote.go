package menu

import (
	"fmt"

	"github.com/litui/monosid/midi"
	"github.com/litui/monosid/midi/notes"
	"github.com/litui/monosid/settings"
	"tinygo.org/x/drivers/ssd1306"
)

func renderIncomingNoteMenu(display *ssd1306.Device) {
	head := fmt.Sprintf("Patch %d", settings.Storage.GetSelectedPatch()+1)
	writeHeader(display, head)

	for ni, n := range midi.CurrentNote[0] {
		if n != -1 {
			noteName := notes.NoteNames[n+notes.FirstNoteMidiOffset]
			write3Box(display, uint8(ni), noteName)
		}
	}
}
