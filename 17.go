package main

// Написать программу, которая в рантайме способна
// определить тип переменной — int, string, bool, channel
// из переменной типа interface{}.

import "fmt"

func getType(v interface{}) {
	switch v.(type) {

	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case chan struct{}:
		fmt.Println("chan struct{}")
	default:
		fmt.Println("undefined")
	}
}

func main() {
	getType("a")
	getType(1)
	getType('1')
	getType(true)
	getType(make(chan struct{}))
}
