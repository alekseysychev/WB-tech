package main

import (
	"fmt"
	"strings"
)

// Написать программу, которая переворачивает слова в строке (snow dog sun - sun dog snow).

func main() {
	str1 := "snow1 dog2 sun3 - sun4 dog5 snow6"
	// разделение на подстроки
	str2 := strings.Split(str1, " ")

	for i, j := 0, len(str2)-1; i < j; i, j = i+1, j-1 {
		str2[i], str2[j] = str2[j], str2[i]
	}
	// объеденение обратно
	str3 := strings.Join(str2, " ")

	fmt.Printf("Было: %s\n", str1)
	fmt.Printf("Стало: %s\n", str3)
}
