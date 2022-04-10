package repository

import (
	"tidy/dbase"
	"tidy/pkg/model"
)

type EmailCheckStorage interface {
}
type SubstringStorage interface {
}
type UserStorage interface {
	Create(m *model.User) error
	ReadOne(id string) (model.User, error)
	Update(m *model.User, id int) error
	Delete(id int) error
}
type CounterStorage interface {
	AddCounter(first string) error
	SubCounter(first string) error
	GetCounter() (string, error)
}

type Repository struct {
	EmailCheckStorage
	SubstringStorage
	CounterStorage
	UserStorage
}

func NewRepository(db dbase.Database) *Repository {
	return &Repository{
		CounterStorage: NewCounterStorage(db.DbRed),
		UserStorage:    NewUserStorage(db.DbSql),
	}
}
