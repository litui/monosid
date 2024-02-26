package settings

import (
	"machine"
	"time"

	"github.com/litui/monosid/log"
)

const (
	// Mapping for General data storage

	flashBlockWriteSize = 4096

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

	// Mapping for Patch data storage

	// Patches are 512 bits long (Sid + 3 voices) * 2
	// Storing SIDs separately just in case I decide to break out left and right settings in places (eg detune?)

	// Make sure first patch is on a 4096 byte (0x1000h) boundary for writing)
	firstPatchStorageAddr = 0x101000
	patchLengthBits       = 512
	patchLengthBytes      = patchLengthBits / 8
	patchCount            = 128

	sid1Addr   = 0
	sid1V1Addr = 64
	sid1V2Addr = 128
	sid1V3Addr = 192
	sid2Addr   = 256
	sid2V1Addr = 320
	sid2V2Addr = 384
	sid2V3Addr = 448

	settingsPatchSidDatatypeOffset      = 0
	settingsPatchSidDatatypeBits        = 4
	settingsPatchSidDatatypeDefault     = 0b0100
	settingsPatchSidVolumeOffset        = settingsPatchSidDatatypeOffset + settingsPatchSidDatatypeBits
	settingsPatchSidVolumeBits          = 4
	settingsPatchSidVolumeDefault       = 6
	settingsPatchSidFilterLPOffset      = settingsPatchSidVolumeOffset + settingsPatchSidVolumeBits
	settingsPatchSidFilterLPBits        = 1
	settingsPatchSidFilterLPDefault     = 0
	settingsPatchSidFilterBPOffset      = settingsPatchSidFilterLPOffset + settingsPatchSidFilterLPBits
	settingsPatchSidFilterBPBits        = 1
	settingsPatchSidFilterBPDefault     = 0
	settingsPatchSidFilterHPOffset      = settingsPatchSidFilterBPOffset + settingsPatchSidFilterBPBits
	settingsPatchSidFilterHPBits        = 1
	settingsPatchSidFilterHPDefault     = 0
	settingsPatchSid3OffOffset          = settingsPatchSidFilterHPOffset + settingsPatchSidFilterHPBits
	settingsPatchSid3OffBits            = 1
	settingsPatchSid3OffDefault         = 0
	settingsPatchSidFilterV1EnOffset    = settingsPatchSid3OffOffset + settingsPatchSid3OffBits
	settingsPatchSidFilterV1EnBits      = 1
	settingsPatchSidFilterV1EnDefault   = 0
	settingsPatchSidFilterV2EnOffset    = settingsPatchSidFilterV1EnOffset + settingsPatchSidFilterV1EnBits
	settingsPatchSidFilterV2EnBits      = 1
	settingsPatchSidFilterV2EnDefault   = 0
	settingsPatchSidFilterV3EnOffset    = settingsPatchSidFilterV2EnOffset + settingsPatchSidFilterV2EnBits
	settingsPatchSidFilterV3EnBits      = 1
	settingsPatchSidFilterV3EnDefault   = 0
	settingsPatchSidFilterExEnOffset    = settingsPatchSidFilterV3EnOffset + settingsPatchSidFilterV3EnBits
	settingsPatchSidFilterExEnBits      = 1
	settingsPatchSidFilterExEnDefault   = 0
	settingsPatchSidFilterResOffset     = settingsPatchSidFilterExEnOffset + settingsPatchSidFilterExEnBits
	settingsPatchSidFilterResBits       = 4
	settingsPatchSidFilterResDefault    = 0
	settingsPatchSidFilterCutoffOffset  = settingsPatchSidFilterResOffset + settingsPatchSidFilterResBits
	settingsPatchSidFilterCutoffBits    = 11
	settingsPatchSidFilterCutoffDefault = 1024

	settingsPatchVDatatypeOffset     = 0
	settingsPatchVDatatypeBits       = 4
	settingsPatchVDatatypeDefault    = 0b0101
	settingsPatchVSyncOffset         = settingsPatchVDatatypeOffset + settingsPatchVDatatypeBits
	settingsPatchVSyncBits           = 1
	settingsPatchVSyncDefault        = 0
	settingsPatchVRingModOffset      = settingsPatchVSyncOffset + settingsPatchVSyncBits
	settingsPatchVRingModBits        = 1
	settingsPatchVRingModDefault     = 0
	settingsPatchVTriangleOffset     = settingsPatchVRingModOffset + settingsPatchVRingModBits
	settingsPatchVTriangleBits       = 1
	settingsPatchVTriangleDefault    = 0
	settingsPatchVSawtoothOffset     = settingsPatchVTriangleOffset + settingsPatchVTriangleBits
	settingsPatchVSawtoothBits       = 1
	settingsPatchVSawtoothDefault    = 1
	settingsPatchVPulseOffset        = settingsPatchVSawtoothOffset + settingsPatchVSawtoothBits
	settingsPatchVPulseBits          = 1
	settingsPatchVPulseDefault       = 0
	settingsPatchVNoiseOffset        = settingsPatchVPulseOffset + settingsPatchVPulseBits
	settingsPatchVNoiseBits          = 1
	settingsPatchVNoiseDefault       = 0
	settingsPatchVPWOffset           = settingsPatchVNoiseOffset + settingsPatchVNoiseBits
	settingsPatchVPWBits             = 12
	settingsPatchVPWDefault          = 1024
	settingsPatchVAttackOffset       = settingsPatchVPWOffset + settingsPatchVPWBits
	settingsPatchVAttackBits         = 4
	settingsPatchVAttackDefault      = 1
	settingsPatchVDecayOffset        = settingsPatchVAttackOffset + settingsPatchVAttackBits
	settingsPatchVDecayBits          = 4
	settingsPatchVDecayDefault       = 0
	settingsPatchVSustainOffset      = settingsPatchVDecayOffset + settingsPatchVDecayBits
	settingsPatchVSustainBits        = 4
	settingsPatchVSustainDefault     = 4
	settingsPatchVReleaseOffset      = settingsPatchVSustainOffset + settingsPatchVSustainBits
	settingsPatchVReleaseBits        = 4
	settingsPatchVReleaseDefault     = 8
	settingsPatchVDetuneCentsOffset  = settingsPatchVReleaseOffset + settingsPatchVReleaseBits
	settingsPatchVDetuneCentsBits    = 8
	settingsPatchVDetuneCentsDefault = 0
)

type StorageDevice struct {
	generalMem      uint64
	generalChanged  bool
	generalLoaded   bool
	patchMem        [2][4]uint64
	patchChanged    bool
	patchLoaded     bool
	lastGeneralSave time.Time
	lastPatchSave   time.Time
}

type Numerics interface {
	~uint64 | ~uint32 | ~uint16 | ~uint8 | uint | ~int64 | ~int32 | ~int16 | ~int8 | ~int | ~float32 | ~float64
}

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

var (
	memDev = machine.Flash
)

func New() StorageDevice {
	return StorageDevice{
		generalMem:      0,
		generalChanged:  false,
		generalLoaded:   false,
		lastGeneralSave: time.Time{},
		lastPatchSave:   time.Time{},
	}
}

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

func (m *StorageDevice) newPatch() {
	for c := 0; c < 2; c++ {
		for b := 0; b < 4; b++ {
			m.patchMem[c][b] = uint64(0)
		}

		m.patchMem[c][0] |= uint64(settingsPatchSidDatatypeDefault) << settingsPatchSidDatatypeOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidVolumeDefault) << settingsPatchSidVolumeOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterLPDefault) << settingsPatchSidFilterLPOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterBPDefault) << settingsPatchSidFilterBPOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterHPDefault) << settingsPatchSidFilterHPOffset
		m.patchMem[c][0] |= uint64(settingsPatchSid3OffDefault) << settingsPatchSid3OffOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterV1EnDefault) << settingsPatchSidFilterV1EnOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterV2EnDefault) << settingsPatchSidFilterV2EnOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterV3EnDefault) << settingsPatchSidFilterV3EnOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterExEnDefault) << settingsPatchSidFilterExEnOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterResDefault) << settingsPatchSidFilterResOffset
		m.patchMem[c][0] |= uint64(settingsPatchSidFilterCutoffDefault) << settingsPatchSidFilterCutoffOffset

		for b := 1; b < 4; b++ {
			m.patchMem[c][b] |= uint64(settingsPatchVDatatypeDefault) << settingsPatchVDatatypeOffset
			m.patchMem[c][b] |= uint64(settingsPatchVSyncDefault) << settingsPatchVSyncOffset
			m.patchMem[c][b] |= uint64(settingsPatchVRingModDefault) << settingsPatchVRingModOffset
			m.patchMem[c][b] |= uint64(settingsPatchVTriangleDefault) << settingsPatchVTriangleOffset
			m.patchMem[c][b] |= uint64(settingsPatchVSawtoothDefault) << settingsPatchVSawtoothOffset
			m.patchMem[c][b] |= uint64(settingsPatchVPulseDefault) << settingsPatchVPulseOffset
			m.patchMem[c][b] |= uint64(settingsPatchVNoiseDefault) << settingsPatchVNoiseOffset
			m.patchMem[c][b] |= uint64(settingsPatchVPWDefault) << settingsPatchVPWOffset
			m.patchMem[c][b] |= uint64(settingsPatchVAttackDefault) << settingsPatchVAttackOffset
			m.patchMem[c][b] |= uint64(settingsPatchVDecayDefault) << settingsPatchVDecayOffset
			m.patchMem[c][b] |= uint64(settingsPatchVSustainDefault) << settingsPatchVSustainOffset
			m.patchMem[c][b] |= uint64(settingsPatchVReleaseDefault) << settingsPatchVReleaseOffset
			m.patchMem[c][b] |= uint64(settingsPatchVDetuneCentsDefault) << settingsPatchVDetuneCentsOffset
		}
	}

	log.Logf("Reset patch")
	m.patchChanged = true
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

func (m *StorageDevice) savePatch(index uint8) bool {
	// Determine which 4096 byte bank we'll need to load up (so we don't lose other patch data when writing)
	bankAddress := int64(firstPatchStorageAddr)
	byteOffset := 0
	if index > 64 {
		bankAddress = int64(firstPatchStorageAddr + flashBlockWriteSize)
		byteOffset = 64
	}

	// Load whole 4096 byte bank into memory
	var patchBankBytes []byte = make([]byte, flashBlockWriteSize)
	_, err := memDev.ReadAt(patchBankBytes, bankAddress)
	if err != nil {
		log.Logf("Failed to load patch bank for writing.")
		return false
	}

	// Replace bytes in upper or lower patchBank with current patch
	relevantByte := (int(index) - byteOffset) * patchLengthBytes
	for c := 0; c < 2; c++ {
		for b := 0; b < 4; b++ {
			for i := 0; i < 8; i++ {
				patchBankBytes[relevantByte] = byte((m.patchMem[c][b] >> (i * 8)) & 0xff)
				relevantByte++
			}
		}
	}

	_, err = memDev.WriteAt(patchBankBytes, bankAddress)
	if err != nil {
		log.Logf("Failed to save patch #%d", index)
		return false
	}

	m.patchChanged = false
	log.Logf("Saved patch #%d", index)
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

func (m *StorageDevice) loadPatch(index uint8) bool {
	// Start with a clear slate
	for c := 0; c < 2; c++ {
		for b := 0; b < 4; b++ {
			m.patchMem[c][b] = uint64(0)
		}
	}

	var patchBytes []byte = make([]byte, patchLengthBytes)

	address := firstPatchStorageAddr + int64(index)*patchLengthBytes
	_, err := memDev.ReadAt(patchBytes, address)
	if err != nil || patchBytes[0]&0xf != settingsPatchSidDatatypeDefault {
		m.newPatch()
		m.patchLoaded = true
		return false
	}

	i := 0
	for c := 0; c < 2; c++ {
		for b := 0; b < 4; b++ {
			m.patchMem[c][b] |= uint64(patchBytes[i]) << (i * 8)
			i++
		}
	}
	m.patchChanged = false

	log.Logf("Loaded patch #%d", index)
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

func (m *StorageDevice) SetMidiChannel(chip uint8, voice uint8, channel uint8) {
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
	setValue(&m.generalMem, chanAddr, channel, chanSize)
}

func (m *StorageDevice) ResetMidiChannel(chip uint8, voice uint8) {
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

func (m *StorageDevice) GetMidiChannel(chip uint8, voice uint8) uint8 {
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

func setValue[InputType Numerics](target *uint64, offset uint8, data InputType, length uint8) {
	for i := 0; i < int(length); i++ {
		bitVal := (uint64(data) >> i) & 1
		if bitVal == 1 {
			// Set
			*target |= (uint64(1) << offset)
		} else {
			// Clear
			*target &= ^(uint64(1) << offset)
		}
	}
}

func getValue[ReturnType Numerics](source uint64, offset uint8, length uint8) ReturnType {
	retval := uint64(0)
	for i := 0; i < int(length); i++ {
		retval |= ((source >> (uint64(i) + uint64(offset))) & 1) << i
	}
	return ReturnType(retval)
}

func (m *StorageDevice) Init() {
	m.loadGeneral()
	m.loadPatch(m.GetSelectedPatch())
}

func (m *StorageDevice) Tick() {
	if m.generalChanged {
		m.saveGeneral()
	}
	if m.patchChanged {
		m.savePatch(m.GetSelectedPatch())
	}
	time.Sleep(time.Second * 5)
}
