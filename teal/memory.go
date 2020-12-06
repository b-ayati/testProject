package teal

import (
	"fmt"
)

const (
	InitialSize     = 10
	ExpansionFactor = 4
)

type MemorySegment struct {
	segment     []DataType
	snapManager snapshotManager
}

func NewMemorySegment() *MemorySegment {
	return &MemorySegment{segment: make([]DataType, InitialSize)}
}

func (ms *MemorySegment) AllocateAt(index int, item DataType) {
	if index >= len(ms.segment) {
		ms.Expand(ExpansionFactor)
	}
	if ms.segment[index] != nil {
		panic("Memory unit at ... is not empty.")
	}
	item.setSnapshotManager(&ms.snapManager)
	ms.snapManager.notifyUpdate(&ms.segment[index], ms.segment[index])
	ms.segment[index] = item
}

func (ms *MemorySegment) Expand(factor float32) {
	newSize := int(1 + float32(len(ms.segment))*(1+factor))
	newSegment := make([]DataType, newSize)

	snapshots := ms.snapManager.savedSnapshots
	if len(snapshots) > 0 {
		for i := range ms.segment {
			if oldValue, exists := snapshots[&ms.segment[i]]; exists {
				snapshots[&newSegment[i]] = oldValue
				delete(snapshots, &ms.segment[i])
			}
		}
	}

	for i, d := range ms.segment {
		newSegment[i] = d
	}

	ms.segment = newSegment
}

func (ms *MemorySegment) Compact() {
	if len(ms.snapManager.savedSnapshots) > 0 {
		return
	}
	last := len(ms.segment) - 1
	for ; last >= 0 && ms.segment[last] == nil; last-- {
	}
	println(last)

	newSegment := make([]DataType, last+1)
	for i := range newSegment {
		newSegment[i] = ms.segment[i]
	}
	ms.segment = newSegment
}

func (ms *MemorySegment) Free(index int) {
	ms.snapManager.notifyUpdate(&ms.segment[index], ms.segment[index])
	ms.segment[index] = nil
}

func (ms *MemorySegment) Get(index int) DataType {
	if d := ms.segment[index]; d != nil {
		return d
	} else {
		panic("segmentation fault!")
	}
}

func (ms *MemorySegment) SaveSnapshot() {
	ms.snapManager.reset()
}

//DiscardSnapshot stops the MemorySegment from
func (ms *MemorySegment) DiscardSnapshot() {
	ms.snapManager.turnOff()
}

func (ms *MemorySegment) RestoreSnapshot() {
	ms.snapManager.restoreSnapshot()
}

func (ms *MemorySegment) String() string {
	return fmt.Sprintf("memory:%v\nsnapshots:%v\n", ms.segment, ms.snapManager.savedSnapshots)
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
		case *uint64:
			*p = value.(uint64)
		case *byte:
			*p = value.(byte)
		case *bool:
			*p = value.(bool)
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
