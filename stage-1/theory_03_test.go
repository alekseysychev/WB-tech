package main

import (
	"testing"
)

func BenchmarkVer1(b *testing.B) {
	m := make(map[int]struct{})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		{
			m[n] = struct{}{}
		}
	}
	b.StopTimer()
}

func BenchmarkVer2(b *testing.B) {
	m := make(map[int]bool)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		{
			m[n] = false
		}
	}
	b.StopTimer()
}
