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
	Create(m *model.User) (int, error)
	ReadOne(id int) (model.User, error)
	Update(m *model.UpdateU, id int) error
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
