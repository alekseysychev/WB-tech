/*
Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

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
	"sort"
	"strconv"
	"strings"
)

type sortStruct struct {
	flags map[string]interface{}
}

func main() {
	flag_k := flag.Int("k", -1, "указание колонки для сортировки") // происходит сортировка по этой и следующих колонках
	flag_n := flag.Bool("n", false, "сортировать по числовому значению")
	flag_r := flag.Bool("r", false, "сортировать в обратном порядке")
	flag_u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag_M := flag.Bool("M", false, "сортировать по названию месяца «JAN» < … < «DEC»")
	flag_b := flag.Bool("b", false, "игнорировать начальные пропуски")
	flag_c := flag.Bool("c", false, "проверять отсортированы ли данные")
	flag_h := flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")

	flag.Parse()

	s := &sortStruct{
		flags: make(map[string]interface{}),
	}

	if isFlagPassed("k") {
		s.flags["k"] = *flag_k
	}
	if isFlagPassed("n") {
		s.flags["n"] = *flag_n
	}
	if isFlagPassed("r") {
		s.flags["r"] = *flag_r
	}
	if isFlagPassed("u") {
		s.flags["u"] = *flag_u
	}
	if isFlagPassed("M") {
		s.flags["M"] = *flag_M
	}
	if isFlagPassed("b") {
		s.flags["b"] = *flag_b
	}
	if isFlagPassed("c") {
		s.flags["c"] = *flag_c
	}
	if isFlagPassed("h") {
		s.flags["h"] = *flag_h
	}

	// чтение пути до файла
	if len(flag.Args()) != 1 {
		fmt.Printf("%s\n", errors.New("не верно задан входящий файл"))
		return
	}
	filePath := flag.Args()[0]

	// чтение строк из файла
	input, err := getRowsFromFile(filePath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	output := s.sort(input)

	// вывод результирующих строк
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}

// сравнение слайсов строк
func compareRows(input, output []string) *int {
	for i := 0; i < len(input); i++ {
		if input[i] != output[i] {
			// fmt.Println(input[i], "!=", output[i])
			return &i
		}
	}
	return nil
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

func (s *sortStruct) sort(input []string) []string {
	months := map[string]int{
		"jan": 0, "feb": 1, "mar": 2, "apr": 3, "may": 4, "jun": 5, "jul": 6, "aug": 7, "sep": 8, "oct": 9, "nov": 10, "dec": 11,
	}

	// структура сортируемой строки
	type rowStruct struct {
		index int      // индекс исходной строки
		value string   // обработанное значение исходной строки
		parts []string // части значений исходной строки (если идёт сортировка по столбцам)
	}

	// преобразуем входные данные к внутренней структуре
	var local []rowStruct
	for i := 0; i < len(input); i++ {
		local = append(local, rowStruct{
			index: i,
			value: input[i],
		})
	}

	// Флаг b игнорировать начальные пропуски
	// делаем trim всех строк перед дальнейшей обработкой
	if _, ok := s.flags["b"]; ok {
		for i := 0; i < len(local); i++ {
			local[i].value = strings.TrimLeft(local[i].value, " ")
		}
	}

	// делим строки на столбцы
	reg := regexp.MustCompile(`[ 	]`)
	for i := 0; i < len(local); i++ {
		local[i].parts = reg.Split(local[i].value, -1)
	}

	// Флаг k указание колонки для сортировки
	// если флаг указан, то оставляем из всех столбцов только этот
	// если такого столбца нет, то ставим пустое значение
	if k, ok := s.flags["k"]; ok {
		index := k.(int) - 1
		for i := 0; i < len(local); i++ {
			if index < len(local[i].parts) {
				local[i].parts = []string{local[i].parts[index]}
			} else {
				local[i].parts = []string{""}
			}
		}
		// fmt.Println(local)
	}

	// сортируем по байтово или по столбцам в зависимости от флагов
	sort.Slice(local, func(i, j int) bool {
		// сравниваем части
		for z := 0; z < len(local[i].parts); z++ {
			left, right := local[i].parts[z], local[j].parts[z]

			// Флаг h сортировать по числовому значению с учётом суффиксов
			// не знаю верного решения, потому тупо заменю буковки на нолики и сравню числа
			if _, ok := s.flags["h"]; ok {
				left, right = strings.Replace(left, "G", "000000000", -1), strings.Replace(right, "G", "000000000", -1)
				left, right = strings.Replace(left, "M", "000000", -1), strings.Replace(right, "M", "000000", -1)
				left, right = strings.Replace(left, "K", "000", -1), strings.Replace(right, "K", "000", -1)
				leftInt, err1 := strconv.Atoi(left)
				rightInt, err2 := strconv.Atoi(right)

				if err1 == nil && err2 == nil {
					if leftInt < rightInt {
						return true
					}
					if leftInt > rightInt {
						return false
					}
				}
			} else
			// Флаг M сортировать по названию месяца
			// преобразуем месяца к числам и их уже сравниваем
			if _, ok := s.flags["M"]; ok {
				leftLower, rightLower := strings.ToLower(left), strings.ToLower(right)
				leftInt, ok1 := months[leftLower]
				rightInt, ok2 := months[rightLower]
				if ok1 && ok2 {
					if leftInt < rightInt {
						return true
					}
					if leftInt > rightInt {
						return false
					}
				}
			} else
			// Флаг n сортировать по числовому значению
			// пытаемся преобразовать части к числам
			if _, ok := s.flags["n"]; ok {
				if left == "" {
					left = "0"
				}
				if right == "" {
					right = "0"
				}
				leftInt, err1 := strconv.Atoi(left)
				rightInt, err2 := strconv.Atoi(right)

				if err1 == nil && err2 == nil {
					if leftInt < rightInt {
						return true
					}
					if leftInt > rightInt {
						return false
					}
				}
			} else { // простое сравнение
				if left < right {
					return true
				}
				if left > right {
					return false
				}
			}
		}

		return len(local[i].parts) < len(local[j].parts)
	})

	// Флаг r сортировать в обратном порядке
	// перетасовываем строки, что бы были в обратном порядке
	if _, ok := s.flags["r"]; ok {
		for i, j := 0, len(local)-1; i < j; i, j = i+1, j-1 {
			local[i], local[j] = local[j], local[i]
		}
	}
	// Флаг u не выводить повторяющиеся строки
	// формируем массив из не повторяющихся строк. так как уже отсортировано, можем проверять только соседние
	if _, ok := s.flags["u"]; ok {
		var temp []rowStruct
		for i := 0; i < len(local); i++ {
			if len(temp) == 0 || temp[len(temp)-1].value != local[i].value {
				temp = append(temp, local[i])
			}
		}
		local = temp
	}
	// Флаг c проверять отсортированы ли данные
	// проверяем идут ли индексы строк друг за другом
	// if _, ok := s.flags["c"]; ok {
	// 	f := compareRows(s.input, local)
	// 	if f == nil {
	// 		local = []string{"данные отсортированы"}
	// 	} else {
	// 		local = []string{"данные не отсортированы, строка:" + strconv.Itoa(*f)}
	// 	}
	// }

	// собираем результат из входяще строки и сортированных индексов
	var result []string
	for i := 0; i < len(local); i++ {
		result = append(result, input[local[i].index])
	}
	return result
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
