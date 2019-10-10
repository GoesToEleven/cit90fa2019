package main

import "testing"

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
