package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(wg sync.WaitGroup, i int) {
				fmt.Println(i)
				wg.Done()
			}(wg, i)
		}
		wg.Wait()
		fmt.Println("exit")
	*/
	// вариант 1
	{
		wg := &sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, i int) {
				fmt.Println(i)
				wg.Done()
			}(wg, i)
		}
		wg.Wait()
		fmt.Println("exit")
	}
	// вариант 2
	{
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(i int) {
				fmt.Println(i)
				wg.Done()
			}(i)
		}
		wg.Wait()
		fmt.Println("exit")
	}
}
