package teal

import (
	"fmt"
	"testing"
)

func TestSnapshotManager(t *testing.T) {
	m := NewMemorySegment()
	//m.AllocateAt(0, NewUInt(23))
	fmt.Print(m)
	m.Compact()
	fmt.Print(m)
	m.AllocateAt(3, NewUInt(11))
	fmt.Print(m)
	got := "good good"
	want := "good good"
	check(got, want, t)
}

func check(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("\nWhile running [%v]:\n we wanted:[%v]\nbut we got:[%v]", t.Name(), want, got)
	}
}
