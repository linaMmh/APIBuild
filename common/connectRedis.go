package common

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// ParamsConnection parameters to make the connection to redis
type ParamsConnection struct {
	Password string
	DBNumber int
	Host     string
}

// ConnectDB interface form redis connection
type ConnectDB interface {
	GetClient() *redis.Client
}

// ConnectRedis struct form redis connection
type ConnectRedis struct {
	Client *redis.Client
}

// GetClient get redis client
func (cm *ConnectRedis) GetClient() *redis.Client {
	return cm.Client
}

// NewCreateConnection create redis connection
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
