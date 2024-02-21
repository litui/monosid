package midi

// Standard MIDI channel messages
const (
	MsgNoteOff           = 0x80
	MsgNoteOn            = 0x90
	MsgPolyAftertouch    = 0xA0
	MsgControlChange     = 0xB0
	MsgProgramChange     = 0xC0
	MsgChannelAftertouch = 0xD0
	MsgPitchBend         = 0xE0
	MsgSysExStart        = 0xF0
	MsgSysExEnd          = 0xF7
)
