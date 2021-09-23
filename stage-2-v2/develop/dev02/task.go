package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"errors"
	"fmt"
	"strings"
)

func UnPack(s string) (string, error) {
	type p struct {
		r rune // символ
		c *int // количество дополнительных повторений
	}
	var store []p
	var escaped bool
	for _, r := range []rune(s) {
		storeCount := len(store)
		// если текущий символ экранирование и не экранирован
		if r == '\\' && !escaped {
			escaped = true
			continue
		}
		// если текущий символ экранирован предыдущим
		// если текущий символ не число
		if escaped || !(48 <= r && r <= 57) {
			store = append(store, p{
				r: r,
			})
			escaped = false
			continue
		}
		if storeCount == 0 {
			return "", errors.New("ошибка ввода. первый символ - число ")
		}
		// символ в число
		count := int(r - 48)
		// если предыдущий символ уже повторялся, увеличчиваем количество повторений
		// если нет , ставим количество повторений
		if store[storeCount-1].c != nil {
			*(store[storeCount-1].c) = *(store[storeCount-1].c)*10 + count
		} else {
			store[storeCount-1].c = &count
		}
	}

	// собираем строку из символов
	b := strings.Builder{}
	for _, v := range store {
		c := 1
		if v.c != nil {
			c = *v.c
		}
		for i := 0; i < c; i++ {
			b.WriteRune(v.r)
		}
	}

	return b.String(), nil
}

// func UnPack(s string) (string, error) {
// 	var stack []rune
// 	var escaped bool
// 	var i int
// 	var lastNum int

// 	// for _, r := range []rune(s) {
// 	for _, r := range s {
// 		// если текущий символ экранирован предыдущим, то мы добавляем его сразу
// 		if escaped {
// 			stack = append(stack, r)
// 			escaped = false
// 		} else
// 		// если текущий символ - экранирование
// 		if r == '\\' {
// 			escaped = true
// 		} else
// 		// если текущий символ - число
// 		if 48 <= r && r <= 57 {
// 			// длина результирующего слайса
// 			lenght := len(stack)
// 			// если текущий элемент первый
// 			if lenght == 0 {
// 				return "", errors.New("ошибка ввода. первый символ - число ")
// 			}
// 			// символ в число
// 			count := int(r - 48)
// 			// если перед этим числом было тоже число
// 			if lastNum > 0 {
// 				count = lastNum*10 - lastNum + count + 1
// 			}
// 			// если число повторений == 0, убираем последний символ
// 			if count == 0 {
// 				stack = append([]rune{}, stack[:len(stack)-1]...)
// 			}
// 			// сохраняем число для последующей обработки чисел из нескольких значений
// 			lastNum = count
// 			// повторяем число вписывая
// 			for count > 1 {
// 				stack = append(stack, stack[lenght-1])
// 				count--
// 			}
// 		} else {
// 			stack = append(stack, r)
// 		}
// 		i++
// 	}
// 	return string(stack), nil
// }

func main() {
	s := []string{
		"вф4кк5",
		"a4bc2d5e",
		"abcd",
		"45",
		"a1",
		"a10",
		"a20",
		"a22",
		"a100",
		"asd0",
		"",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
		`\5\\5`,
		`\\5\5`,
	}

	var r string
	var err error
	for _, v := range s {
		fmt.Println(v)
		r, err = UnPack(v)
		fmt.Println("unpack  :  ", r, len([]rune(r)), err)
	}
}
