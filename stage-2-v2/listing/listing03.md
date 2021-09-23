Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:

```text
error это интерфейс, для реализации должен быть метод Error() string
интерфейс внутри это 2 элемента тип T и значение V
проверить можно через рефлект, но лучше избегать вариантов, когда возвращается пустая ошибка
reflect.ValueOf(err).IsNil()
```
