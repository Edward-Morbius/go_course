package main

import "testing"

func BenchmarkRepeat1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = repeat1("Hello, World\n", 1000)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = repeat2("Hello, World\n", 1000)
	}
}

func BenchmarkRepeat3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = repeat3("Hello, World\n", 1000)
	}
}
