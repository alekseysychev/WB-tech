package main

// Какие существуют способы остановить выполнения горутины?
// Написать примеры использования.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Завершится по достижении конца кода")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Завершится при досрочном возврате данных из функции")
		return
		fmt.Println("не отобразится")
	}()

	// контекст
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("ctx ")
			}
			time.Sleep(time.Second / 10)
		}
	}(ctx)
	time.Sleep(time.Second)
	cancel()
	fmt.Printf("\n")

	// семафор
	wg.Add(1)
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-ch:
				return
			default:
				fmt.Printf("sem ")
			}
			time.Sleep(time.Second / 10)
		}
	}(ch)
	time.Sleep(time.Second)
	close(ch)

	wg.Wait()
	fmt.Printf("\n")
}
