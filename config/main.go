package config

import "machine"

// MIDI config
const (
	PIN_MIDI_TX = machine.UART0_TX_PIN
	PIN_MIDI_RX = machine.UART0_RX_PIN
)

var (
	MIDI_UART = machine.UART0
)

// Display & I2C config
const (
	PIN_I2C_SDA = machine.GP26
	PIN_I2C_SCL = machine.GP27

	DISPLAY_WIDTH       = 128
	DISPLAY_HEIGHT      = 32
	DISPLAY_I2C_ADDRESS = 0x3c

	EXPANDER_I2C_ADDRESS = 0x20

	// The following are pins on the i2c expander
	ENCODER1_PIN1   = 3
	ENCODER1_PIN2   = 4
	ENCODER1_BUTTON = 5
	ENCODER2_PIN1   = 0
	ENCODER2_PIN2   = 1
	ENCODER2_BUTTON = 2
	ENCODER3_PIN1   = 11
	ENCODER3_PIN2   = 12
	ENCODER3_BUTTON = 13
	ENCODER4_PIN1   = 8
	ENCODER4_PIN2   = 9
	ENCODER4_BUTTON = 10
)

var (
	MAIN_I2C = machine.I2C1

	ENCODER_OFFSETS = [4]int{ENCODER1_PIN1, ENCODER2_PIN1, ENCODER3_PIN1, ENCODER4_PIN1}
)

// SID config
const (
	PIN_SID_RW   = machine.GP18
	PIN_SID_PHI2 = machine.GP28

	PIN_SID_PWM  = PIN_SID_PHI2
	SID_PWM_FREQ = 1000000

	// 50% duty cycle; 1.0 / 0.5 = 2
	SID_PWM_DUTY = 2
)

var (
	// Pin 28 is PWM6
	SID_PWM = machine.PWM6

	SID_CS_PINS = [2]machine.Pin{
		machine.GP17,
		machine.GP16,
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
