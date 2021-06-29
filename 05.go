package main

// Написать программу, которая будет последовательно
// писать значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершиться.

import (
	"fmt"
	"time"
)

func main() {
	var delay int = 1 // время выполнения
	channel := make(chan string)
	go func() {
		// канал дляя ззавершения через N секунд
		e := time.After(time.Second * time.Duration(delay))
		// таймер каждую секуну для записи в канал
		t := time.NewTicker(time.Second / 10)
		// по завершении горутины останавливаем таймер
		defer t.Stop()
		for {
			select {
			case <-t.C:
				// пишем в канал
				channel <- "tick"
			case <-e:
				// закрываем канал и выходим из горутины, если время истекло
				close(channel)
				return
			}
		}
	}()
	// читаем данные из канала
	for c := range channel {
		fmt.Printf("%s ", c)
	}
	fmt.Printf("\n")
}
