package main

// Дана последовательность температурных колебаний
// (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5).
// Объединить данный значения в группы с шагов в 10 градусов.
// Последовательность в подмножностве не важна.

import (
	"fmt"
)

func main() {
	result := make(map[int][]float32)
	t := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, c := range t {
		cInt := int(c)
		index := cInt - cInt%10
		result[index] = append(result[index], c)

	}
	fmt.Printf("%+v\n", result)
}
