package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

type state interface {
	setClock(hour int)
	getTimesOfDay() string
}

type timeOfDay struct {
	day   state
	night state

	currentState state
}

func newTimeOfDay() *timeOfDay {
	v := &timeOfDay{}
	v.day = &dayState{timeOfDay: v}
	v.night = &nightState{timeOfDay: v}
	v.currentState = &dayState{timeOfDay: v}
	return v
}

func (d *timeOfDay) changeState(s state) {
	d.currentState = s
}

func (d *timeOfDay) setClock(hour int) {
	d.currentState.setClock(hour)
}

func (d *timeOfDay) getTimesOfDay() string {
	return d.currentState.getTimesOfDay()
}

type dayState struct {
	timeOfDay *timeOfDay
}

func (d *dayState) setClock(hour int) {
	if hour < 9 || 17 <= hour {
		fmt.Println("here", d.timeOfDay)
		d.timeOfDay.changeState(d.timeOfDay.night)
	}
}

func (d *dayState) getTimesOfDay() string {
	return "день"
}

type nightState struct {
	timeOfDay *timeOfDay
}

func (d *nightState) setClock(hour int) {
	if hour >= 9 || 17 > hour {
		d.timeOfDay.changeState(d.timeOfDay.day)
	}
}

func (d *nightState) getTimesOfDay() string {
	return "ночь"
}
