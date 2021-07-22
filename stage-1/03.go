package main

// Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(22+32+42….)
// с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	var result int32
	var resultWrong int32

	for _, v := range []int{2, 4, 6, 8, 10, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3} {
		wg.Add(1)
		go func(v int) {
			atomic.AddInt32(&result, int32(v*v))
			resultWrong += int32(v * v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("result :      ", atomic.LoadInt32(&result))
	fmt.Println("resultWrong : ", resultWrong)
}
