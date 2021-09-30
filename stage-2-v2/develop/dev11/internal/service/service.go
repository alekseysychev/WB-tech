package service

import (
	"dev11/internal/dto"
	"dev11/internal/repository"
	"time"
)

type Service interface {
	EventsList() []dto.Event
	EventsGetForDay(int, time.Time) []dto.Event
	EventsGetForWeek(int, time.Time) []dto.Event
	EventsGetForMonth(int, time.Time) []dto.Event
	EventCreate(dto.Event)
	EventUpdate(dto.Event)
	EventDelete(dto.Event)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	repo.EventCreate(dto.Event{
		Text:   "user 1 event 1",
		Date:   time.Now().AddDate(0, 0, -2),
		UserId: 1,
	})
	repo.EventCreate(dto.Event{
		Text:   "user 1 event 2",
		Date:   time.Now().AddDate(0, -1, 10),
		UserId: 1,
	})
	repo.EventCreate(dto.Event{
		Text:   "user 1 event 3",
		Date:   time.Now().AddDate(0, 0, -2),
		UserId: 1,
	})
	repo.EventCreate(dto.Event{
		Text:   "user 2 event 1",
		Date:   time.Now().AddDate(0, -1, -2),
		UserId: 2,
	})
	repo.EventCreate(dto.Event{
		Text:   "user 2 event 2",
		Date:   time.Now().AddDate(0, 1, -2),
		UserId: 2,
	})
	repo.EventCreate(dto.Event{
		Text:   "user 3 event 1",
		Date:   time.Now().AddDate(0, 1, 10),
		UserId: 3,
	})
	return &service{
		repo: repo,
	}
}

func (s *service) EventsList() []dto.Event {
	return s.repo.EventsList()
}

func (s *service) EventsGetForDay(userId int, date time.Time) []dto.Event {
	// year, month, day := date.Date()
	// dateStart := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	// dateStop := dateStart.AddDate(0, 0, 1)
	dateStart := date.AddDate(0, 0, -1)
	dateStop := date.AddDate(0, 0, 1)
	return s.repo.EventGetForDate(userId, dateStart, dateStop)
}

func (s *service) EventsGetForWeek(userId int, date time.Time) []dto.Event {
	// year, month, day := date.Date()
	// dateStart := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	// dateStop := dateStart.AddDate(0, 0, 1)
	dateStart := date.AddDate(0, 0, -7)
	dateStop := date.AddDate(0, 0, 7)
	return s.repo.EventGetForDate(userId, dateStart, dateStop)
}

func (s *service) EventsGetForMonth(userId int, date time.Time) []dto.Event {
	// year, month, day := date.Date()
	// dateStart := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	// dateStop := dateStart.AddDate(0, 0, 1)
	dateStart := date.AddDate(0, -1, 0)
	dateStop := date.AddDate(0, 1, 0)
	return s.repo.EventGetForDate(userId, dateStart, dateStop)
}

func (s *service) EventCreate(event dto.Event) {
	s.repo.EventCreate(event)
}

func (s *service) EventUpdate(event dto.Event) {
	s.repo.EventUpdate(event)
}

func (s *service) EventDelete(event dto.Event) {
	s.repo.EventDelete(event)
}
