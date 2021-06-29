package main

import (
	"fmt"
	"sync"
	"time"
)

// Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	current := 0
	// mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	reader := func(readerNumber int) {
		defer wg.Done()
		for {
			// mu.Lock()
			if len(a) <= current {
				return
			}
			v := a[current]
			fmt.Printf("%d : %d\n", readerNumber, v)
			current++
			time.Sleep(time.Millisecond * time.Duration(v))
			// mu.Unlock()
		}
	}

	wg.Add(1)
	go reader(1)
	wg.Add(1)
	go reader(2)

	wg.Wait()
}
