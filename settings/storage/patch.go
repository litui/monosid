package storage

import (
	"github.com/litui/monosid/log"
	"github.com/litui/monosid/shared"
)

const (
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

func (m *StorageDevice) SetVolume(chip shared.SidChip, volume uint8) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSidVolumeOffset, volume, settingsPatchSidVolumeBits)
}

func (m *StorageDevice) ResetVolume(chip shared.SidChip) {
	m.SetVolume(chip, settingsPatchSidVolumeDefault)
}

func (m *StorageDevice) GetVolume(chip shared.SidChip) uint8 {
	return getValue[uint8](m.patchMem[chip][sidPatchData], settingsPatchSidVolumeOffset, settingsPatchSidVolumeBits)
}

func (m *StorageDevice) SetFilterLP(chip shared.SidChip, enable bool) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSidFilterLPOffset, shared.BToI(enable), settingsPatchSidFilterLPBits)
}

func (m *StorageDevice) ResetFilterLP(chip shared.SidChip) {
	m.SetFilterLP(chip, shared.IToB(settingsPatchSidFilterLPDefault))
}

func (m *StorageDevice) GetFilterLP(chip shared.SidChip) bool {
	return shared.IToB(getValue[uint8](m.patchMem[chip][sidPatchData], settingsPatchSidFilterLPOffset, settingsPatchSidFilterLPBits))
}

func (m *StorageDevice) SetFilterBP(chip shared.SidChip, enable bool) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSidFilterBPOffset, shared.BToI(enable), settingsPatchSidFilterBPBits)
}

func (m *StorageDevice) ResetFilterBP(chip shared.SidChip) {
	m.SetFilterBP(chip, shared.IToB(settingsPatchSidFilterBPDefault))
}

func (m *StorageDevice) GetFilterBP(chip shared.SidChip) bool {
	return shared.IToB(getValue[uint8](m.patchMem[chip][sidPatchData], settingsPatchSidFilterBPOffset, settingsPatchSidFilterBPBits))
}

func (m *StorageDevice) SetFilterHP(chip shared.SidChip, enable bool) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSidFilterHPOffset, shared.BToI(enable), settingsPatchSidFilterHPBits)
}

func (m *StorageDevice) ResetFilterHP(chip shared.SidChip) {
	m.SetFilterHP(chip, shared.IToB(settingsPatchSidFilterHPDefault))
}

func (m *StorageDevice) GetFilterHP(chip shared.SidChip) bool {
	return shared.IToB(getValue[uint8](m.patchMem[chip][sidPatchData], settingsPatchSidFilterHPOffset, settingsPatchSidFilterHPBits))
}

func (m *StorageDevice) Set3Off(chip shared.SidChip, enable bool) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSid3OffOffset, shared.BToI(enable), settingsPatchSid3OffBits)
}

func (m *StorageDevice) Reset3Off(chip shared.SidChip) {
	m.Set3Off(chip, shared.IToB(settingsPatchSid3OffDefault))
}

func (m *StorageDevice) GetFilter3Off(chip shared.SidChip) bool {
	return shared.IToB(getValue[uint8](m.patchMem[chip][sidPatchData], settingsPatchSid3OffOffset, settingsPatchSid3OffBits))
}

func (m *StorageDevice) SetFilterEn(chip shared.SidChip, voice shared.VoiceIndex, enable bool) {
	addr := uint8(settingsPatchSidFilterV1EnOffset)
	size := uint8(settingsPatchSidFilterV1EnBits)
	switch voice {
	case shared.Voice2:
		addr = settingsPatchSidFilterV2EnOffset
	case shared.Voice3:
		addr = settingsPatchSidFilterV3EnOffset
	case shared.VoiceEx:
		addr = settingsPatchSidFilterExEnOffset
	}
	setValue(&m.patchMem[chip][sidPatchData], addr, shared.BToI(enable), size)
}

func (m *StorageDevice) ResetFilterEn(chip shared.SidChip, voice shared.VoiceIndex) {
	defVal := uint8(settingsPatchSidFilterV1EnDefault)
	switch voice {
	case shared.Voice2:
		defVal = settingsPatchSidFilterV2EnDefault
	case shared.Voice3:
		defVal = settingsPatchSidFilterV3EnDefault
	case shared.VoiceEx:
		defVal = settingsPatchSidFilterExEnDefault
	}
	m.SetFilterEn(chip, voice, shared.IToB(defVal))
}

func (m *StorageDevice) GetFilterEn(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchSidFilterV1EnOffset)
	size := uint8(settingsPatchSidFilterV1EnBits)
	switch voice {
	case shared.Voice2:
		addr = settingsPatchSidFilterV2EnOffset
	case shared.Voice3:
		addr = settingsPatchSidFilterV3EnOffset
	case shared.VoiceEx:
		addr = settingsPatchSidFilterExEnOffset
	}
	return shared.IToB(getValue[uint8](m.patchMem[chip][sidPatchData], addr, size))
}

func (m *StorageDevice) SetFilterRes(chip shared.SidChip, resonance uint8) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSidFilterResOffset, resonance, settingsPatchSidFilterResBits)
}

func (m *StorageDevice) ResetFilterRes(chip shared.SidChip) {
	m.SetFilterRes(chip, settingsPatchSidFilterResDefault)
}

func (m *StorageDevice) GetFilterRes(chip shared.SidChip) uint8 {
	return getValue[uint8](m.patchMem[chip][sidPatchData], settingsPatchSidFilterResOffset, settingsPatchSidFilterResBits)
}

func (m *StorageDevice) SetFilterCutoff(chip shared.SidChip, cutoff uint16) {
	setValue(&m.patchMem[chip][sidPatchData], settingsPatchSidFilterCutoffOffset, cutoff, settingsPatchSidFilterCutoffBits)
}

func (m *StorageDevice) ResetFilterCutoff(chip shared.SidChip) {
	m.SetFilterCutoff(chip, settingsPatchSidFilterCutoffDefault)
}

func (m *StorageDevice) GetFilterCutoff(chip shared.SidChip) uint16 {
	return getValue[uint16](m.patchMem[chip][sidPatchData], settingsPatchSidFilterCutoffOffset, settingsPatchSidFilterCutoffBits)
}

func (m *StorageDevice) SetVoiceSync(chip shared.SidChip, voice shared.VoiceIndex, sync bool) {
	addr := uint8(settingsPatchVSyncOffset)
	size := uint8(settingsPatchVSyncBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, shared.BToI(sync), size)
}

func (m *StorageDevice) ResetVoiceSync(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceSync(chip, voice, shared.IToB(settingsPatchVSyncDefault))
}

func (m *StorageDevice) GetVoiceSync(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchVSyncOffset)
	size := uint8(settingsPatchVSyncBits)
	bank := patchBank(voice + 1)

	return shared.IToB(getValue[uint8](m.patchMem[chip][bank], addr, size))
}

func (m *StorageDevice) SetVoiceRingMod(chip shared.SidChip, voice shared.VoiceIndex, sync bool) {
	addr := uint8(settingsPatchVRingModOffset)
	size := uint8(settingsPatchVRingModBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, shared.BToI(sync), size)
}

func (m *StorageDevice) ResetVoiceRingMod(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceRingMod(chip, voice, shared.IToB(settingsPatchVRingModDefault))
}

func (m *StorageDevice) GetVoiceRingMod(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchVRingModOffset)
	size := uint8(settingsPatchVRingModBits)
	bank := patchBank(voice + 1)

	return shared.IToB(getValue[uint8](m.patchMem[chip][bank], addr, size))
}

func (m *StorageDevice) SetVoiceTriangle(chip shared.SidChip, voice shared.VoiceIndex, sync bool) {
	addr := uint8(settingsPatchVTriangleOffset)
	size := uint8(settingsPatchVTriangleBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, shared.BToI(sync), size)
}

func (m *StorageDevice) ResetVoiceTriangle(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceTriangle(chip, voice, shared.IToB(settingsPatchVTriangleDefault))
}

func (m *StorageDevice) GetVoiceTriangle(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchVTriangleOffset)
	size := uint8(settingsPatchVTriangleBits)
	bank := patchBank(voice + 1)

	return shared.IToB(getValue[uint8](m.patchMem[chip][bank], addr, size))
}

func (m *StorageDevice) SetVoiceSawtooth(chip shared.SidChip, voice shared.VoiceIndex, sync bool) {
	addr := uint8(settingsPatchVSawtoothOffset)
	size := uint8(settingsPatchVSawtoothBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, shared.BToI(sync), size)
}

func (m *StorageDevice) ResetVoiceSawtooth(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceSawtooth(chip, voice, shared.IToB(settingsPatchVSawtoothDefault))
}

func (m *StorageDevice) GetVoiceSawtooth(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchVSawtoothOffset)
	size := uint8(settingsPatchVSawtoothBits)
	bank := patchBank(voice + 1)

	return shared.IToB(getValue[uint8](m.patchMem[chip][bank], addr, size))
}

func (m *StorageDevice) SetVoicePulse(chip shared.SidChip, voice shared.VoiceIndex, sync bool) {
	addr := uint8(settingsPatchVPulseOffset)
	size := uint8(settingsPatchVPulseBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, shared.BToI(sync), size)
}

func (m *StorageDevice) ResetVoicePulse(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoicePulse(chip, voice, shared.IToB(settingsPatchVPulseDefault))
}

func (m *StorageDevice) GetVoicePulse(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchVPulseOffset)
	size := uint8(settingsPatchVPulseBits)
	bank := patchBank(voice + 1)

	return shared.IToB(getValue[uint8](m.patchMem[chip][bank], addr, size))
}

func (m *StorageDevice) SetVoiceNoise(chip shared.SidChip, voice shared.VoiceIndex, sync bool) {
	addr := uint8(settingsPatchVNoiseOffset)
	size := uint8(settingsPatchVNoiseBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, shared.BToI(sync), size)
}

func (m *StorageDevice) ResetVoiceNoise(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceNoise(chip, voice, shared.IToB(settingsPatchVNoiseDefault))
}

func (m *StorageDevice) GetVoiceNoise(chip shared.SidChip, voice shared.VoiceIndex) bool {
	addr := uint8(settingsPatchVNoiseOffset)
	size := uint8(settingsPatchVNoiseBits)
	bank := patchBank(voice + 1)

	return shared.IToB(getValue[uint8](m.patchMem[chip][bank], addr, size))
}

func (m *StorageDevice) SetVoicePW(chip shared.SidChip, voice shared.VoiceIndex, pulsewidth uint16) {
	addr := uint8(settingsPatchVPWOffset)
	size := uint8(settingsPatchVPWBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, pulsewidth, size)
}

func (m *StorageDevice) ResetVoicePW(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoicePW(chip, voice, settingsPatchVPWDefault)
}

func (m *StorageDevice) GetVoicePW(chip shared.SidChip, voice shared.VoiceIndex) uint16 {
	addr := uint8(settingsPatchVPWOffset)
	size := uint8(settingsPatchVPWBits)
	bank := patchBank(voice + 1)

	return getValue[uint16](m.patchMem[chip][bank], addr, size)
}

func (m *StorageDevice) SetVoiceAttack(chip shared.SidChip, voice shared.VoiceIndex, attack shared.AttackRate) {
	addr := uint8(settingsPatchVAttackOffset)
	size := uint8(settingsPatchVAttackBits)
	bank := patchBank(voice + 1)

	setValue[shared.AttackRate](&m.patchMem[chip][bank], addr, attack, size)
}

func (m *StorageDevice) ResetVoiceAttack(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceAttack(chip, voice, settingsPatchVAttackDefault)
}

func (m *StorageDevice) GetVoiceAttack(chip shared.SidChip, voice shared.VoiceIndex) shared.AttackRate {
	addr := uint8(settingsPatchVAttackOffset)
	size := uint8(settingsPatchVAttackBits)
	bank := patchBank(voice + 1)

	return getValue[shared.AttackRate](m.patchMem[chip][bank], addr, size)
}

func (m *StorageDevice) SetVoiceDecay(chip shared.SidChip, voice shared.VoiceIndex, decay shared.DecayRate) {
	addr := uint8(settingsPatchVDecayOffset)
	size := uint8(settingsPatchVDecayBits)
	bank := patchBank(voice + 1)

	setValue[shared.DecayRate](&m.patchMem[chip][bank], addr, decay, size)
}

func (m *StorageDevice) ResetVoiceDecay(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceDecay(chip, voice, settingsPatchVDecayDefault)
}

func (m *StorageDevice) GetVoiceDecay(chip shared.SidChip, voice shared.VoiceIndex) shared.DecayRate {
	addr := uint8(settingsPatchVDecayOffset)
	size := uint8(settingsPatchVDecayBits)
	bank := patchBank(voice + 1)

	return getValue[shared.DecayRate](m.patchMem[chip][bank], addr, size)
}

func (m *StorageDevice) SetVoiceSustain(chip shared.SidChip, voice shared.VoiceIndex, sustain uint8) {
	addr := uint8(settingsPatchVSustainOffset)
	size := uint8(settingsPatchVSustainBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, sustain, size)
}

func (m *StorageDevice) ResetVoiceSustain(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceSustain(chip, voice, settingsPatchVSustainDefault)
}

func (m *StorageDevice) GetVoiceSustain(chip shared.SidChip, voice shared.VoiceIndex) uint8 {
	addr := uint8(settingsPatchVSustainOffset)
	size := uint8(settingsPatchVSustainBits)
	bank := patchBank(voice + 1)

	return getValue[uint8](m.patchMem[chip][bank], addr, size)
}

func (m *StorageDevice) SetVoiceRelease(chip shared.SidChip, voice shared.VoiceIndex, release shared.ReleaseRate) {
	addr := uint8(settingsPatchVReleaseOffset)
	size := uint8(settingsPatchVReleaseBits)
	bank := patchBank(voice + 1)

	setValue[shared.ReleaseRate](&m.patchMem[chip][bank], addr, release, size)
}

func (m *StorageDevice) ResetVoiceRelease(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceRelease(chip, voice, settingsPatchVReleaseDefault)
}

func (m *StorageDevice) GetVoiceRelease(chip shared.SidChip, voice shared.VoiceIndex) shared.ReleaseRate {
	addr := uint8(settingsPatchVReleaseOffset)
	size := uint8(settingsPatchVReleaseBits)
	bank := patchBank(voice + 1)

	return getValue[shared.ReleaseRate](m.patchMem[chip][bank], addr, size)
}

func (m *StorageDevice) SetVoiceDetune(chip shared.SidChip, voice shared.VoiceIndex, cents int8) {
	addr := uint8(settingsPatchVDetuneCentsOffset)
	size := uint8(settingsPatchVDetuneCentsBits)
	bank := patchBank(voice + 1)

	setValue(&m.patchMem[chip][bank], addr, cents, size)
}

func (m *StorageDevice) ResetVoiceDetune(chip shared.SidChip, voice shared.VoiceIndex) {
	m.SetVoiceDetune(chip, voice, settingsPatchVDetuneCentsDefault)
}

func (m *StorageDevice) GetVoiceDetune(chip shared.SidChip, voice shared.VoiceIndex) int8 {
	addr := uint8(settingsPatchVDetuneCentsOffset)
	size := uint8(settingsPatchVDetuneCentsBits)
	bank := patchBank(voice + 1)

	return getValue[int8](m.patchMem[chip][bank], addr, size)
}
