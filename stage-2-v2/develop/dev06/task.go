package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func (cS *cutStruct) cut(rows []string) []string {
	type colStruct struct {
		haveDelimter bool
		col          string
	}

	var result []string

	// проходимся по строкам
	for _, row := range rows {
		var cols []colStruct
		var haveDelimter bool // есть ли в строке разделитель
		// разделяем строку на столбцы
		cols = append(cols, colStruct{
			haveDelimter: false,
			col:          "",
		})

		for _, r := range row {
			// если текущий символ не разделитель, добавляем в предыдущий столбей
			if r != cS.flags.d {
				cols[len(cols)-1].col += string(r)
				continue
			}
			// если текущий символ разделитель, начинаем новый столбец, а в старом ставим флаг
			cols[len(cols)-1].haveDelimter = true
			haveDelimter = true
			cols = append(cols, colStruct{
				haveDelimter: false,
				col:          "",
			})
		}

		index := cS.flags.f - 1

		// fmt.Println(cols)
		// если такой столбец есть
		if len(cols) > index {
			// если задан флаг печатать только с разделителями и есть разделитель
			// или это последний столбец
			if cS.flags.s && cols[index].haveDelimter || !cS.flags.s || (len(cols)-1 == index && haveDelimter) {
				result = append(result, cols[index].col)
			}
			continue
		}
		// если 1 столбец и не срезаем столбцы, показываем его
		if !cS.flags.s && len(cols) == 1 {
			result = append(result, cols[0].col)
		}
		// если такого столбца нет
		if haveDelimter {
			result = append(result, "")
			continue
		}
	}

	return result
}

type flagsStrusct struct {
	f int
	d rune
	s bool
}

type cutStruct struct {
	flags flagsStrusct
}

func main() {
	flag_f := flag.Int("f", 1, "")
	flag_d := flag.String("d", "	", "")
	flag_s := flag.Bool("s", false, "")
	flag.Parse()

	if !isFlagPassed("f") {
		fmt.Println("флаг f должен быть задан")
		return
	}

	flag_d_rune := []rune(*flag_d)[0]

	cS := cutStruct{
		flags: flagsStrusct{
			f: *flag_f,
			d: flag_d_rune,
			s: *flag_s,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for _, row := range cS.cut(input) {
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

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
