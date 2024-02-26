package voice

import (
	"math"

	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid/gpio"
)

const (
	SID_REG_VOICE_FREQ_LO = 0
	SID_REG_VOICE_FREQ_HI = 1
	SID_REG_VOICE_PW_LO   = 2
	SID_REG_VOICE_PW_HI   = 3
	SID_REG_VOICE_CTRL    = 4
	SID_REG_VOICE_AD      = 5
	SID_REG_VOICE_SR      = 6
)

type Voice struct {
	index uint8
	chip  shared.SidChip

	// Pitch & Timbre
	freq      float32
	dutyCycle float32

	// Note options
	gate     bool
	sync     bool
	ringMod  bool
	test     bool
	triangle bool
	sawtooth bool
	pulse    bool
	noise    bool

	// Envelope
	attack  shared.AttackRate
	decay   shared.DecayRate
	sustain uint8
	release shared.ReleaseRate
}

func New(index uint8, chip shared.SidChip) Voice {
	return Voice{
		index:     index,
		chip:      chip,
		freq:      0,
		dutyCycle: 0.5,
		gate:      false,
		sync:      false,
		ringMod:   false,
		test:      false,
		triangle:  false,
		sawtooth:  true,
		pulse:     false,
		noise:     false,
		attack:    0,
		decay:     0,
		sustain:   6,
		release:   8,
	}
}

func (v *Voice) SetFrequency(freq float32) {
	v.freq = freq

	regFreqLo := uint8(7*v.index + SID_REG_VOICE_FREQ_LO)
	regFreqHi := uint8(7*v.index + SID_REG_VOICE_FREQ_HI)

	realFreq := uint16(math.Round(float64(freq * 16.777216)))

	gpio.WriteReg(v.chip, regFreqHi, uint8(realFreq>>8))
	gpio.WriteReg(v.chip, regFreqLo, uint8(realFreq&0xff))
}

func (v *Voice) SetPulseWidth(dutyCycle float32) {
	v.dutyCycle = dutyCycle

	pulseWidth := uint16(math.Round(float64(dutyCycle*4096 - 1)))

	regPWLo := uint8(7*v.index + SID_REG_VOICE_PW_LO)
	regPWHi := uint8(7*v.index + SID_REG_VOICE_PW_HI)

	gpio.WriteReg(v.chip, regPWHi, uint8(pulseWidth>>8)&0xf)
	gpio.WriteReg(v.chip, regPWLo, uint8(pulseWidth&0xff))
}

func (v *Voice) SetEnvelope(attack shared.AttackRate, decay shared.DecayRate, sustain uint8, release shared.ReleaseRate) {
	v.attack = attack & 0xf
	v.decay = decay & 0xf
	v.sustain = sustain & 0xf
	v.release = release & 0xf

	regAD := uint8(7*v.index + SID_REG_VOICE_AD)
	regSR := uint8(7*v.index + SID_REG_VOICE_SR)

	adBits := uint8(v.attack)<<4 | uint8(v.decay)
	srBits := uint8(v.sustain)<<4 | uint8(v.release)

	gpio.WriteReg(v.chip, regAD, adBits)
	gpio.WriteReg(v.chip, regSR, srBits)
}

func (v *Voice) SetAttack(attack shared.AttackRate) {
	v.SetEnvelope(attack, v.decay, v.sustain, v.release)
}

func (v *Voice) SetDecay(decay shared.DecayRate) {
	v.SetEnvelope(v.attack, decay, v.sustain, v.release)
}

func (v *Voice) SetSustain(sustain uint8) {
	v.SetEnvelope(v.attack, v.decay, sustain, v.release)
}

func (v *Voice) SetRelease(release shared.ReleaseRate) {
	v.SetEnvelope(v.attack, v.decay, v.sustain, release)
}

func (v *Voice) SetWaveform(triangle bool, sawtooth bool, pulse bool, noise bool) {
	v.triangle = triangle
	v.sawtooth = sawtooth
	v.pulse = pulse
	v.noise = noise
}

func (v *Voice) Trigger() {
	regCtrl := uint8(7*v.index + SID_REG_VOICE_CTRL)

	v.gate = true

	outBits := uint8(shared.BToI(v.gate))
	outBits |= shared.BToI(v.sync) << 1
	outBits |= shared.BToI(v.ringMod) << 2
	outBits |= shared.BToI(v.test) << 3
	outBits |= shared.BToI(v.triangle) << 4
	outBits |= shared.BToI(v.sawtooth) << 5
	outBits |= shared.BToI(v.pulse) << 6
	outBits |= shared.BToI(v.noise) << 7

	gpio.WriteReg(v.chip, regCtrl, outBits)
}

func (v *Voice) Release() {
	regCtrl := uint8(7*v.index + SID_REG_VOICE_CTRL)

	v.gate = false

	outBits := uint8(shared.BToI(v.gate))
	outBits |= shared.BToI(v.sync) << 1
	outBits |= shared.BToI(v.ringMod) << 2
	outBits |= shared.BToI(v.test) << 3
	outBits |= shared.BToI(v.triangle) << 4
	outBits |= shared.BToI(v.sawtooth) << 5
	outBits |= shared.BToI(v.pulse) << 6
	outBits |= shared.BToI(v.noise) << 7

	gpio.WriteReg(v.chip, regCtrl, outBits)
}
