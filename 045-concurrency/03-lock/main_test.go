package main

import "testing"

func BenchmarkLaunchRoutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRoutines()
	}
}
// 711282 ns/op
