package main

import "fmt"

// Написать программу, которая проверяет, что все символы в строке уникальные.

func check(s string) bool {
	fr := make(map[rune]struct{})

	for _, r := range []rune(s) {
		if _, ok := fr[r]; ok {
			return false
		}
		fr[r] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(check("absdefg"))
	fmt.Println(check("test"))
}
