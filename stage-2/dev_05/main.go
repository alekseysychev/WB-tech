/*
Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type flagsStrusct struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

type grepStruct struct {
	word  string
	flags flagsStrusct
}

func main() {
	flag_A := flag.Int("A", 0, `"after" печатать +N строк после совпадения`)
	flag_B := flag.Int("B", 0, `"before" печатать +N строк до совпадения`)
	flag_C := flag.Int("C", 0, `"context" (A+B) печатать ±N строк вокруг совпадения`)
	flag_c := flag.Bool("c", false, `"count" (количество строк)`)
	flag_i := flag.Bool("i", false, `"ignore-case" (игнорировать регистр)`)
	flag_v := flag.Bool("v", false, `"invert" (вместо совпадения, исключать)`)
	flag_F := flag.Bool("F", false, `"fixed", точное совпадение со строкой, не паттерн`)
	flag_n := flag.Bool("n", false, `"line num", печатать номер строки`)
	flag.Parse()

	gS := grepStruct{}

	gS.flags.A = *flag_A
	gS.flags.B = *flag_B
	gS.flags.C = *flag_C
	gS.flags.c = *flag_c
	gS.flags.i = *flag_i
	gS.flags.v = *flag_v
	gS.flags.F = *flag_F
	gS.flags.n = *flag_n

	if len(flag.Args()) != 2 {
		fmt.Printf("%s\n", errors.New("ошибка формата вызова команды"))
		return
	}

	args := flag.Args()
	gS.word = args[0]
	filePath := args[1]

	input, err := getRowsFromFile(filePath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	output := gS.grep(input)

	for _, row := range output {
		fmt.Println(row)
	}
}

func getRowsFromFile(path string) ([]string, error) {
	// открываем файл на чтение
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("code 01: %w", err)
	}
	// читаем строки в слайс
	rows := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, nil
}

func (gS *grepStruct) grep(input []string) []string {
	rowsMap := make([]byte, len(input))

	deltaA := gS.flags.C
	if gS.flags.A > 0 {
		deltaA = gS.flags.A
	}
	deltaB := gS.flags.C
	if gS.flags.B > 0 {
		deltaB = gS.flags.B
	}

	var pattern *regexp.Regexp
	if gS.flags.i {
		pattern, _ = regexp.Compile("(?i)" + gS.word)
	} else {
		pattern, _ = regexp.Compile(gS.word)
	}

	set := func(index int) {
		from := index - deltaB
		to := index + 1 + deltaA
		for i := from; i < to; i++ {
			if i >= 0 && i < len(rowsMap) {
				if i == index { // строка совпадения
					rowsMap[i] = 2
				}
				if rowsMap[i] == 0 { // ставим что это соседняя строка
					rowsMap[i] = 1
				}
			}
		}
	}

	for index, row := range input {
		switch {
		case gS.flags.F:
			if strings.Contains(row, gS.word) {
				set(index)
			}
		default:
			if pattern.MatchString(row) {
				set(index)
			}
		}
	}

	result := make([]string, 0)

	var appendDelimiter bool
	for index, rowMap := range rowsMap {
		var appendString string
		var appendStringBool bool
		switch {
		case gS.flags.n:
			if rowMap == 1 { // соседняя строка
				appendString = strconv.Itoa(index+1) + "-" + input[index]
				appendStringBool = true
			} else if rowMap == 2 { // точное совпадение
				appendString = strconv.Itoa(index+1) + ":" + input[index]
				appendStringBool = true
			}
		case gS.flags.v:
			if rowMap == 0 { // нет совпадения
				appendString = input[index]
				appendStringBool = true
			}
		default:
			if rowMap > 0 { // есть совпадение
				appendString = input[index]
				appendStringBool = true
			}
		}

		if appendStringBool {
			if (gS.flags.A > 0 || gS.flags.B > 0 || gS.flags.C > 0) && !gS.flags.c && appendDelimiter && len(result) > 0 {
				result = append(result, "--")
			}
			result = append(result, appendString)
			appendDelimiter = false
		} else {
			appendDelimiter = true
		}
	}

	if gS.flags.c {
		result = []string{strconv.Itoa(len(result))}
	}

	return result
}

// сравнение слайсов строк
func compareRows(input, output []string) *int {
	if len(input) != len(output) {
		i := 0
		return &i
	}
	for i := 0; i < len(input); i++ {
		if input[i] != output[i] {
			// fmt.Println(input[i], "!=", output[i])
			return &i
		}
	}
	return nil
}
