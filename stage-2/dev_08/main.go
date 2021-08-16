/*
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		command := sc.Text()
		switch {
		case len(command) > 5 && command[:5] == "echo ":
			fmt.Fprintln(os.Stdout, command[5:])
		case len(command) > 3 && command[:3] == "cd ":
			err := os.Chdir(command[3:])
			if err != nil {
				fmt.Println("задан не верный путь: ", command[:3])
				continue
			}
		case command == "pwd":
			path, err := filepath.Abs(".")
			if err != nil {
				fmt.Println("ошибка получения пути")
				continue
			}
			fmt.Println(path, err)
		default:
			fmt.Println("несуществующая команда")
		}
	}

	// dirs, err := os.ReadDir("./..")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))
	os.Chdir("./..")
	fmt.Println(filepath.Abs("."))

	// for _, dir := range dirs {
	// fmt.Println(dir.Info())
	// }

	// for {
	// }
}
