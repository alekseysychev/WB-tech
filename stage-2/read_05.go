// Что выведет программа? Объяснить вывод программы.

package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// добавлено
func test2() error {
	if false {
		return &customError{
			msg: "error msg",
		}
	}
	return nil
}

// на мой взгляд пример аналогичен 02
// test возвращает структуру, которая приводится к интерфейсу error, так как мы заранее объъявили тип переменной
// что бы поправить и дать ожидаемое поведение, надо убрать 21 строку, на 22 сделать :=
// но лучше возвращать как test2
