package main

// Реализовать конкурентную запись в map.

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mx sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mx.Lock()
			m[1] = 1
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(m)
}
