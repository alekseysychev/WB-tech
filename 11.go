package main

// Написать пересечение двух неупорядоченных массивов.

import "fmt"

func main() {
	a := []int{44, 1, 22, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{11, 22, 33, 1, 44, 22, 22}
	// карта частот
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	for _, v := range b {
		// если в карте есть такой элемент и у него ещё есть повторения, выводи
		if c, ok := freq[v]; ok && c > 0 {
			freq[v]--
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n")
}
