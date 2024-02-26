package chip

import (
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
		volume:        6,
		filterLP:      false,
		filterBP:      false,
		filterHP:      false,
		filter3Off:    false,
		filterCutoff:  0,
		filterRes:     0,
		Voice:         [3]voice.Voice{voice.New(0, chip), voice.New(1, chip), voice.New(2, chip)},
		voiceFilterEn: [3]bool{false, false, false},
	}
	return dev
}

func (sd *SIDDevice) SetVolume(volume uint8) {
	sd.volume = volume

	outbits := uint8(volume & 0xf)
	outbits |= btoi(sd.filterLP) << 4
	outbits |= btoi(sd.filterBP) << 5
	outbits |= btoi(sd.filterHP) << 6
	outbits |= btoi(sd.filter3Off) << 7

	gpio.WriteReg(sd.chip, SID_REG_FILT_MODE_VOL, outbits)
}

func (sd *SIDDevice) SetFilterMode(lp bool, bp bool, hp bool) {
	sd.filterLP = lp
	sd.filterBP = bp
	sd.filterHP = hp

	outbits := uint8(sd.volume & 0xf)
	outbits |= btoi(lp) << 4
	outbits |= btoi(bp) << 5
	outbits |= btoi(hp) << 6
	outbits |= btoi(sd.filter3Off) << 7

	gpio.WriteReg(sd.chip, SID_REG_FILT_MODE_VOL, outbits)
}

func (sd *SIDDevice) SetFilterCutoff(cutoff uint16) {
	sd.filterCutoff = cutoff

	cutoffLo := uint8(cutoff & 0b111)
	cutoffHi := uint8((cutoff >> 3) & 0xFF)

	gpio.WriteReg(sd.chip, SID_REG_FILT_FC_LO, cutoffLo)
	gpio.WriteReg(sd.chip, SID_REG_FILT_FC_HI, cutoffHi)
}

func (sd *SIDDevice) SetFilterRes(resonance uint8) {
	sd.filterRes = resonance & 0xF

	outbits := uint8(0)
	for i := 0; i < 3; i++ {
		outbits |= btoi(sd.voiceFilterEn[i]) << i
	}
	outbits |= sd.filterRes << 4

	gpio.WriteReg(sd.chip, SID_REG_RES_FILT_EN, outbits)
}

func (sd *SIDDevice) SetFilterEn(voiceIndex uint8, enable bool) {
	sd.voiceFilterEn[voiceIndex] = enable

	outbits := uint8(0)
	for i := 0; i < 3; i++ {
		outbits |= btoi(sd.voiceFilterEn[i]) << i
	}
	outbits |= sd.filterRes << 4

	gpio.WriteReg(sd.chip, SID_REG_RES_FILT_EN, outbits)
}

func btoi(tf bool) uint8 {
	if tf {
		return 1
	}
	return 0
}
