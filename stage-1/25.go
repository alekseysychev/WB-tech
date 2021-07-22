package main

import (
	"fmt"
	"sync"
)

// Написать свою структуру счетчик, которая будет инкрементировать и выводить значения в конкурентной среде.

type counter struct {
	i int
	sync.Mutex
}

func (s *counter) show() {
	s.Lock()
	fmt.Printf("%d ", s.i)
	s.i++
	s.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	item := &counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				item.show()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\n")
}
