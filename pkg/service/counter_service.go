package service

import "tidy/pkg/repository"

type CountServ struct {
	storage repository.CounterStorage
}

func NewCounterService(storage repository.CounterStorage) *CountServ {
	return &CountServ{
		storage: storage,
	}
}
func (cs *CountServ) AddCounter(first string) error {
	cs.storage.AddCounter(first)

	return nil
}

func (cs *CountServ) SubCounter(num string) error {
	cs.storage.SubCounter(num)
	return nil
}

func (cs *CountServ) GetCounter() (string, error) {
	res, err := cs.storage.GetCounter()
	if err != nil {
		return "", err
	}
	return res, nil
}
