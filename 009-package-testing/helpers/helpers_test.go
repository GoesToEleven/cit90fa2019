package helpers

import "testing"

func TestBar(t *testing.T) {
	want := 42
	if got := Bar(); got != want {
		t.Errorf("Bar() = %d, want %d", got, want)
	}
}
