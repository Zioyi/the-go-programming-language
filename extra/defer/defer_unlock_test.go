package main

import (
	"math/rand"
	"sync"
	"testing"
)

var i int

func unlockDefer(l *sync.Mutex) {
	l.Lock()
	defer l.Unlock()
	i = rand.Int()
}

func BenchmarkUnlockDefer(b *testing.B) {
	l := &sync.Mutex{}
	for n := 0; n < b.N; n++ {
		unlockDefer(l)
	}
}

func unlockDirect(l *sync.Mutex) {
	l.Lock()
	i = rand.Int()
	l.Unlock()
}

func BenchmarkUnlockDirect(b *testing.B) {
	l := &sync.Mutex{}
	for n := 0; n < b.N; n++ {
		unlockDefer(l)
	}
}

