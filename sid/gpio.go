package sid

import (
	"machine"
	"time"

	"github.com/litui/monosid/config"
)

type chip uint8

const (
	chipLeft chip = iota
	chipRight
)

var (
	clockReady = false
	gpioReady  = false
)

func initGpio() bool {
	if gpioReady {
		return gpioReady
	}

	config.PIN_SID_PWM.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	config.PIN_SID_PWM.Low()

	config.PIN_SID_RW.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	config.PIN_SID_RW.High()

	// Set both CS pins high to start
	for _, v := range config.SID_CS_PINS {

		v.Configure(machine.PinConfig{
			Mode: machine.PinOutput,
		})

		v.High()
	}

	// Set addr pins low to start
	for _, v := range config.SID_ADDR_PINS {

		v.Configure(machine.PinConfig{
			Mode: machine.PinOutput,
		})

		v.Low()
	}

	// Set addr pins high to start
	for _, v := range config.SID_DATA_PINS {

		v.Configure(machine.PinConfig{
			Mode: machine.PinOutput,
		})

		v.Low()
	}

	gpioReady = true
	return gpioReady
}

func initClock() bool {
	if !gpioReady || clockReady {
		return clockReady
	}

	pwm := config.SID_PWM
	pwm.Configure(machine.PWMConfig{
		Period: 1e9 / config.SID_PWM_FREQ,
	})

	ch, err := config.SID_PWM.Channel(config.PIN_SID_PWM)
	if err != nil {
		return false
	}

	// 50% duty cycle
	pwm.Set(ch, pwm.Top()/config.SID_PWM_DUTY)

	clockReady = true
	return clockReady
}

// Enable one CS pin and disable the other. index is 0 or 1
func csOn(index chip) {
	pol := intToBool(uint8(index))
	config.SID_CS_PINS[0].Set(pol)
	config.SID_CS_PINS[1].Set(!pol)
}

// Disable both CS pins (set both to high)
func csClear() {
	config.SID_CS_PINS[0].High()
	config.SID_CS_PINS[1].High()
}

// Sets the SID RW pin low
func enableWrite() {
	config.PIN_SID_RW.Low()
}

// Sets the SID RW pin high
func disableWrite() {
	config.PIN_SID_RW.High()
}

// Populates the SID ADDR register pins
func setAddr(address uint8) bool {
	if !clockReady {
		return false
	}

	for i := 0; i < 5; i++ {
		bitVal := intToBool((address >> i) & 1)
		config.SID_ADDR_PINS[i].Set(bitVal)
	}
	return true
}

// Populates the SID data register pins
func setData(data uint8) bool {
	if !clockReady {
		return false
	}

	for i := 0; i < 8; i++ {
		bitVal := intToBool((data >> i) & 1)
		config.SID_DATA_PINS[i].Set(bitVal)
	}
	return true
}

// Writes data to specified SID chip
func writeReg(chip chip, address uint8, data uint8) bool {
	if !clockReady {
		return false
	}
	setData(data)
	setAddr(address)

	time.Sleep(time.Microsecond * 100)
	enableWrite()
	csOn(chip)
	time.Sleep(time.Microsecond * 100)
	csClear()
	disableWrite()

	setData(0)

	return true
}

func intToBool(value uint8) bool {
	if value == 0 {
		return false
	}
	return true
}
