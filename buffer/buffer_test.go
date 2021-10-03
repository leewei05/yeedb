package buffer

import "testing"

func TestSize(t *testing.T) {
	frames := []int{1, 2, 3, 4, 5, 6}
	cr := &ClockReplacer{frames}

	got := cr.Size()
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
