package main

import (
	"os"
	"pi-api/common"
	v1 "pi-api/v1"
	"strconv"
)

const (
	// MaxRandomPrecisionDefault const default for random generation
	MaxRandomPrecisionDefault string = "300"
	// RedisEnabledDefault redis is default in true
	RedisEnabledDefault string = "true"
	// RedisUrlDefault redis posrt default
	RedisUrlDefault string = "localhost:6379"
)

// main function in charge of inverting the dependencies of the repository and the use case
func main() {
	// set params from Redis connection
	redisURL := os.Getenv("REDIS_LOCAL_URL")
	if redisURL == "" {
		redisURL = RedisUrlDefault
	}

	pc := common.ParamsConnection{
		Password: "",
		DBNumber: 0,
		Host:     redisURL,
	}
	// Connect with redis
	connect, err := common.NewCreateConnection(pc)
	if err != nil {
		panic(err)
		return
	}

	// Instantiate repository
	redisRepository := v1.NewKeepPiRepository(connect.GetClient())

	// Instantiate useCase
	// set params from env vars
	randomPrecision := os.Getenv("MAX_RANDOM_PRECISION")
	if randomPrecision == "" {
		randomPrecision = MaxRandomPrecisionDefault
	}

	redisEnabled := os.Getenv("REDIS_ENABLED")
	if redisEnabled == "" {
		redisEnabled = RedisEnabledDefault
	}

	randomPrecisionInt, err := strconv.Atoi(randomPrecision)
	if err != nil {
		panic(err)
		return
	}

	redisEnabledBool, err := strconv.ParseBool(redisEnabled)
	if err != nil {
		panic(err)
		return
	}
	getPi := v1.NewGetPi(redisRepository, randomPrecisionInt, redisEnabledBool)

	// Instantiate handler
	repo := v1.NewApi(getPi)
	err = repo.Handler()
	if err != nil {
		panic(err)
		return
	}
}
