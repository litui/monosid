package storage

import (
	"machine"
	"time"
)

type patchBank uint8

const (
	// Mapping for General data storage

	flashBlockWriteSize = 4096
)

const (
	sidPatchData patchBank = iota
	v1PatchData
	v2PatchData
	v3PatchData
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

var (
	memDev = machine.Flash
)

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

func New() StorageDevice {
	return StorageDevice{
		generalMem:      0,
		generalChanged:  false,
		generalLoaded:   false,
		lastGeneralSave: time.Time{},
		lastPatchSave:   time.Time{},
	}
}

func setValue[InputType Numerics](target *uint64, offset uint8, data InputType, length uint8) {
	for i := 0; i < int(length); i++ {
		bitVal := (uint64(data) >> i) & 1
		if bitVal == 1 {
			// Set
			*target |= uint64(1) << (offset + uint8(i))
		} else {
			// Clear
			*target &= ^(uint64(1) << (offset + uint8(i)))
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
