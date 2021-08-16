// Создать программу печатающую точное время с использованием NTP библиотеки.
// Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
// Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

// Программа должна быть оформлена с использованием как go module.
// Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
// Программа должна проходить проверки go vet и golint.

package main

import (
	"fmt"
	"log"

	"example.com/customTime"
)

func main() {
	t := customTime.New()

	err := t.SetHost("0.beevik-ntp.pool.ntp.1org")
	if err != nil {
		log.Fatalln(err)
	}
	time, err := t.Time()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(time)
}
