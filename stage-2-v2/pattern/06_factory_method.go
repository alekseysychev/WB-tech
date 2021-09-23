package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "errors"

var ErrWrongType = errors.New("невверный тип")

type Transport interface {
	Delivery() string
}

type Factory struct{}

type car struct{}

func (c *car) Delivery() string {
	return "доставка автомобилем"
}

type ship struct{}

func (c *ship) Delivery() string {
	return "доставка судном"
}

func getTransport(transportType string) (Transport, error) {
	switch transportType {
	case "car":
		return &car{}, nil
	case "ship":
		return &ship{}, nil
	}
	return nil, ErrWrongType
}
