package midi

import (
	"math"

	"github.com/litui/monosid/led"
	"github.com/litui/monosid/log"
	"github.com/litui/monosid/midi/notes"
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid"
)

var (
	CurrentNote = [2][3]int8{
		{-1, -1, -1},
		{-1, -1, -1},
	}
)

func handleNoteOff(channel uint8, note uint8, velocity uint8) {
	led.Flash()
	log.Logf("NoteOff - Ch%d - %d - %d", channel, note, velocity)

	for si, s := range sid.SID {
		for vi, v := range s.Voice {
			if settings.Storage.GetMidiChannel(shared.SidChip(si), shared.VoiceIndex(vi)) != channel {
				continue
			}

			v.Release()
			CurrentNote[si][vi] = -1 // off
		}
	}
}

func handleNoteOn(channel uint8, note uint8, velocity uint8) {
	led.Flash()
	log.Logf("NoteOn - Ch%d - %d - %d", channel, note, velocity)

	if int8(note)+notes.FirstNoteMidiOffset >= int8(len(notes.NotePitches)) {
		return
	} else if int8(note) < notes.FirstNoteMidiOffset {
		return
	}

	baseFreq := notes.NotePitches[int8(note)+notes.FirstNoteMidiOffset]

	for si, s := range sid.SID {
		for vi, v := range s.Voice {
			chip := shared.SidChip(si)
			voice := shared.VoiceIndex(vi)

			if settings.Storage.GetMidiChannel(chip, voice) != channel {
				continue
			}

			cents := settings.Storage.GetVoiceDetune(chip, voice)
			adjFreq := baseFreq

			if cents != 0 {
				adjFreq = baseFreq * float32(math.Pow(2.0, float64(cents)/1200.0))
			}

			v.SetFrequency(adjFreq)
			v.Trigger()
			CurrentNote[si][vi] = int8(note)
		}
	}
}

func handlePolyAftertouch(channel uint8, note uint8, pressure uint8) {

}

func handleControlChange(channel uint8, cc uint8, value uint8) {

}

func handleProgramChange(channel uint8, program uint8) {

}

func handleChannelAftertouch(channel uint8, pressure uint8) {

}

// Pitchbend represented as range from -8192 to 8191 with 0 as off
func handlePitchBend(channel uint8, pitchBend int16) {

}

func handleSysex() {

	// When handled, reset the length
	sysexLen = 0
}
