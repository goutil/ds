package ds

import "testing"

func TestConnectedness(t *testing.T) {
	ds := NewDisjointSet(5)
	if ds.Connected(0, 4) {
		t.Fatalf("%v is not connected to %v", 0, 4)
	}

	ds.Connect(0, 1)
	if !ds.Connected(0, 1) {
		t.Fatalf("%v is connected to %v", 0, 1)
	}

	ds.Connect(1, 2)
	ds.Connect(2, 3)
	ds.Connect(3, 4)
	if !ds.Connected(0, 4) {
		t.Fatalf("%v is connected to %v", 0, 4)
	}
}
