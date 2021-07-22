package main

// Дана переменная int64. Написать программу которая
// устанавливает i-й бит в 1 или 0.

import (
	"fmt"
	"strconv"
)

func main() {
	var a int64 = 1234567890 // переменная Range: -9223372036854775808 through 9223372036854775807.
	var n int = 9            // i-ый бит
	var s string = "0"       // в 1 или 0
	{
		// число в двоичном виде в строке
		str1 := strconv.FormatInt(a, 2)
		// формируем новую строку
		str2 := str1[:len(str1)-n-1] + s + str1[len(str1)-n:]

		b, err := strconv.ParseInt(str2, 2, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Было  : (%s)  %d\n", str1, a)
			fmt.Printf("Стало : (%s)  %d\n", str2, b)
		}
		fmt.Println()
	}
	{
		b := a | 1<<n
		if s == "0" {
			b ^= 1 << n
		}

		fmt.Printf("Было  : (%s)  %d\n", strconv.FormatInt(a, 2), a)
		fmt.Printf("Стало : (%s)  %d\n", strconv.FormatInt(b, 2), b)
		fmt.Println()
	}
}
