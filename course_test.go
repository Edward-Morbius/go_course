package main

import "testing"

func BenchmarkSlowRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slowRepeat("Hello, World\n", 1000)
	}
}
