package repository

import (
	"log"
	"os"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	client *redis.Client
)

var (
	key = "1"
	val = "val"
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	code := m.Run()
	os.Exit(code)
}

func TestSet(t *testing.T) {

	mock := redismock.NewNiceMock(client)
	mock.On("Set", "counter", key).Return(redis.NewStatusResult("", nil))

	r := NewCounterStorage(mock)

	err := r.AddCounter(key)
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Get", "counter").Return(redis.NewStringResult(val, nil))

	r := NewCounterStorage(mock)
	res, err := r.GetCounter()
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}
