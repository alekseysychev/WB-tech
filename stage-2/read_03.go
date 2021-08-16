// Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

package main

import (
	"fmt"
	"io/fs"
	"reflect"
)

func Foo() error {
	var err *fs.PathError = nil

	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Printf("%#v\n", err)                  // добавлено
	fmt.Println(reflect.ValueOf(err).IsNil()) // добавлено
}

// error это интерфейс, для реализации должен быть метод Error() string
// интерфейс внутри это 2 элемента тип T и значение V
// проверить можно через рефлект, но лучше избегать вариантов, когда возвращается пустая ошибка
