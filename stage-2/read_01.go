package main

// Что выведет программа? Объяснить вывод программы.

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

// слайсинг [1:4] делает выбор включая элемент с индексом 1(77) и не включая элемент с индексом 4(80)
// математическая запись [1:4)
