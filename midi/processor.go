package midi

import (
	"machine"
	"runtime"

	"github.com/lunux2008/xulu"
)

var (
	// Define sysex buffer as global since sysex could require multiple passes
	sysexLen    uint16
	sysexBuffer []byte
	inSysex     bool = false
)

// Determine what kind of midi message we're looking at and how to handle it
func processBuffer(uart *machine.UART) {
	bc := uart.Buffered()
	if bc < 1 {
		return
	}

	var tempBuf []byte
	len, err := uart.Read(tempBuf)
	if err != nil {
		return
	}
	xulu.Use(len)

	// continue receiving sysex message if mid-message
	for i := 0; i < len; i++ {
		nextByte := tempBuf[i]

		// Assume we're still receiving sysex until the end byte
		if inSysex {
			if nextByte == MsgSysExEnd {
				inSysex = false
				handleSysex()

			} else {
				sysexBuffer = append(sysexBuffer, nextByte)
				sysexLen++
			}
		} else {
			if nextByte == MsgSysExStart {
				// prepare for receiving sysex message
				sysexLen = 0
				inSysex = true

			} else {
				msgType := nextByte & 0xF0
				channel := nextByte & 0xF
				if i+2 < len {
					switch msgType {
					case MsgNoteOff:
						handleNoteOff(channel, tempBuf[i+1], tempBuf[i+2])
					case MsgNoteOn:
						handleNoteOn(channel, tempBuf[i+1], tempBuf[i+2])
					case MsgPolyAftertouch:
						handlePolyAftertouch(channel, tempBuf[i+1], tempBuf[i+2])
					case MsgControlChange:
						handleControlChange(channel, tempBuf[i+1], tempBuf[i+2])
					case MsgPitchBend:
						pitchBend := (int16(tempBuf[i+1]&0x7f) | int16(tempBuf[i+1]&0x7f)<<7) - 0x2000
						handlePitchBend(channel, pitchBend)
					}
				}
				if i+1 < len {
					switch msgType {
					case MsgProgramChange:
						handleProgramChange(channel, tempBuf[i+1])
					case MsgChannelAftertouch:
						handleChannelAftertouch(channel, tempBuf[i+1])
					}
				}
			}
		}
		runtime.Gosched()
	}
}
