package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const testString = "test"

func BenchmarkConcat(b *testing.B) {
	var str string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += testString
	}
	b.StopTimer()
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString(testString)
	}
	b.StopTimer()
}

func BenchmarkSpririntf(b *testing.B) {
	var str string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str = fmt.Sprintf("%s %s", str, testString)
	}
	b.StopTimer()
}

func BenchmarkJoin(b *testing.B) {
	var str string
	var arr []string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		arr = append(arr, testString)
	}
	str = strings.Join(arr, " ")
	_ = str
	b.StopTimer()
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], testString)
	}
	b.StopTimer()
}
