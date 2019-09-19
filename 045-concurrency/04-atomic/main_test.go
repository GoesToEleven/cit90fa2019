package main

import "testing"

func BenchmarkLaunchRoutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRoutines()
	}
}
// 705524 ns/op
