package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{
		"пятак", "тяпка", "пятка", "листок", "слиток", "столик", "оон", "но", "он", "оно", "ноо",
		"Волос", "слово", "каша", "маша", "шама", "кама", "столик", "слово_одиночка", "листок",
	}
	fmt.Println(find(words))
}

type wordsStruct struct {
	first string              // первое найденное слово
	words map[string]struct{} //
}

func find(words []string) map[string][]string {
	m := make(map[string]wordsStruct)
	for _, word := range words {
		// преобразуем к нижнему регистру
		wordLower := strings.ToLower(word)
		// получаем ключ
		key := makeKey(wordLower)
		// собираем по ключу список слов
		if v, ok := m[key]; ok { // если у нас уже есть такой ключ и текущее слово это не ключ
			// добавляем слово, если это не наш ключ
			if wordLower != v.first {
				v.words[wordLower] = struct{}{}
			}
		} else {
			m[key] = wordsStruct{
				first: wordLower,
				words: make(map[string]struct{}),
			}
		}
	}
	result := make(map[string][]string)
	for _, me := range m {
		// пропуск элементов у которых меньше 3 анаграмм
		if len(me.words) < 2 {
			continue
		}
		// собираем слайс уникальных слов
		for word, _ := range me.words {
			result[me.first] = append(result[me.first], word)
		}
		// сортируем слайс слов
		sort.Slice(result[me.first], func(i, j int) bool {
			return result[me.first][i] < result[me.first][j]
		})
	}
	return result
}

func makeKey(wordLower string) string {
	// разбиваем на буквы
	wordRune := []rune(wordLower)
	// сортируем
	sort.Slice(wordRune, func(i, j int) bool {
		return wordRune[i] < wordRune[j]
	})
	// возвращаем в строку
	key := string(wordRune)
	return key
}
