package service

import (
	"dev_11/internal/dto"
	"dev_11/internal/repository"
	"time"
)

type Service interface {
	Create(dto.Event) error
	Update(dto.Event) error
	Delete(dto.Event) error
	GetByTime(time.Timer) ([]dto.Event, error)
}

type service struct{}

func (s *service) Create(e dto.Event) error {
	return nil
}

func (s *service) Update(e dto.Event) error {
	return nil
}

func (s *service) Delete(e dto.Event) error {
	return nil
}

func (s *service) GetById(period time.Timer) ([]dto.Event, error) {
	return nil, nil
}

func (s *service) GetByTime(period time.Timer) ([]dto.Event, error) {
	return nil, nil
}

func New(repo repository.Repository) (Service, error) {
	return &service{}, nil
}
