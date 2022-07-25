package main

import (
	"fmt"
	"test3/common"
	v1 "test3/v1"
)

const MAX_RANDOM_PRECISION int = 300

const REDIS_ENABLED bool = true

func main() {
	pc := common.ParamsConnection{
		Password: "",
		DBNumber: 0,
		Host:     "localhost:6379",
	}
	connect, err := common.NewCreateConnection(pc)
	if err != nil {
		fmt.Println(err)
		return
	}

	redisRepo := v1.NewKeepPiRepository(connect.GetClient())
	getpi := v1.NewGetPi(redisRepo, MAX_RANDOM_PRECISION, REDIS_ENABLED)
	repo := v1.NewApi(getpi)
	err = repo.Handler()
	if err != nil {
		fmt.Println(err)
		return
	}
}
