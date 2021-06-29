package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	f := make(map[string]struct{})
	for _, v := range s {
		if _, ok := f[v]; !ok {
			f[v] = struct{}{}
		}
	}

	fmt.Println(f)
}
