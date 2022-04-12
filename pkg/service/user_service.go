package service

import (
	"fmt"
	"strconv"
	"tidy/pkg/model"
	"tidy/pkg/repository"
)

type UserServ struct {
	storage repository.UserStorage
}

func NewUserService(storage repository.UserStorage) *UserServ {
	return &UserServ{
		storage: storage,
	}
}

func (s *UserServ) Create(m *model.User) (int, error) {
	id, err := s.storage.Create(m)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (s *UserServ) ReadOne(id string) (model.User, error) {
	id1, err := strconv.Atoi(id)
	if err != nil {
		return model.User{}, err
	}
	m, err := s.storage.ReadOne(id1)
	if err != nil {

		return m, err
	}

	if err != nil {
		fmt.Println(err)
		return m, err
	}

	return m, nil
}
func (s *UserServ) Update(m *model.UpdateU, id string) error {
	id1, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = s.storage.Update(m, id1)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *UserServ) Delete(m string) error {
	id, err := strconv.Atoi(m)
	if err != nil {
		return err
	}
	err = s.storage.Delete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
