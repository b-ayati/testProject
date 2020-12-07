package teal

import "fmt"

type DataType interface {
	setSnapshotManager(*snapshotManager)
	//String() string
}

type UInt struct {
	manager *snapshotManager
	value   uint64
}

func NewUInt(value uint64) *UInt {
	return &UInt{value: value}
}

func (i *UInt) setSnapshotManager(sm *snapshotManager) {
	i.manager = sm
}

func (i *UInt) Value() uint64 {
	return i.value
}

func (i *UInt) SetValue(value uint64) {
	if i.manager != nil {
		i.manager.notifyUpdate(&i.value, i.value)
	}
	i.value = value
}

/*
func (i *UInt) String() string {
	return fmt.Sprintf("TealUInt: %d", i.value)
}*/

//OutOfRangeError is an error type indicating that some integer value is out of its valid range.
type OutOfRangeError struct {
	value       int
	lowerBound  int
	higherBound int
}

func (orErr *OutOfRangeError) Error() string {
	return fmt.Sprintf("%d is out of range. it must be between %d and %d", orErr.value, orErr.lowerBound, orErr.higherBound)
}
