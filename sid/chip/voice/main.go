package voice

import (
	"math"

	"github.com/litui/monosid/settings"
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
	index shared.VoiceIndex
	chip  shared.SidChip

	// Pitch & Timbre
	freq      float32
	dutyCycle float32

	// Note options
	gate bool
	test bool
}

func New(index shared.VoiceIndex, chip shared.SidChip) Voice {
	return Voice{
		index:     index,
		chip:      chip,
		freq:      0,
		dutyCycle: 0.5,
		gate:      false,
		test:      false,
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
	v.SetRawPulseWidth(pulseWidth)
}

func (v *Voice) SetRawPulseWidth(pulseWidth uint16) {
	settings.Storage.SetVoicePW(v.chip, v.index, pulseWidth)

	regPWLo := uint8(7*v.index + SID_REG_VOICE_PW_LO)
	regPWHi := uint8(7*v.index + SID_REG_VOICE_PW_HI)

	gpio.WriteReg(v.chip, regPWHi, uint8(pulseWidth>>8)&0xf)
	gpio.WriteReg(v.chip, regPWLo, uint8(pulseWidth&0xff))
}

func (v *Voice) SetEnvelope(attack shared.AttackRate, decay shared.DecayRate, sustain uint8, release shared.ReleaseRate) {
	settings.Storage.SetVoiceAttack(v.chip, v.index, attack&0xf)
	settings.Storage.SetVoiceDecay(v.chip, v.index, decay&0xf)
	settings.Storage.SetVoiceSustain(v.chip, v.index, sustain&0xf)
	settings.Storage.SetVoiceRelease(v.chip, v.index, release&0xf)

	regAD := uint8(7*v.index + SID_REG_VOICE_AD)
	regSR := uint8(7*v.index + SID_REG_VOICE_SR)

	adBits := uint8(attack&0xf)<<4 | uint8(decay&0xf)
	srBits := uint8(sustain&0xf)<<4 | uint8(release&0xf)

	gpio.WriteReg(v.chip, regAD, adBits)
	gpio.WriteReg(v.chip, regSR, srBits)
}

func (v *Voice) SetAttack(attack shared.AttackRate) {
	decay := settings.Storage.GetVoiceDecay(v.chip, v.index)
	sustain := settings.Storage.GetVoiceSustain(v.chip, v.index)
	release := settings.Storage.GetVoiceRelease(v.chip, v.index)
	v.SetEnvelope(attack, decay, sustain, release)
}

func (v *Voice) SetDecay(decay shared.DecayRate) {
	attack := settings.Storage.GetVoiceAttack(v.chip, v.index)
	sustain := settings.Storage.GetVoiceSustain(v.chip, v.index)
	release := settings.Storage.GetVoiceRelease(v.chip, v.index)
	v.SetEnvelope(attack, decay, sustain, release)
}

func (v *Voice) SetSustain(sustain uint8) {
	attack := settings.Storage.GetVoiceAttack(v.chip, v.index)
	decay := settings.Storage.GetVoiceDecay(v.chip, v.index)
	release := settings.Storage.GetVoiceRelease(v.chip, v.index)
	v.SetEnvelope(attack, decay, sustain, release)
}

func (v *Voice) SetRelease(release shared.ReleaseRate) {
	attack := settings.Storage.GetVoiceAttack(v.chip, v.index)
	decay := settings.Storage.GetVoiceDecay(v.chip, v.index)
	sustain := settings.Storage.GetVoiceSustain(v.chip, v.index)
	v.SetEnvelope(attack, decay, sustain, release)
}

func (v *Voice) SetWaveform(triangle bool, sawtooth bool, pulse bool, noise bool) {
	settings.Storage.SetVoiceTriangle(v.chip, v.index, triangle)
	settings.Storage.SetVoiceSawtooth(v.chip, v.index, sawtooth)
	settings.Storage.SetVoicePulse(v.chip, v.index, pulse)
	settings.Storage.SetVoiceNoise(v.chip, v.index, noise)
}

func (v *Voice) Trigger() {
	regCtrl := uint8(7*v.index + SID_REG_VOICE_CTRL)

	v.gate = true

	outBits := uint8(shared.BToI(v.gate))
	outBits |= shared.BToI(settings.Storage.GetVoiceSync(v.chip, v.index)) << 1
	outBits |= shared.BToI(settings.Storage.GetVoiceRingMod(v.chip, v.index)) << 2
	outBits |= shared.BToI(v.test) << 3
	outBits |= shared.BToI(settings.Storage.GetVoiceTriangle(v.chip, v.index)) << 4
	outBits |= shared.BToI(settings.Storage.GetVoiceSawtooth(v.chip, v.index)) << 5
	outBits |= shared.BToI(settings.Storage.GetVoicePulse(v.chip, v.index)) << 6
	outBits |= shared.BToI(settings.Storage.GetVoiceNoise(v.chip, v.index)) << 7

	gpio.WriteReg(v.chip, regCtrl, outBits)
}

func (v *Voice) Release() {
	regCtrl := uint8(7*v.index + SID_REG_VOICE_CTRL)

	v.gate = false

	outBits := uint8(shared.BToI(v.gate))
	outBits |= shared.BToI(settings.Storage.GetVoiceSync(v.chip, v.index)) << 1
	outBits |= shared.BToI(settings.Storage.GetVoiceRingMod(v.chip, v.index)) << 2
	outBits |= shared.BToI(v.test) << 3
	outBits |= shared.BToI(settings.Storage.GetVoiceTriangle(v.chip, v.index)) << 4
	outBits |= shared.BToI(settings.Storage.GetVoiceSawtooth(v.chip, v.index)) << 5
	outBits |= shared.BToI(settings.Storage.GetVoicePulse(v.chip, v.index)) << 6
	outBits |= shared.BToI(settings.Storage.GetVoiceNoise(v.chip, v.index)) << 7

	gpio.WriteReg(v.chip, regCtrl, outBits)
}
