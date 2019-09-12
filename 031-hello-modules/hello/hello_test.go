package hello

import "testing"

func TestWorld(t *testing.T) {
	want := "hello world"
	if got := World(); got != want {
		t.Errorf("World() = %q, want %q", got, want)
	}
}
