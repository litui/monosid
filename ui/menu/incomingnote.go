package menu

import (
	"github.com/litui/monosid/midi"
	"github.com/litui/monosid/midi/notes"
	"tinygo.org/x/drivers/ssd1306"
)

func renderIncomingNoteMenu(display *ssd1306.Device) {
	writeHeader(display, "MIDI Notes")

	for ni, n := range midi.CurrentNote[0] {
		if n != -1 {
			noteName := notes.NoteNames[n+notes.FirstNoteMidiOffset]
			write3Box(display, uint8(ni), noteName)
		}
	}
}
