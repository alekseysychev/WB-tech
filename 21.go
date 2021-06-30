package main

import (
	"fmt"
	"sync"
	"time"
)

// Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	current := 0
	wg := sync.WaitGroup{}

	// простое чтение с блокировкой счётика текущего элемента
	{
		currentMutex := sync.RWMutex{}
		reader := func(readerNumber int) {
			defer wg.Done()
			for {
				currentMutex.Lock()
				if len(a) <= current {
					currentMutex.Unlock()
					return
				}
				v := a[current]
				fmt.Printf("%d(%d) ", v, readerNumber)
				current++
				currentMutex.Unlock()
				time.Sleep(time.Millisecond * time.Duration(v))
			}
		}

		wg.Add(2)
		go reader(1)
		go reader(2)
		wg.Wait()
		fmt.Println()
	}
	// чтение через кусок слайса
	{
		reader := func(readerNumber int, sllice []int) {
			defer wg.Done()
			for _, v := range sllice {
				fmt.Printf("%d(%d) ", v, readerNumber)
			}
		}

		countReaders := 5
		for i := 0; i < countReaders; i++ {
			wg.Add(1)
			start := len(a) / countReaders * i
			stop := len(a) / countReaders * (i + 1)
			if i == countReaders-1 {
				stop = len(a)
			}
			go reader(i+1, a[start:stop])
		}
		wg.Wait()
		fmt.Println()
	}
}
