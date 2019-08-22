package main

import "testing"

func TestFoo(t *testing.T) {
	want := "hello world"
	if got := foo(); got != want {
		t.Errorf("Foo() = %q, want %q", got, want)
	}
}
