package main

import "fmt"

// Создать слайс с предварительно выделенными 100 элементами.

func main() {
	a := make([]int, 100)

	fmt.Println(len(a))
}
