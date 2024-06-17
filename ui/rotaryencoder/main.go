package rotaryencoder

// Borrowed from tinygo drivers and adjusted to use the MCP23017

import (
	"math"
	"time"

	"tinygo.org/x/drivers/mcp23017"
)

const (
	clickSum = 4
)

var (
	states           = []int8{0, -1, 1, 0, 1, 0, 0, -1, -1, 0, 0, 1, 0, 1, -1, 0}
	lastClick        time.Time
	debounceInterval time.Duration = 200 * time.Millisecond
)

// New creates a new rotary encoder.
func New(pinA, pinB, pinS mcp23017.Pin) Device {
	return Device{pinA: pinA, pinB: pinB, pinS: pinS,
		oldAB: 0b00000011, value: 0,
		swValue: false, oldSwValue: false, wasClicked: false,
		Dir:         make(chan int, 8),
		Switch:      make(chan bool),
		rangeBottom: math.MinInt,
		rangeTop:    math.MaxInt,
	}
}

// Device represents a rotary encoder.
type Device struct {
	pinA mcp23017.Pin // CLK pin
	pinB mcp23017.Pin // DT pin
	pinS mcp23017.Pin // SW pin

	oldAB      int
	value      int
	swValue    bool
	oldSwValue bool
	wasClicked bool
	Dir        chan int
	Switch     chan bool

	rangeBottom int
	rangeTop    int
}

// Configure configures the rotary encoder.
func (enc *Device) Configure() {
	enc.pinA.SetMode(mcp23017.Input)
	enc.pinB.SetMode(mcp23017.Input)
	enc.pinS.SetMode(mcp23017.Input)
}

func (enc *Device) UpdateSW(pinState bool) {
	if pinState { // the switch is released -- because of pullup
		if time.Since(lastClick) > debounceInterval {
			lastClick = time.Now()
			enc.swValue = false
			enc.wasClicked = true
			select {
			case enc.Switch <- true:
			default:
			}
		}
	} else { //the switch is pressed
		enc.swValue = true
	}
}

func (enc *Device) Update(aHigh bool, bHigh bool) {
	enc.oldAB <<= 2
	if aHigh {
		enc.oldAB |= 1
	}
	if bHigh {
		enc.oldAB |= 1 << 1
	}

	enc.value += int(states[enc.oldAB&0x0f])
	// enc.value = 0

	// Each full click of the encoder generates 4 interupts.
	// Each interrupt adds 1 or -1 to the value.
	// We send the direction to the channel only when we have a full click, i.e. 4 interrupts.

	if enc.value%clickSum == 0 {
		divVal := enc.value / clickSum

		direction := -(divVal)
		if direction != 0 {
			select { // non-blocking way of sending to channel
			case enc.Dir <- direction:
			default:
			}
		}

		// Limit encoder to within defined bounds
		if divVal < enc.rangeBottom {
			enc.value = enc.rangeBottom * clickSum
		} else if divVal > enc.rangeTop {
			enc.value = enc.rangeTop * clickSum
		}
	}
}

// Value returns the value of the rotary encoder.
func (enc *Device) Value() int {
	return enc.value / clickSum
}

// SetValue sets the value of the rotary encoder.
func (enc *Device) SetValue(v int) {
	enc.value = int(v * clickSum)
}

// SwitchValue returns the value of the switch.
func (enc *Device) SwitchWasClicked() bool {
	if enc.wasClicked {
		enc.wasClicked = false
		return true
	} else {
		return false
	}
}

// Returns the current defined range for the encoder
func (enc *Device) Range() (int, int) {
	return enc.rangeBottom, enc.rangeTop
}

// Sets the possible value range for the encoder
// This won't be used until the next tick
func (enc *Device) SetRange(bottom int, top int) {
	enc.rangeBottom = bottom
	enc.rangeTop = top

	divVal := enc.value / clickSum

	if divVal < enc.rangeBottom {
		enc.value = enc.rangeBottom * clickSum
	} else if divVal > enc.rangeTop {
		enc.value = enc.rangeTop * clickSum
	}
}
