package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "А" // многобайтовая кодировка
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:11]
}

func main() {
	someFunc()

	// так же проблема такого выделения может встретиться со слайсами, когда cap берётся от исходного слайса
	i1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	i2 := i1[:0]
	fmt.Printf("str %s\n", justString)
	fmt.Printf("len/cap %d/%d\n", len(i1), cap(i1))
	fmt.Printf("len/cap %d/%d\n", len(i2), cap(i2))
}
