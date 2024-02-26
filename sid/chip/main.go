package chip

import (
	"github.com/litui/monosid/settings"
	"github.com/litui/monosid/shared"
	"github.com/litui/monosid/sid/chip/voice"
	"github.com/litui/monosid/sid/gpio"
)

const (
	SID_REG_FILT_FC_LO    = 21
	SID_REG_FILT_FC_HI    = 22
	SID_REG_RES_FILT_EN   = 23
	SID_REG_FILT_MODE_VOL = 24
)

type SIDDevice struct {
	chip shared.SidChip

	volume uint8

	filterLP   bool
	filterBP   bool
	filterHP   bool
	filter3Off bool

	filterCutoff uint16
	filterRes    uint8

	Voice         [3]voice.Voice
	voiceFilterEn [3]bool
}

func New(chip shared.SidChip) SIDDevice {
	dev := SIDDevice{
		chip:          chip,
		Voice:         [3]voice.Voice{voice.New(0, chip), voice.New(1, chip), voice.New(2, chip)},
		voiceFilterEn: [3]bool{false, false, false},
	}
	return dev
}

func (sd *SIDDevice) SetVolume(volume uint8) {
	settings.Storage.SetVolume(sd.chip, volume&0xf)

	outbits := uint8(volume & 0xf)
	outbits |= btoi(settings.Storage.GetFilterLP(sd.chip)) << 4
	outbits |= btoi(settings.Storage.GetFilterBP(sd.chip)) << 5
	outbits |= btoi(settings.Storage.GetFilterHP(sd.chip)) << 6
	outbits |= btoi(settings.Storage.GetFilter3Off(sd.chip)) << 7

	gpio.WriteReg(sd.chip, SID_REG_FILT_MODE_VOL, outbits)
}

func (sd *SIDDevice) SetFilterMode(lp bool, bp bool, hp bool) {
	settings.Storage.SetFilterLP(sd.chip, lp)
	settings.Storage.SetFilterBP(sd.chip, bp)
	settings.Storage.SetFilterHP(sd.chip, hp)

	outbits := uint8(settings.Storage.GetVolume(sd.chip))
	outbits |= btoi(lp) << 4
	outbits |= btoi(bp) << 5
	outbits |= btoi(hp) << 6
	outbits |= btoi(settings.Storage.GetFilter3Off(sd.chip)) << 7

	gpio.WriteReg(sd.chip, SID_REG_FILT_MODE_VOL, outbits)
}

func (sd *SIDDevice) SetFilterCutoff(cutoff uint16) {
	settings.Storage.SetFilterCutoff(sd.chip, cutoff)

	cutoffLo := uint8(cutoff & 0b111)
	cutoffHi := uint8((cutoff >> 3) & 0xFF)

	gpio.WriteReg(sd.chip, SID_REG_FILT_FC_LO, cutoffLo)
	gpio.WriteReg(sd.chip, SID_REG_FILT_FC_HI, cutoffHi)
}

func (sd *SIDDevice) SetFilterRes(resonance uint8) {
	settings.Storage.SetFilterRes(sd.chip, resonance)

	outbits := uint8(0)
	for i := 0; i < 4; i++ {
		outbits |= btoi(settings.Storage.GetFilterEn(sd.chip, shared.VoiceIndex(i))) << i
	}
	outbits |= resonance << 4

	gpio.WriteReg(sd.chip, SID_REG_RES_FILT_EN, outbits)
}

func (sd *SIDDevice) SetFilterEn(voice shared.VoiceIndex, enable bool) {
	settings.Storage.SetFilterEn(sd.chip, shared.VoiceIndex(voice), enable)

	outbits := uint8(0)
	for i := 0; i < 4; i++ {
		outbits |= btoi(settings.Storage.GetFilterEn(sd.chip, shared.VoiceIndex(i))) << i
	}
	outbits |= settings.Storage.GetFilterRes(sd.chip) << 4

	gpio.WriteReg(sd.chip, SID_REG_RES_FILT_EN, outbits)
}

func btoi(tf bool) uint8 {
	if tf {
		return 1
	}
	return 0
}
