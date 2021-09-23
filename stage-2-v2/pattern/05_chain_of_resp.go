package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import (
	"strconv"
)

// обобщенная проблема
type Trouble struct {
	number int
}

func (t *Trouble) getNumber() int {
	return t.number
}

// интерфейс поддержки
type support interface {
	resolve(trouble Trouble) bool
	Handle(support support, trouble Trouble) string
}

type defaultSupport struct {
	support
	name string
	next support
}

func (dS *defaultSupport) SetNext(next support) {
	dS.next = next
}

// обработка пролемы
func (dS *defaultSupport) Handle(support support, trouble Trouble) string {
	if support.resolve(trouble) {
		return "trouble:" + strconv.Itoa(trouble.getNumber()) + " is resolved by " + dS.name
	} else if dS.next != nil {
		return dS.next.Handle(dS.next, trouble)
	} else {
		return "trouble:" + strconv.Itoa(trouble.getNumber()) + " cannot be resolved"
	}
}

// структура без поддержки
type noSupport struct {
	*defaultSupport
}

// решает ли структура проблему
func (nS *noSupport) resolve(trouble Trouble) bool {
	return false
}

func NewNoSupport(name string) *noSupport {
	return &noSupport{&defaultSupport{name: name}}
}

// структура с ограниченной поддержкой
type limitSupport struct {
	*defaultSupport
	limit int
}

// решает ли структура проблему
func (lS *limitSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber() < lS.limit {
		return true
	} else {
		return false
	}
}

func NewLimitSupport(name string, limit int) *limitSupport {
	return &limitSupport{&defaultSupport{name: name}, limit}
}
