package ui

import (
	"machine"

	"github.com/litui/monosid/config"
	"github.com/litui/monosid/ui/menu"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/mcp23017"
)

var (
	expander *mcp23017.Device
	Encoder  []*rotaryencoder.Device
	encReady bool = true

	pinCache mcp23017.Pins
)

func initEncoders(i2c *machine.I2C) {
	// Expander / encoder setup
	expander, _ = mcp23017.NewI2C(i2c, config.EXPANDER_I2C_ADDRESS)

	Encoder = make([]*rotaryencoder.Device, 0)

	for _, pos := range config.ENCODER_OFFSETS {
		pinA := expander.Pin(pos)
		pinB := expander.Pin(pos + 1)
		pinS := expander.Pin(pos + 2)

		r := rotaryencoder.New(pinA, pinB, pinS)

		r.Configure()
		Encoder = append(Encoder, &r)
	}

	// Lock in the range for the main menu encoder
	Encoder[0].SetRange(0, int(menu.MENU_LENGTH)-1)

	pinCache, _ = expander.GetPins()
	encReady = true
}

// Since we can't use interrupts with the expander, here's a tick function
func tickEncoders() {
	if !encReady {
		return
	}

	pins, _ := expander.GetPins()

	for i, enc := range Encoder {
		offset := config.ENCODER_OFFSETS[i]
		valA := pins.Get(offset)
		cValA := pinCache.Get(offset)
		valB := pins.Get(offset + 1)
		cValB := pinCache.Get(offset + 1)

		if valA != cValA || valB != cValB {
			enc.Update(valA, valB)
			// log.Logf("Val: %d", enc.Value())
		}

		valS := pins.Get(offset + 2)
		cValS := pinCache.Get(offset + 2)
		if valS && valS != cValS {
			enc.UpdateSW(valS)
			// log.Logf("Sw: %t", enc.SwitchWasClicked())
		}
	}

	pinCache = pins
}
