package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type counterStorage struct {
	db *redis.Client
}

func NewCounterStorage(db *redis.Client) CounterStorage {
	return &counterStorage{
		db: db,
	}
}
func (cr *counterStorage) AddCounter(first string) error {
	ctx := context.Background()

	firstNum, err := strconv.Atoi(first)

	val, err := cr.db.Get(ctx, "counter").Result()
	if err != nil {
		return err
	}
	secondNum, err := strconv.Atoi(val)

	res := firstNum + secondNum
	result := strconv.Itoa(res)
	err = cr.db.Set(ctx, "counter", result, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr *counterStorage) SubCounter(first string) error {
	ctx := context.Background()

	firstNum, err := strconv.Atoi(first)

	val, err := cr.db.Get(ctx, "counter").Result()
	if err != nil {
		return err
	}
	secondNum, err := strconv.Atoi(val)

	res := secondNum - firstNum
	result := strconv.Itoa(res)
	err = cr.db.Set(ctx, "counter", result, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cr *counterStorage) GetCounter() (string, error) {
	ctx := context.Background()
	val, err := cr.db.Get(ctx, "counter").Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
