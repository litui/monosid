package config

import "machine"

const (
	PIN_I2C_SDA = machine.GP26
	PIN_I2C_SCL = machine.GP27

	PIN_MIDI_TX = machine.UART0_TX_PIN
	PIN_MIDI_RX = machine.UART0_RX_PIN

	PIN_SID_RW   = machine.GP18
	PIN_SID_PHI2 = machine.GP28

	PIN_SID_PWM  = PIN_SID_PHI2
	SID_PWM_FREQ = 1000000

	// 50% duty cycle; 1.0 / 0.5 = 2
	SID_PWM_DUTY = 2
)

var (
	MAIN_I2C  = machine.I2C1
	MIDI_UART = machine.UART0

	// Pin 28 is PWM6
	SID_PWM = machine.PWM6

	SID_CS_PINS = [2]machine.Pin{
		machine.GP16,
		machine.GP17,
	}
	SID_ADDR_PINS = [5]machine.Pin{
		machine.GP2,
		machine.GP3,
		machine.GP4,
		machine.GP5,
		machine.GP6,
	}
	SID_DATA_PINS = [8]machine.Pin{
		machine.GP8,
		machine.GP9,
		machine.GP10,
		machine.GP11,
		machine.GP12,
		machine.GP13,
		machine.GP14,
		machine.GP15,
	}
)
