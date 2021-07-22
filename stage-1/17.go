package main

// Написать программу, которая в рантайме способна
// определить тип переменной — int, string, bool, channel
// из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) {
	switch v.(type) {

	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case int32:
		fmt.Println("int32")
	case bool:
		fmt.Println("bool")
	case chan struct{}:
		fmt.Println("chan struct{}")
	default:
		fmt.Println("undefined")
	}
}

func getType2(v interface{}) {
	switch reflect.TypeOf(v).String() {
	case "string":
		fmt.Println("string")
	case "int":
		fmt.Println("int")
	case "int32":
		fmt.Println("int32")
	case "bool":
		fmt.Println("bool")
	case "chan struct {}":
		fmt.Println("chan struct {}")
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

	fmt.Println()

	getType2("a")
	getType2(1)
	getType2('1')
	getType2(true)
	getType2(make(chan struct{}))
}
