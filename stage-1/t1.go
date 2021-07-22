package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	ErrNotFound    error = errors.New("stack is empty")
	ErrWrongInput  error = errors.New("wrong input")
	ErrWrongInput1 error = errors.New("wrong input")
)

type stack struct {
	s []string
	sync.RWMutex
}

// добавление элемента
func (s *stack) push(i string) {
	s.Lock()
	defer s.Unlock()
	s.s = append(s.s, i)
}

// извлечение первого элемента
func (s *stack) pop() (string, error) {
	s.Lock()
	defer s.Unlock()
	n := len(s.s) - 1
	if n < 0 {
		return "", ErrNotFound
	}
	el := s.s[n]
	s.s = s.s[:n] // текущему значению присваиваем всё, кроме изъятого
	return el, nil
}

// чтение первого элемента
func (s *stack) peek() (string, error) {
	s.RLock()
	defer s.RUnlock()
	n := len(s.s) - 1
	if n < 0 {
		return "", ErrNotFound
	}
	el := s.s[n]
	return el, nil
}

func (s *stack) String() string {
	return strings.Join(s.s, "")
}

func zip(s string) (string, error) {
	var st stack

	for _, r := range []rune(s) {
		var counter int = 1
		var err error
		// достаём последний элемент
		lastSign, err := st.peek()
		if err != nil && err != ErrNotFound {
			return "", errors.New("error #1")
		}
		// если последний элемент равен текущему
		// достаём ещё 1, счётчик, увеличиваем на 1, последний кладём обратно
		if lastSign == string(r) {
			st.pop() // достаём lastSign
			counterString, err := st.pop()
			if err != nil && err != ErrNotFound {
				return "", errors.New("error #2")
			}
			counter, err = strconv.Atoi(string(counterString))
			if err != nil {
				return "", errors.New("error #3")
			}
			counter++
		}

		st.push(strconv.Itoa(counter))
		st.push(string(r))
	}
	return st.String(), nil
}

func IsPalindrome(i interface{}) (bool, error) {
	var s string

	switch i.(type) {
	case int:
		s = strconv.Itoa(i.(int))
	case string:
		s = i.(string)
	default:
		return false, ErrWrongInput
	}

	var back strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		back.WriteRune(runes[i])
	}

	return s == back.String(), nil
}

func main() {
	fmt.Println("zip")
	str := []string{
		"АААBBCCCDDEEE",
		"sdasdddddffFFF!!!!",
		"ssАААBBCCCDDEEE",
		"ssААА3333BBCCCDDEEE",
	}
	for _, s := range str {
		r, err := zip(s)
		fmt.Printf("in : %s\n", s)
		fmt.Printf("out: %s \n", r)
		fmt.Printf("err: %e\n", err)
	}
	fmt.Println("IsPalindrome string")
	str = []string{
		"0123456789",
		"01234567899876543210",
		"ABsddsBB",
		"ABsddsBA",
	}
	for _, s := range str {
		p, err := IsPalindrome(s)
		fmt.Printf("in:  %s \n", s)
		fmt.Printf("out: %v\n", p)
		fmt.Printf("err: %e\n", err)
	}
	fmt.Println("IsPalindrome int")
	ints := []int{123456789, 123456789987654321}
	for _, i := range ints {
		p, err := IsPalindrome(i)
		fmt.Printf("in:  %d \n", i)
		fmt.Printf("out: %v\n", p)
		fmt.Printf("err: %e\n", err)
	}
	fmt.Println(errors.Is(ErrWrongInput, ErrWrongInput1))
}
