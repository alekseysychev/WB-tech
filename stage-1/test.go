package main

import (
	"fmt"
	"runtime"
	"strings"
)

var s0 string // a package-level variable

func f(s1 string) {
	// s0 = s1[:50]
	copy([]byte(s0), s1[:50])
}

func createStringWithLengthOnHeap(len int) string {
	var s strings.Builder
	for i := 0; i < len; i++ {
		s.WriteString("a")
	}

	return s.String()
}

func a() {
	s := createStringWithLengthOnHeap(1 << 30) // 1M bytes
	f(s)
}

func main() {
	// a()
	s := createStringWithLengthOnHeap(1 << 30) // 1M bytes
	runtime.GC()

	fmt.Println("load")
	var input string
	fmt.Scanf("%d", input)
	fmt.Println(s)
}
