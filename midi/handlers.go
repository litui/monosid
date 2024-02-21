package midi

func handleNoteOff(channel uint8, note uint8, velocity uint8) {

}

func handleNoteOn(channel uint8, note uint8, velocity uint8) {

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
