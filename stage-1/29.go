package main

import (
	"fmt"
	"math/big"
)

// Написать программу, которая перемножает, делит, складывает,
// вычитает 2 числовых переменных a,b, значение которые > 2^20

func main() {
	x := 2 << 22 // первое число
	y := 2 << 30 // второе число

	{
		fmt.Println("easy")
		a := float64(x)
		b := float64(y)

		fmt.Println("+ ", a+b)
		fmt.Println("- ", a-b)
		fmt.Println("* ", a*b)
		fmt.Println("/ ", a/b)
	}
	{
		fmt.Println("math/big")
		a := big.NewFloat(float64(x))
		b := big.NewFloat(float64(y))

		fmt.Println("+ ", new(big.Float).Add(a, b))
		fmt.Println("- ", new(big.Float).Sub(a, b))
		fmt.Println("* ", new(big.Float).Mul(a, b))
		fmt.Println("/ ", new(big.Float).Quo(a, b))
	}
}
