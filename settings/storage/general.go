package storage

import (
	"github.com/litui/monosid/log"
	"github.com/litui/monosid/shared"
)

const (
	// Going to use the entire first block of 4096 bits for general just in case we need the storage for future functionality

	settingsGeneralStorageAddr = 0x100000
	settingsGeneralTotalBits   = 64
	settingsGeneralTotalBytes  = settingsGeneralTotalBits / 8

	settingsGeneralDatatypeOffset         = 0
	settingsGeneralDatatypeBits           = 4
	settingsGeneralDatatypeDefault        = 0b1000
	settingsGeneralTemperamentOffset      = settingsGeneralDatatypeOffset + settingsGeneralDatatypeBits
	settingsGeneralTemperamentBits        = 4
	settingsGeneralTemperamentDefault     = 0b0000
	settingsGeneralQuantizerModeOffset    = settingsGeneralTemperamentOffset + settingsGeneralTemperamentBits
	settingsGeneralQuantizerModeBits      = 2
	settingsGeneralQuantizerModeDefault   = 0b00
	settingsGeneralQuantizerRootOffset    = settingsGeneralQuantizerModeOffset + settingsGeneralQuantizerModeBits
	settingsGeneralQuantizerRootBits      = 4
	settingsGeneralQuantizerRootDefault   = 0b0000
	settingsGeneralQuantizerScaleOffset   = settingsGeneralQuantizerRootOffset + settingsGeneralQuantizerRootBits
	settingsGeneralQuantizerScaleBits     = 4
	settingsGeneralQuantizerScaleDefault  = 0b0000
	settingsGeneralMidiC1V1ChannelOffset  = settingsGeneralQuantizerScaleOffset + settingsGeneralQuantizerScaleBits
	settingsGeneralMidiC1V1ChannelBits    = 4
	settingsGeneralMidiC1V1ChannelDefault = 0b0000
	settingsGeneralMidiC1V2ChannelOffset  = settingsGeneralMidiC1V1ChannelOffset + settingsGeneralMidiC1V1ChannelBits
	settingsGeneralMidiC1V2ChannelBits    = 4
	settingsGeneralMidiC1V2ChannelDefault = 0b0000
	settingsGeneralMidiC1V3ChannelOffset  = settingsGeneralMidiC1V2ChannelOffset + settingsGeneralMidiC1V2ChannelBits
	settingsGeneralMidiC1V3ChannelBits    = 4
	settingsGeneralMidiC1V3ChannelDefault = 0b0000
	settingsGeneralMidiC2V1ChannelOffset  = settingsGeneralMidiC1V3ChannelOffset + settingsGeneralMidiC1V3ChannelBits
	settingsGeneralMidiC2V1ChannelBits    = 4
	settingsGeneralMidiC2V1ChannelDefault = 0b0000
	settingsGeneralMidiC2V2ChannelOffset  = settingsGeneralMidiC2V1ChannelOffset + settingsGeneralMidiC2V1ChannelBits
	settingsGeneralMidiC2V2ChannelBits    = 4
	settingsGeneralMidiC2V2ChannelDefault = 0b0000
	settingsGeneralMidiC2V3ChannelOffset  = settingsGeneralMidiC2V2ChannelOffset + settingsGeneralMidiC2V2ChannelBits
	settingsGeneralMidiC2V3ChannelBits    = 4
	settingsGeneralMidiC2V3ChannelDefault = 0b0000
	settingsGeneralPatchSelOffset         = settingsGeneralMidiC2V3ChannelOffset + settingsGeneralMidiC2V3ChannelBits
	settingsGeneralPatchSelBits           = 7
	settingsGeneralPatchSelDefault        = 0b0000000
)

type QuantizerMode uint8
type ScaleNote uint8
type Scale uint8

const (
	QuantizerOff QuantizerMode = iota
	QuantizerOn
	QuantizerRoundUp
	QuantizerRoundDown

	NoteC ScaleNote = iota
	NoteCSharp
	NoteD
	NoteDSharp
	NoteE
	NoteF
	NoteFSharp
	NoteG
	NoteGSharp
	NoteA
	NoteASharp
	NoteB

	ScaleIonian Scale = iota
	ScaleDorian
	ScalePhrygian
	ScaleLydian
	ScaleMixolydian
	ScaleAeolian
	ScaleLocrian
	ScalePentatonicMajor
	ScalePentatonicMinor
	ScaleEgyptian
	ScaleBlues
	ScaleHarmonicMajor
	ScaleHarmonicMinor
)

func (m *StorageDevice) newGeneral() {
	m.generalMem = uint64(0)
	m.generalMem |= uint64(settingsGeneralDatatypeDefault) << settingsGeneralDatatypeOffset
	m.generalMem |= uint64(settingsGeneralTemperamentDefault) << settingsGeneralTemperamentOffset
	m.generalMem |= uint64(settingsGeneralQuantizerModeDefault) << settingsGeneralQuantizerModeOffset
	m.generalMem |= uint64(settingsGeneralQuantizerRootDefault) << settingsGeneralQuantizerRootOffset
	m.generalMem |= uint64(settingsGeneralQuantizerScaleDefault) << settingsGeneralQuantizerScaleOffset
	m.generalMem |= uint64(settingsGeneralMidiC1V1ChannelDefault) << settingsGeneralMidiC1V1ChannelOffset
	m.generalMem |= uint64(settingsGeneralMidiC1V2ChannelDefault) << settingsGeneralMidiC1V2ChannelOffset
	m.generalMem |= uint64(settingsGeneralMidiC1V3ChannelDefault) << settingsGeneralMidiC1V3ChannelOffset
	m.generalMem |= uint64(settingsGeneralMidiC2V1ChannelDefault) << settingsGeneralMidiC2V1ChannelOffset
	m.generalMem |= uint64(settingsGeneralMidiC2V2ChannelDefault) << settingsGeneralMidiC2V2ChannelOffset
	m.generalMem |= uint64(settingsGeneralMidiC2V3ChannelDefault) << settingsGeneralMidiC1V3ChannelOffset
	m.generalMem |= uint64(settingsGeneralPatchSelDefault) << settingsGeneralPatchSelOffset

	log.Logf("Reset general settings")
	m.generalChanged = true
}

func (m *StorageDevice) saveGeneral() bool {
	var settingsBytes []byte = make([]byte, settingsGeneralTotalBytes)
	for i := 0; i < settingsGeneralTotalBytes; i++ {
		settingsBytes[i] = byte((m.generalMem >> (i * 8)) & 0xff)
	}
	_, err := memDev.WriteAt(settingsBytes, settingsGeneralStorageAddr)
	if err != nil {
		log.Logf("Failed to write settings")
		return false
	}

	m.generalChanged = false
	log.Logf("Saved general settings")
	return true
}

func (m *StorageDevice) loadGeneral() bool {
	// Start with a clear slate
	m.generalMem = uint64(0)

	// Attempt to load settings. If that fails, generate new settings.
	var settingsBytes []byte = make([]byte, settingsGeneralTotalBits/8)
	_, err := memDev.ReadAt(settingsBytes, settingsGeneralStorageAddr)
	if err != nil || settingsBytes[0]&0xf != settingsGeneralDatatypeDefault {
		m.newGeneral()
		m.generalLoaded = true
		return false
	}

	// Remap loaded bytes into a uint64
	for i, b := range settingsBytes {
		m.generalMem |= uint64(b) << (i * 8)
	}
	m.generalChanged = false

	log.Logf("Loaded general settings")
	return true
}

func (m *StorageDevice) SetQuantizerMode(mode QuantizerMode) {
	setValue(&m.generalMem, settingsGeneralQuantizerModeOffset, mode, settingsGeneralQuantizerModeBits)
}

func (m *StorageDevice) ResetQuantizerMode() {
	m.SetQuantizerMode(settingsGeneralQuantizerModeDefault)
}

func (m *StorageDevice) GetQuantizerMode() QuantizerMode {
	return getValue[QuantizerMode](m.generalMem, settingsGeneralQuantizerModeOffset, settingsGeneralQuantizerModeBits)
}

func (m *StorageDevice) SetQuantizerRoot(note ScaleNote) {
	setValue(&m.generalMem, settingsGeneralQuantizerRootOffset, note, settingsGeneralQuantizerRootBits)
}

func (m *StorageDevice) ResetQuantizerRoot() {
	m.SetQuantizerRoot(settingsGeneralQuantizerRootDefault)
}

func (m *StorageDevice) GetQuantizerRoot() ScaleNote {
	return getValue[ScaleNote](m.generalMem, settingsGeneralQuantizerRootOffset, settingsGeneralQuantizerRootBits)
}

func (m *StorageDevice) SetQuantizerScale(scale Scale) {
	setValue(&m.generalMem, settingsGeneralQuantizerScaleOffset, scale, settingsGeneralQuantizerScaleBits)
}

func (m *StorageDevice) ResetQuantizerScale() {
	m.SetQuantizerScale(settingsGeneralQuantizerScaleDefault)
}

func (m *StorageDevice) GetQuantizerScale() Scale {
	return getValue[Scale](m.generalMem, settingsGeneralQuantizerScaleOffset, settingsGeneralQuantizerScaleBits)
}

func (m *StorageDevice) SetMidiChannel(chip shared.SidChip, voice shared.VoiceIndex, channel uint8) {
	chanSize := uint8(settingsGeneralMidiC1V1ChannelBits)
	chanAddr := uint8(settingsGeneralMidiC1V1ChannelOffset)

	if chip == 1 {
		switch voice {
		case 0:
			chanAddr = settingsGeneralMidiC2V1ChannelOffset
		case 1:
			chanAddr = settingsGeneralMidiC2V2ChannelOffset
		case 2:
			chanAddr = settingsGeneralMidiC2V3ChannelOffset
		}
	} else {
		switch voice {
		case 0:
			chanAddr = settingsGeneralMidiC1V1ChannelOffset
		case 1:
			chanAddr = settingsGeneralMidiC1V2ChannelOffset
		case 2:
			chanAddr = settingsGeneralMidiC1V3ChannelOffset
		}
	}
	setValue[uint8](&m.generalMem, chanAddr, channel, chanSize)
}

func (m *StorageDevice) ResetMidiChannel(chip shared.SidChip, voice shared.VoiceIndex) {
	defVal := uint8(settingsGeneralMidiC1V1ChannelDefault)

	if chip == 1 {
		switch voice {
		case 0:
			defVal = settingsGeneralMidiC2V1ChannelDefault
		case 1:
			defVal = settingsGeneralMidiC2V2ChannelDefault
		case 2:
			defVal = settingsGeneralMidiC2V3ChannelDefault
		}
	} else {
		switch voice {
		case 0:
			defVal = settingsGeneralMidiC1V1ChannelDefault
		case 1:
			defVal = settingsGeneralMidiC1V2ChannelDefault
		case 2:
			defVal = settingsGeneralMidiC1V3ChannelDefault
		}
	}

	m.SetMidiChannel(chip, voice, defVal)
}

func (m *StorageDevice) GetMidiChannel(chip shared.SidChip, voice shared.VoiceIndex) uint8 {
	chanSize := uint8(settingsGeneralMidiC1V1ChannelBits)
	chanAddr := uint8(settingsGeneralMidiC1V1ChannelOffset)

	if chip == 1 {
		switch voice {
		case 0:
			chanAddr = settingsGeneralMidiC2V1ChannelOffset
		case 1:
			chanAddr = settingsGeneralMidiC2V2ChannelOffset
		case 2:
			chanAddr = settingsGeneralMidiC2V3ChannelOffset
		}
	} else {
		switch voice {
		case 0:
			chanAddr = settingsGeneralMidiC1V1ChannelOffset
		case 1:
			chanAddr = settingsGeneralMidiC1V2ChannelOffset
		case 2:
			chanAddr = settingsGeneralMidiC1V3ChannelOffset
		}
	}

	return getValue[uint8](m.generalMem, chanAddr, chanSize)
}

func (m *StorageDevice) SetSelectedPatch(index uint8) {
	setValue(&m.generalMem, settingsGeneralPatchSelOffset, index, settingsGeneralPatchSelBits)
}

func (m *StorageDevice) ResetSelectedPatch() {
	m.SetSelectedPatch(settingsGeneralPatchSelDefault)
}

func (m *StorageDevice) GetSelectedPatch() uint8 {
	return getValue[uint8](m.generalMem, settingsGeneralPatchSelOffset, settingsGeneralPatchSelBits)
}
