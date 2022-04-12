package service

import (
	"tidy/pkg/model"
	"tidy/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserService interface {
	Create(m *model.User) (int, error)
	ReadOne(id string) (model.User, error)
	Update(m *model.UpdateU, id string) error
	Delete(m string) error
}

type SubstringService interface {
	MaxLength(s *string) (string, error)
}
type EmailCheckService interface {
	EmailCheck(s string) (string, error)
	IinCheck(s string) (string, error)
}
type CounterService interface {
	AddCounter(first string) error
	SubCounter(num string) error
	GetCounter() (string, error)
}
type HashCalcService interface {
	GetID() (string, error)
	GetHash(s string) (string, error)
}
type Service struct {
	UserService
	SubstringService
	CounterService
	EmailCheckService
	HashCalcService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService:       NewUserService(repos.UserStorage),
		SubstringService:  NewSubService(),
		CounterService:    NewCounterService(repos.CounterStorage),
		EmailCheckService: NewEmailService(),
		HashCalcService:   NewHashService(),
	}
}
