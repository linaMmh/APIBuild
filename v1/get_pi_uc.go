package v1

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"test3/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const IndexRedis string = "pi-decimals-%b"

type KeepPiInterface interface {
	setPi(index string, response common.Response) error
	getPi(index string) (common.Response, error)
	deletePi(index string) error
}

type ValidatePiRandom struct {
	InputNumber int `form:"input_number" binding:"required,numeric,excludesall=-."`
}

type ValidatePi struct {
	RandomNumber int `form:"random_number" binding:"required,numeric,excludesall=-."`
}

type GetPi struct {
	keepPiInterface    KeepPiInterface
	maxRandomPrecision int
	redisEnabled       bool
}

func (uc *GetPi) GetPiRandom(c *gin.Context) {
	// First: validate correct data
	var random ValidatePiRandom
	if err := c.ShouldBindQuery(&random); err != nil {
		er := common.ErrorResponse{
			UserMessage:     "The data sent is not valid",
			InternalMessage: "CONFLICT_BAD_PARAMS",
			MoreInfo:        err.Error(),
		}

		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}

	// Assign random number received
	inputNumber := random.InputNumber

	// calculate random number between (inputNumber/2 and inputNumber+1).
	randomCalculate, err := calculateRandom(inputNumber, uc.maxRandomPrecision)
	if err != nil {
		er := common.ErrorResponse{
			UserMessage:     "Random parameter would cause overflow",
			RandomGenerate:  randomCalculate,
			InternalMessage: "CONFLICT_RANDOM_NOT_VALID",
			MoreInfo:        err.Error(),
		}

		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}

	// If Redis is disabled: calculate pi and return the value.
	if !uc.redisEnabled {
		pi, _ := calculatePI(float64(randomCalculate))
		response := common.Response{
			Param:  inputNumber,
			Random: randomCalculate,
			PiCalc: fmt.Sprint(pi),
		}

		c.IndentedJSON(http.StatusOK, response)
		return
	}

	response, err := saveInRedis(randomCalculate, uc.keepPiInterface)
	// if exist some error return it
	if err != nil {
		er := common.ErrorResponse{
			UserMessage:     "A problem occurred when querying Redis",
			RandomGenerate:  randomCalculate,
			InternalMessage: "CONFLICT_WITH_REDIS",
			MoreInfo:        err.Error(),
		}

		c.IndentedJSON(http.StatusConflict, er)
		return
	}

	// assign missing parameter
	response.Param = inputNumber
	response.Random = randomCalculate

	c.IndentedJSON(http.StatusOK, response)
}

func (uc *GetPi) GetPi(c *gin.Context) {
	// First: validate correct data
	var random ValidatePi
	if err := c.ShouldBindQuery(&random); err != nil {
		er := common.ErrorResponse{
			UserMessage:     "The data sent is not valid",
			InternalMessage: "CONFLICT_BAD_PARAMS",
			MoreInfo:        err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}

	// Assign random number received
	inputNumber := random.RandomNumber

	// If Redis is disabled: calculate pi and return the value.
	if !uc.redisEnabled {
		pi, _ := calculatePI(float64(inputNumber))
		response := common.Response{
			Random: inputNumber,
			PiCalc: fmt.Sprint(pi),
		}

		c.IndentedJSON(http.StatusOK, response)
		return
	}

	response, err := saveInRedis(inputNumber, uc.keepPiInterface)
	// if exist some error return it
	if err != nil {
		er := common.ErrorResponse{
			UserMessage:     "A problem occurred when querying Redis",
			InternalMessage: "CONFLICT_WITH_REDIS",
			MoreInfo:        err.Error(),
		}

		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}
	// assign missing parameter
	response.Random = inputNumber

	c.IndentedJSON(http.StatusOK, response)
}

func (uc *GetPi) DeletePi(c *gin.Context) {
	// First: validate correct data
	var random ValidatePi
	if err := c.ShouldBindQuery(&random); err != nil {
		er := common.ErrorResponse{
			UserMessage:     "The data sent is not valid",
			InternalMessage: "CONFLICT_BAD_PARAMS",
			MoreInfo:        err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}

	if !uc.redisEnabled {
		er := common.ErrorResponse{
			UserMessage:     "Redis Server Disabled",
			InternalMessage: "CONFLICT_REDIS_DISABLED",
			MoreInfo:        "Redis Server Disabled",
		}
		c.IndentedJSON(http.StatusConflict, er)
		return
	}
	// search the index
	index := fmt.Sprintf(IndexRedis, random.RandomNumber)
	_, err := uc.keepPiInterface.getPi(index)
	if err != nil {
		er := common.ErrorResponse{
			UserMessage:     "A problem occurred when querying Redis",
			InternalMessage: "CONFLICT_WITH_REDIS",
			MoreInfo:        err.Error(),
		}

		c.IndentedJSON(http.StatusConflict, er)
		return
	}

	// Delete the index
	err = uc.keepPiInterface.deletePi(index)
	if err != nil {
		er := common.ErrorResponse{
			UserMessage:     "A problem occurred when querying Redis",
			InternalMessage: "CONFLICT_WITH_REDIS",
			MoreInfo:        err.Error(),
		}

		c.IndentedJSON(http.StatusConflict, er)
		return
	}
}

func calculateRandom(number, maxRandom int) (int, error) {
	min := number / 2
	max := number + 1
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min) + min

	// validate if random generated is valid
	if maxRandom != 0 && random > maxRandom {
		return 0, errors.New("random is not valid")
	}

	return random, nil
}

func calculatePI(decimals float64) (*big.Float, uint) {
	n := int64(2 + int(float64(decimals)/14.181647462))
	prec := uint(int(math.Ceil(math.Log2(10)*decimals)) + int(math.Ceil(math.Log10(decimals))) + 2)

	c := new(big.Float).Mul(
		big.NewFloat(float64(426880)),
		new(big.Float).SetPrec(prec).Sqrt(big.NewFloat(float64(10005))),
	)

	k := big.NewInt(int64(6))
	k12 := big.NewInt(int64(12))
	l := big.NewFloat(float64(13591409))
	lc := big.NewFloat(float64(545140134))
	x := big.NewFloat(float64(1))
	xc := big.NewFloat(float64(-262537412640768000))
	m := big.NewFloat(float64(1))
	sum := big.NewFloat(float64(13591409))

	pi := big.NewFloat(0)

	x.SetPrec(prec)
	m.SetPrec(prec)
	sum.SetPrec(prec)
	pi.SetPrec(prec)

	bigI := big.NewInt(0)
	bigOne := big.NewInt(1)

	for ; n > 0; n-- {
		// L calculation
		l.Add(l, lc)

		// X calculation
		x.Mul(x, xc)

		// M calculation
		kpower3 := big.NewInt(0)
		kpower3.Exp(k, big.NewInt(3), nil)
		ktimes16 := new(big.Int).Mul(k, big.NewInt(16))
		mtop := big.NewFloat(0).SetPrec(prec)
		mtop.SetInt(new(big.Int).Sub(kpower3, ktimes16))
		mbot := big.NewFloat(0).SetPrec(prec)
		mbot.SetInt(new(big.Int).Exp(new(big.Int).Add(bigI, bigOne), big.NewInt(3), nil))
		mtmp := big.NewFloat(0).SetPrec(prec)
		mtmp.Quo(mtop, mbot)
		m.Mul(m, mtmp)

		// Sum calculation
		t := big.NewFloat(0).SetPrec(prec)
		t.Mul(m, l)
		t.Quo(t, x)
		sum.Add(sum, t)

		// Pi calculation
		pi.Quo(c, sum)
		k.Add(k, k12)
		bigI.Add(bigI, bigOne)
	}
	return pi, prec
}

func saveInRedis(numberDecimals int, uc KeepPiInterface) (common.Response, error) {
	response := common.Response{}

	// Create index for Redis
	index := fmt.Sprintf(IndexRedis, numberDecimals)

	// search index
	resp, err := uc.getPi(index)
	if err != nil && err != redis.Nil {
		return response, err
	}

	// assign response
	response = resp
	// if the response is empty: calculate Pi
	if (common.Response{} == response) {
		pi, _ := calculatePI(float64(numberDecimals))
		response = common.Response{
			PiCalc: fmt.Sprint(pi),
		}
	}

	err = uc.setPi(index, response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func NewGetPi(
	keepPiInterface KeepPiInterface,
	maxRandomPrecision int,
	redisEnabled bool,
) *GetPi {
	return &GetPi{
		keepPiInterface:    keepPiInterface,
		maxRandomPrecision: maxRandomPrecision,
		redisEnabled:       redisEnabled,
	}
}
