package midi

import (
	"machine"
	"runtime"
)

var (
	// Define sysex buffer as global since sysex could require multiple passes
	sysexLen    uint16
	sysexBuffer []byte
	inSysex     bool = false
)

func readNextByte(uart *machine.UART) byte {
	for {
		b, err := uart.ReadByte()
		if err == nil {
			return b
		}

		runtime.Gosched()
	}
}

// Determine what kind of midi message we're looking at and how to handle it
func processBuffer(uart *machine.UART) {
	for {
		b, err := uart.ReadByte()
		if err != nil {
			break
		}

		// Assume we're still receiving sysex until the end byte
		if inSysex {
			if b == MsgSysExEnd {
				inSysex = false
				handleSysex()

			} else {
				sysexBuffer = append(sysexBuffer, b)
				sysexLen++
			}
		} else {
			if b == MsgSysExStart {
				// prepare for receiving sysex message
				sysexLen = 0
				inSysex = true

			} else {
				msgType := b & 0xF0
				channel := b & 0xF
				switch msgType {
				case MsgNoteOff:
					note := readNextByte(uart)
					velocity := readNextByte(uart)
					handleNoteOff(channel, note, velocity)
				case MsgNoteOn:
					note := readNextByte(uart)
					velocity := readNextByte(uart)
					handleNoteOn(channel, note, velocity)
				case MsgPolyAftertouch:
					note := readNextByte(uart)
					pressure := readNextByte(uart)
					handlePolyAftertouch(channel, note, pressure)
				case MsgControlChange:
					cc := readNextByte(uart)
					value := readNextByte(uart)
					handleControlChange(channel, cc, value)
				case MsgPitchBend:
					pbLow := readNextByte(uart)
					pbHigh := readNextByte(uart)
					pitchBend := (int16(pbLow&0x7f) | int16(pbHigh&0x7f)<<7) - 0x2000
					handlePitchBend(channel, pitchBend)
				case MsgProgramChange:
					program := readNextByte(uart)
					handleProgramChange(channel, program)
				case MsgChannelAftertouch:
					pressure := readNextByte(uart)
					handleChannelAftertouch(channel, pressure)
				}
			}
		}

		runtime.Gosched()
	}
}
