package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "А"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:11]
}

func main() {
	someFunc()
	fmt.Printf("str %s\n", justString)
	fmt.Printf("len/cap %d/%d\n", len(justString), cap([]byte(justString)))
}
