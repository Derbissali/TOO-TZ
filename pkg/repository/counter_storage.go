package repository

import (
	"strconv"

	"github.com/go-redis/redis"
)

type counterStorage struct {
	db redis.Cmdable
}

func NewCounterStorage(db redis.Cmdable) CounterStorage {
	return &counterStorage{
		db: db,
	}
}
func (cr *counterStorage) AddCounter(first string) error {

	firstNum, err := strconv.Atoi(first)

	val, err := cr.db.Get("counter").Result()
	if err != nil {
		return err
	}
	secondNum, err := strconv.Atoi(val)

	res := firstNum + secondNum
	result := strconv.Itoa(res)
	err = cr.db.Set("counter", result, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr *counterStorage) SubCounter(first string) error {

	firstNum, err := strconv.Atoi(first)

	val, err := cr.db.Get("counter").Result()
	if err != nil {
		return err
	}
	secondNum, err := strconv.Atoi(val)

	res := secondNum - firstNum
	result := strconv.Itoa(res)
	err = cr.db.Set("counter", result, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr *counterStorage) GetCounter() (string, error) {

	val := cr.db.Get("counter")

	return val.Result()
}
