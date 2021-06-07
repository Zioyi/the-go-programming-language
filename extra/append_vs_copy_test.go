package main

import (
	"crypto/rand"
	"testing"
)

var src = make([]byte, 512)
var dst = make([]byte, 512)

func genSource() {
	rand.Read(src)
}

func BenchmarkCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		genSource()
		b.StartTimer()
		copy(dst, src)
	}
}

func BenchmarkAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		genSource()
		b.StartTimer()
		dst = append(dst, src...)
	}
}
