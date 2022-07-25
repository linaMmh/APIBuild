package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"test3/common"

	"github.com/go-redis/redis/v8"
)

type KeepPi struct {
	ClientDB *redis.Client
}

func (r *KeepPi) setPi(indice string, response common.Response) error {
	client := r.ClientDB
	cacheField := "some-field-like-account-id"
	jsonsend, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}

	err = client.HSet(context.Background(), indice, cacheField, jsonsend).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("llave guardada")
	return nil
}

func (r *KeepPi) getPi(indice string) (common.Response, error) {
	client := r.ClientDB
	cacheField := "some-field-like-account-id"
	var response common.Response
	val, err := client.HGet(context.Background(), indice, cacheField).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return response, nil
	case err != nil:
		fmt.Println("Get failed", err)
		return response, err
	case val == "":
		fmt.Println("value is empty")
		return response, nil
	}

	fmt.Println(val)
	err = json.Unmarshal([]byte(val), &response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	return response, err
}

func NewKeepPiRepository(
	client *redis.Client,
) *KeepPi {
	return &KeepPi{
		ClientDB: client,
	}
}
