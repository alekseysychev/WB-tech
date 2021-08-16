// Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func testDefer() { // добавлено
	fmt.Println("--- testDefer ---") // добавлено
	defer func() {                   // добавлено
		fmt.Println("3") // добавлено
	}() // добавлено
	defer func() { // добавлено
		fmt.Println("2") // добавлено
	}() // добавлено
	fmt.Println("1") // добавлено
} // добавлено

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
	testDefer() // добавлено
}

// defer помещает функцию в стек, который выполняется после возврата окружающей функции
// test в данном случае имеет именованный возврат переменной и после return переменную всё ещё можно изменить через defer
// anotherTest не имеет именованного возврата и
// так же важная особенность, что работает принцип LIFO
