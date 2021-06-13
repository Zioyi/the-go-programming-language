package main

import (
	"testing"
)

func BenchmarkDynamicSize(b *testing.B) {
	m := map[int]int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkPredefinedSize(b *testing.B) {
	m := make(map[int]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N+10; i++ {
		m[i] = i
	}
}
