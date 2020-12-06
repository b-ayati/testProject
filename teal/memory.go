package teal

import "fmt"

type MemorySegment struct {
	segment     []DataType
	snapManager snapshotManager
}

func (ms *MemorySegment) SaveSnapshot() {
	ms.snapManager.reset()
}

func (ms *MemorySegment) RestoreSnapshot() {
	ms.snapManager.restoreSnapshot()
}

func (ms *MemorySegment) Add(item DataType) {
	item.setSnapshotManager(&ms.snapManager)
	ms.segment = append(ms.segment, item)
}

func (ms *MemorySegment) String() string {
	return fmt.Sprintf("%s", ms.segment)
}

type snapshotManager struct {
	savedSnapshots map[interface{}]interface{}
}

func (sm *snapshotManager) reset() {
	sm.savedSnapshots = make(map[interface{}]interface{})
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
	if _, exists := sm.savedSnapshots[pointer]; !exists {
		sm.savedSnapshots[pointer] = oldValue
	} else {
		println("we have", pointer, oldValue)
	}
}
