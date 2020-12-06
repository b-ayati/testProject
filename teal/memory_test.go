package teal

import "testing"

func TestSnapshotManager(t *testing.T) {
	got := "bad bad"
	want := "good good"
	check(got, want, t)
}

func check(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("\nWhile running [%v]:\n we wanted:[%v]\nbut we got:[%v]", t.Name(), want, got)
	}
}
