package teal

import (
	"bytes"
	"fmt"
)

type DataType interface {
	setSnapshotManager(*snapshotManager)
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

type ConstByteArray struct {
	values []byte
}

func NewConstByteArray(b []byte) *ConstByteArray {
	temp := make([]byte, len(b))
	copy(b, temp)
	return &ConstByteArray{values: temp}
}

func (cba *ConstByteArray) setSnapshotManager(sm *snapshotManager) {
	//do nothing!
}

func (cba *ConstByteArray) Get(i int) (byte, *OutOfRangeError) {
	if l := len(cba.values); i < 0 || i >= l {
		return 0, &OutOfRangeError{value: i, lowerBound: 0, higherBound: l - 1}
	}
	return cba.values[i], nil
}

func (cba *ConstByteArray) EqualsToSlice(b []byte) bool {
	return bytes.Equal(cba.values, b)
}

func (cba *ConstByteArray) Equals(other *ConstByteArray) bool {
	return bytes.Equal(cba.values, other.values)
}

type ByteArray struct {
	ConstByteArray
	manager *snapshotManager
}

func NewByteArray(size int) *ByteArray {
	return &ByteArray{ConstByteArray: ConstByteArray{values: make([]byte, size)}}
}

func (ba *ByteArray) setSnapshotManager(sm *snapshotManager) {
	ba.manager = sm
}

func (ba *ByteArray) Set(i int, b byte) *OutOfRangeError {
	if l := len(ba.values); i < 0 || i >= l {
		return &OutOfRangeError{value: i, lowerBound: 0, higherBound: l - 1}
	}
	if ba.manager != nil {
		ba.manager.notifyUpdate(&ba.values[i], ba.values[i])
	}
	ba.values[i] = b
	return nil
}

//OutOfRangeError is an error type indicating that some integer value is out of its valid range.
type OutOfRangeError struct {
	value       int
	lowerBound  int
	higherBound int
}

func (orErr *OutOfRangeError) Error() string {
	return fmt.Sprintf("%d is out of range. it must be between %d and %d", orErr.value, orErr.lowerBound, orErr.higherBound)
}
