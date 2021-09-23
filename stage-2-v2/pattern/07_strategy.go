package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "errors"

var ErrTransportIsEmpty = errors.New("не выбран транспорт")

type Transport interface {
	getRoute() string
}

type Navigator interface {
	SetTransport(t Transport)
	GetRoute() (string, error)
}

type navigator struct {
	Transport Transport
}

func NewNavigator() Navigator {
	return &navigator{}
}

func (s *navigator) SetTransport(t Transport) {
	s.Transport = t
}

func (s *navigator) GetRoute() (string, error) {
	if s.Transport != nil {
		return s.Transport.getRoute(), nil
	}
	return "", ErrTransportIsEmpty
}

type bike struct{}

func (v *bike) getRoute() string {
	return "приятная поездка на велосипеде"
}

type auto struct{}

func (v *auto) getRoute() string {
	return "быстрая поездка на авто"
}

type bus struct{}

func (v *bus) getRoute() string {
	return "поездка в автобусе"
}
