package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

type NtpTime struct {
	Host string
}

func (n *NtpTime) Get() (time.Time, error) {
	t, err := ntp.Time(n.Host)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func main() {
	ntpTime := NtpTime{
		Host: "0.beevik-ntp.pool.ntp.org",
	}
	t, err := ntpTime.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(t)
}
