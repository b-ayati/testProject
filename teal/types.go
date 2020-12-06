package teal

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
