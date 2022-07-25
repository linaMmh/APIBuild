package common

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

type ConnectDBMock interface {
	GetClient() *redis.Client
}

type ConnectRedisMock struct {
	Client *redis.Client
}

func (cm *ConnectRedisMock) GetClient() *redis.Client {
	return cm.Client
}

func NewCreateConnectionMock() (ConnectDBMock, redismock.ClientMock) {
	db, mockRedis := redismock.NewClientMock()

	return &ConnectRedisMock{
		Client: db,
	}, mockRedis
}
