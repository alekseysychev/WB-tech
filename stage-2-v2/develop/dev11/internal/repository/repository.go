package repository

import (
	"dev11/internal/dto"
	"sync"
	"time"
)

type Repository interface {
	EventCreate(dto.Event)
	EventUpdate(dto.Event)
	EventDelete(dto.Event)
	EventsList() []dto.Event
	EventGetForDate(int, time.Time, time.Time) []dto.Event
}

type repository struct {
	events map[int]dto.Event
	lastId int
	sync.RWMutex
}

func New() Repository {
	return &repository{
		events: make(map[int]dto.Event),
	}
}

func (r *repository) EventCreate(event dto.Event) {
	r.Lock()
	event.Id = r.lastId
	r.events[event.Id] = event
	r.lastId++
	r.Unlock()
}

func (r *repository) EventUpdate(event dto.Event) {
	r.Lock()
	r.events[event.Id] = event
	r.Unlock()
}

func (r *repository) EventDelete(event dto.Event) {
	r.Lock()
	delete(r.events, event.Id)
	r.Unlock()
}

func (r *repository) EventsList() []dto.Event {
	result := make([]dto.Event, 0)
	for _, v := range r.events {
		result = append(result, v)
	}
	return result
}

func (r *repository) EventGetForDate(userId int, dateStart, dateStop time.Time) []dto.Event {
	result := make([]dto.Event, 0)
	for _, v := range r.events {
		if v.UserId == userId && v.Date.After(dateStart) && v.Date.Before(dateStop) {
			result = append(result, v)
		}
	}
	return result
}
