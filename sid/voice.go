package sid

import "math"

type Voice struct {
	index uint8
	chip  chip

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
	attack  uint8
	decay   uint8
	sustain uint8
	release uint8
}

// const (
// 	ATTACK_2MS    = 0
// 	ATTACK_8MS    = 1
// 	ATTACK_16MS   = 2
// 	ATTACK_24MS   = 3
// 	ATTACK_38MS   = 4
// 	ATTACK_56MS   = 5
// 	ATTACK_68MS   = 6
// 	ATTACK_80MS   = 7
// 	ATTACK_100MS  = 8
// 	ATTACK_250MS  = 9
// 	ATTACK_500MS  = 10
// 	ATTACK_800MS  = 11
// 	ATTACK_1000MS = 12
// 	ATTACK_3000MS = 13
// 	ATTACK_5000MS = 14
// 	ATTACK_6000MS = 15

// 	RELEASE_6MS     = 0
// 	RELEASE_24MS    = 1
// 	RELEASE_48MS    = 2
// 	RELEASE_72MS    = 3
// 	RELEASE_114MS   = 4
// 	RELEASE_168MS   = 5
// 	RELEASE_204MS   = 6
// 	RELEASE_240MS   = 7
// 	RELEASE_300MS   = 8
// 	RELEASE_750MS   = 9
// 	RELEASE_1500MS  = 10
// 	RELEASE_2400MS  = 11
// 	RELEASE_3000MS  = 12
// 	RELEASE_9000MS  = 13
// 	RELEASE_15000MS = 14
// 	RELEASE_24000MS = 15
// )

func NewVoice(index uint8, chip chip) Voice {
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

	writeReg(v.chip, regFreqHi, uint8(realFreq>>8))
	writeReg(v.chip, regFreqLo, uint8(realFreq&0xff))
}

func (v *Voice) SetPulseWidth(dutyCycle float32) {
	v.dutyCycle = dutyCycle

	pulseWidth := uint16(math.Round(float64(dutyCycle*4096 - 1)))

	regPWLo := uint8(7*v.index + SID_REG_VOICE_PW_LO)
	regPWHi := uint8(7*v.index + SID_REG_VOICE_PW_HI)

	writeReg(v.chip, regPWHi, uint8(pulseWidth>>8)&0xf)
	writeReg(v.chip, regPWLo, uint8(pulseWidth&0xff))
}

func (v *Voice) SetEnvelope(attack uint8, decay uint8, sustain uint8, release uint8) {
	v.attack = attack & 0xf
	v.decay = decay & 0xf
	v.sustain = sustain & 0xf
	v.release = release & 0xf

	regAD := uint8(7*v.index + SID_REG_VOICE_AD)
	regSR := uint8(7*v.index + SID_REG_VOICE_SR)

	adBits := v.attack<<4 | v.decay
	srBits := v.sustain<<4 | v.release

	writeReg(v.chip, regAD, adBits)
	writeReg(v.chip, regSR, srBits)
}

func (v *Voice) SetAttack(attack uint8) {
	v.SetEnvelope(attack, v.decay, v.sustain, v.release)
}

func (v *Voice) SetDecay(decay uint8) {
	v.SetEnvelope(v.attack, decay, v.sustain, v.release)
}

func (v *Voice) SetSustain(sustain uint8) {
	v.SetEnvelope(v.attack, v.decay, sustain, v.release)
}

func (v *Voice) SetRelease(release uint8) {
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

	outBits := uint8(btoi(v.gate))
	outBits |= btoi(v.sync) << 1
	outBits |= btoi(v.ringMod) << 2
	outBits |= btoi(v.test) << 3
	outBits |= btoi(v.triangle) << 4
	outBits |= btoi(v.sawtooth) << 5
	outBits |= btoi(v.pulse) << 6
	outBits |= btoi(v.noise) << 7

	writeReg(v.chip, regCtrl, outBits)
}

func (v *Voice) Release() {
	regCtrl := uint8(7*v.index + SID_REG_VOICE_CTRL)

	v.gate = false

	outBits := uint8(btoi(v.gate))
	outBits |= btoi(v.sync) << 1
	outBits |= btoi(v.ringMod) << 2
	outBits |= btoi(v.test) << 3
	outBits |= btoi(v.triangle) << 4
	outBits |= btoi(v.sawtooth) << 5
	outBits |= btoi(v.pulse) << 6
	outBits |= btoi(v.noise) << 7

	writeReg(v.chip, regCtrl, outBits)
}
