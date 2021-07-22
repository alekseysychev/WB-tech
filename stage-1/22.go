package main

import (
	"fmt"
	"sort"
)

// Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

type ByAsc []int

func (a ByAsc) Len() int           { return len(a) }
func (a ByAsc) Less(i, j int) bool { return a[i] < a[j] }
func (a ByAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	a := []int{1, 2, 3, 244, 53, 2, 12, 43, 35, 6, 54, 3, 2}

	fmt.Println(a)
	sort.Sort(ByAsc(a))
	fmt.Println(a)
}
