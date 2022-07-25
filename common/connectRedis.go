package common

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type ParamsConnection struct {
	Password string
	DBNumber int
	Host     string
}

type ConnectDB interface {
	GetClient() *redis.Client
}

type ConnectRedis struct {
	Client *redis.Client
}

func (cm *ConnectRedis) GetClient() *redis.Client {
	return cm.Client
}

func NewCreateConnection(pc ParamsConnection) (ConnectDB, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     pc.Host,
		Password: pc.Password,
		DB:       pc.DBNumber,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &ConnectRedis{
		Client: client,
	}, nil
}
