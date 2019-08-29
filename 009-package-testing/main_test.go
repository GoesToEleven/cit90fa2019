package main

import "testing"

func TestFoo(t *testing.T) {
	want := "hello"
	if got := foo(); got != want {
		t.Errorf("foo() = %s, want %s", got, want)
	}
}
