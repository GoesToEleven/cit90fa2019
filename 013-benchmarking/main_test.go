package main

import "testing"

func BenchmarkComplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compLit()
	}
}

func BenchmarkMakeway(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeWay()
	}
}
