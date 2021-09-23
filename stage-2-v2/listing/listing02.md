Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
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


func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

Ответ:

```text
2 - так как переменная x увеличится через defer, а вывод именованный
1 - так как хоть х и увеличится, вывод не именованный и х останется внутри функции

defer добавляет в стек(LIFO) вызовов после выхода из функции
```
