package sid

const (
	SID_REG_VOICE_FREQ_LO = 0
	SID_REG_VOICE_FREQ_HI = 1
	SID_REG_VOICE_PW_LO   = 2
	SID_REG_VOICE_PW_HI   = 3
	SID_REG_VOICE_CTRL    = 4
	SID_REG_VOICE_AD      = 5
	SID_REG_VOICE_SR      = 6

	SID_REG_FILT_FC_LO    = 21
	SID_REG_FILT_FC_HI    = 22
	SID_REG_RES_FILT_EN   = 23
	SID_REG_FILT_MODE_VOL = 24
)

type SIDDevice struct {
	chip chip

	volume uint8

	filterLP   bool
	filterBP   bool
	filterHP   bool
	filter3Off bool

	filterCutoff uint16
	filterRes    uint8

	Voice         [3]Voice
	voiceFilterEn [3]bool
}

func NewSID(chip chip) SIDDevice {
	dev := SIDDevice{
		chip:          chip,
		volume:        6,
		filterLP:      false,
		filterBP:      false,
		filterHP:      false,
		filter3Off:    false,
		filterCutoff:  0,
		filterRes:     0,
		Voice:         [3]Voice{NewVoice(0, chip), NewVoice(1, chip), NewVoice(2, chip)},
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

	writeReg(sd.chip, SID_REG_FILT_MODE_VOL, outbits)
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

	writeReg(sd.chip, SID_REG_FILT_MODE_VOL, outbits)
}

func (sd *SIDDevice) SetFilterCutoff(cutoff uint16) {
	sd.filterCutoff = cutoff

	cutoffLo := uint8(cutoff & 0b111)
	cutoffHi := uint8((cutoff >> 3) & 0xFF)

	writeReg(sd.chip, SID_REG_FILT_FC_LO, cutoffLo)
	writeReg(sd.chip, SID_REG_FILT_FC_HI, cutoffHi)
}

func (sd *SIDDevice) SetFilterRes(resonance uint8) {
	sd.filterRes = resonance & 0xF

	outbits := uint8(0)
	for i := 0; i < 3; i++ {
		outbits |= btoi(sd.voiceFilterEn[i]) << i
	}
	outbits |= sd.filterRes << 4

	writeReg(sd.chip, SID_REG_RES_FILT_EN, outbits)
}

func (sd *SIDDevice) SetFilterEn(voiceIndex uint8, enable bool) {
	sd.voiceFilterEn[voiceIndex] = enable

	outbits := uint8(0)
	for i := 0; i < 3; i++ {
		outbits |= btoi(sd.voiceFilterEn[i]) << i
	}
	outbits |= sd.filterRes << 4

	writeReg(sd.chip, SID_REG_RES_FILT_EN, outbits)
}

func btoi(tf bool) uint8 {
	if tf {
		return 1
	}
	return 0
}
