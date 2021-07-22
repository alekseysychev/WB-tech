package main

import (
	"fmt"
	"sync"
)

// Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются
// на четность и отправляются во второй канал. Результаты работы из второго канала пишутся в stdout.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	// канал 1
	wg.Add(1)
	go func(ch1 chan int, ch2 chan int) {
		for v := range ch1 {
			if v%2 == 0 {
				ch2 <- v
			}
		}
		// закрываем канал сообщая, что больше сообщений не будет
		close(ch2)
		wg.Done()
	}(ch1, ch2)
	// вывод
	wg.Add(1)
	go func(ch2 chan int) {
		for v := range ch2 {
			fmt.Printf("%d ", v)
		}
		wg.Done()
	}(ch2)
	// заприсываем в 1 канал числа
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		ch1 <- v
	}
	// закрываем канал сообщая, что больше сообщений не будет
	close(ch1)
	wg.Wait()
	fmt.Printf("\n")
}
