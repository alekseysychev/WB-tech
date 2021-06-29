package main

import (
	"fmt"
	"sort"
)

// Написать бинарный поиск встроенными методами языка.

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sort.SearchInts(a, 5))
}
