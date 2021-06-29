package main

import (
	"fmt"
)

func main() {
	var a map[int]int
	b := map[int]int{}
	c := new(map[int]int)
	d := make(map[int]int)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(d)

	b[0] = 0

	d[0] = 0

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(d)
}
