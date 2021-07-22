package main

// Реализовать композицию структуры Action от родительской структуры Human.

import "fmt"

type Human struct {
	age int
}

type Action struct {
	Human
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func main() {
	a := &Action{}
	a.SetAge(123)
	fmt.Printf("%+v\n", a)
}
