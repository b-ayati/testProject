package teal

import (
	"fmt"
	"testing"
)

func TestSnapshotManager(t *testing.T) {

	m := NewMemorySegment(10)
	//m.AllocateAt(0, NewUInt(23))
	fmt.Print(m)
	m.Compact()
	fmt.Print(m)
	x := NewUInt(11)
	m.AllocateAt(0, x)
	fmt.Print(m)
	m.Compact()
	x.SetValue(15)
	fmt.Print(m)
	m.AllocateAt(6, NewUInt(51))
	fmt.Print(m)
	got := m.String()
	want := "Memory Segment: (maxSize:10)\n[0, *teal.UInt]--->&{0xc00005e348 15}\n[1, <nil>]\n[2, <nil>]\n[3, <nil>]\n[4, <nil>]\n[5, <nil>]\n[6, *teal.UInt]--->&{0xc00005e348 51}\n[7, <nil>]\n[8, <nil>]\n[9, <nil>]\nsavedSnapshots:map[]\n============================\n"
	check(got, want, t)
	check("bad", "good", t)
}

func check(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("\nWhile running [%v]:\nwe wanted:\n%v\nbut we got:\n%v", t.Name(), want, got)
	}
}
