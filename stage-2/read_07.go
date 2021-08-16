// Что выведет программа? Объяснить вывод программы.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	var counter int = 2
	go func() {
		for {
			select {
			// case v := <-a:
			// 	c <- v
			// case v := <-b:
			// 	c <- v
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					counter--
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					counter--
				}
			}
			if counter == 0 {
				close(c)
				break
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

// asChan из слайса int пишет данные в канал
// merge из данные из 2-х каналов объединяет в один канал
// в основном потоке данные из объединенного канала читаются

// проблема в коде в том, что при чтении из канала не проверяется закрыт канал или нет
// и после того как они закрыты он оттуда читает стандартные значения типа, для int это 0
// если оба канала закрыты, надо закрыть и исходящий. код переделан
