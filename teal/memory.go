package teal

import (
	"fmt"
)

type MemorySegment struct {
	segment     []DataType
	snapManager snapshotManager
	maxSize     int
}

func NewMemorySegment(size int) *MemorySegment {
	return &MemorySegment{maxSize: size, segment: make([]DataType, size)}
}

func (ms *MemorySegment) AllocateAt(index int, item DataType) error {
	if index < 0 || index >= ms.maxSize {
		return &OutOfRangeError{value: index, lowerBound: 0, higherBound: ms.maxSize - 1}
	}
	if index >= len(ms.segment) {
		ms.expand()
	}
	if ms.segment[index] != nil {
		return fmt.Errorf("memory unit at %d is not empty", index)
	}
	//we need to notify our snapshot manager about change in segment[]
	ms.snapManager.notifyUpdate(&ms.segment[index], ms.segment[index])
	//adding item
	item.setSnapshotManager(&ms.snapManager)
	ms.segment[index] = item
	return nil
}

func (ms *MemorySegment) Free(index int) {
	//we need to notify our snapshot manager about change in segment[]
	ms.snapManager.notifyUpdate(&ms.segment[index], ms.segment[index])
	//removing item
	ms.segment[index] = nil
}

func (ms *MemorySegment) Get(index int) (DataType, error) {
	if index < 0 || index >= ms.maxSize {
		return nil, &OutOfRangeError{value: index, lowerBound: 0, higherBound: ms.maxSize - 1}
	}
	if index >= len(ms.segment) || ms.segment[index] == nil {
		return nil, fmt.Errorf("memory unit at %d is empty", index)
	}
	return ms.segment[index], nil
}

func (ms *MemorySegment) SaveSnapshot() {
	ms.snapManager.reset()
}

//DiscardSnapshot stops the MemorySegment from
func (ms *MemorySegment) DiscardSnapshot() {
	ms.snapManager.turnOff()
	ms.Compact()
}

func (ms *MemorySegment) RestoreSnapshot() {
	ms.snapManager.restoreSnapshot()
	//we don't need the old snapshot anymore, so we reset snapManager to improve performance of MemorySegment
	ms.snapManager.reset()
}

func (ms *MemorySegment) expand() {
	if len(ms.segment) == ms.maxSize {
		return
	}
	if len(ms.snapManager.savedSnapshots) > 0 {
		panic("We can not expand while there is a saved snapshot!")
	}
	newSegment := make([]DataType, ms.maxSize)
	copy(newSegment, ms.segment)
	ms.segment = newSegment
}

func (ms *MemorySegment) Compact() {
	if len(ms.snapManager.savedSnapshots) > 0 {
		return
	}
	//we need to find last element like this cuz we have a Free() function which can remove elements
	last := len(ms.segment) - 1
	for ; last >= 0 && ms.segment[last] == nil; last-- {
	}
	newSegment := make([]DataType, last+1)
	copy(newSegment, ms.segment)
	ms.segment = newSegment
}

func (ms *MemorySegment) String() string {
	str := fmt.Sprintf("Memory Segment: (maxSize:%d)", ms.maxSize)
	for i, data := range ms.segment {
		str += fmt.Sprintf("\n[%d, %T)]--->%v", i, data, data)
	}
	str += fmt.Sprintf("\nsavedSnapshots:%v\n", ms.snapManager.savedSnapshots)
	return str + "============================\n"
}

type snapshotManager struct {
	savedSnapshots map[interface{}]interface{}
}

func (sm *snapshotManager) reset() {
	sm.savedSnapshots = make(map[interface{}]interface{})
}

func (sm *snapshotManager) turnOff() {
	sm.savedSnapshots = nil
}

func (sm *snapshotManager) restoreSnapshot() {
	if sm.savedSnapshots == nil {
		panic("For restoring a snapshot u need to save one first!")
	}
	for pointer, value := range sm.savedSnapshots {
		switch p := pointer.(type) {
		case *DataType:
			if value == nil {
				*p = nil
			} else {
				*p = value.(DataType)
			}
		case *uint64:
			*p = value.(uint64)
		case *byte:
			*p = value.(byte)
		case *bool:
			*p = value.(bool)
		case *float64:
			*p = value.(float64)
		default:
			panic("It seems that you are trying to add a new teal.DataType but you forgot to change this function!")
		}
	}
}

func (sm *snapshotManager) notifyUpdate(pointer interface{}, oldValue interface{}) {
	//if sm.savedSnapshots is nil that means the snapshotManager is turned off and this function has no effect.
	if sm.savedSnapshots == nil {
		return
	}
	if _, exists := sm.savedSnapshots[pointer]; !exists {
		sm.savedSnapshots[pointer] = oldValue
	} else {
		println("we have", pointer, oldValue)
	}
}
