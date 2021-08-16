package main

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
