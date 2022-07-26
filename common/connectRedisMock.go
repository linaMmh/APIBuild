package common

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

// ConnectDBMock connect mock from test.
type ConnectDBMock interface {
	GetClient() *redis.Client
}

// ConnectRedisMock connect mock from test.
type ConnectRedisMock struct {
	Client *redis.Client
}

// GetClient mock from test.
func (cm *ConnectRedisMock) GetClient() *redis.Client {
	return cm.Client
}

// NewCreateConnectionMock from test.
func NewCreateConnectionMock() (ConnectDBMock, redismock.ClientMock) {
	db, mockRedis := redismock.NewClientMock()

	return &ConnectRedisMock{
		Client: db,
	}, mockRedis
}
