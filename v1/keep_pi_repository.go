package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"test3/common"

	"github.com/go-redis/redis/v8"
)

const CacheField string = "pi-decimals-number"

type KeepPi struct {
	ClientDB *redis.Client
}

func (r *KeepPi) setPi(index string, response common.Response) error {
	jsonSend, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = r.ClientDB.HSet(context.Background(), index, CacheField, jsonSend).Err()
	if err != nil {
		return err
	}

	fmt.Println("key saved successfully")
	return nil
}

func (r *KeepPi) getPi(index string) (common.Response, error) {
	client := r.ClientDB
	val, err := client.HGet(context.Background(), index, CacheField).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return common.Response{}, err
	case err != nil:
		return common.Response{}, err
	case val == "":
		fmt.Println("value is empty")
		return common.Response{}, nil
	}

	// unmarshal response
	var response common.Response
	err = json.Unmarshal([]byte(val), &response)
	if err != nil {
		return common.Response{}, err
	}

	fmt.Println("key obtained successfully")
	return response, nil
}

func (r *KeepPi) deletePi(index string) error {
	client := r.ClientDB
	_, err := client.Del(context.Background(), index).Result()
	if err != nil {
		return err
	}

	fmt.Println("key deleted successfully")
	return nil
}

func NewKeepPiRepository(
	client *redis.Client,
) *KeepPi {
	return &KeepPi{
		ClientDB: client,
	}
}
