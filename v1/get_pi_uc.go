package v1

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	UserMessage     string `json:"userMessage"`
	RandomGenerate  int    `json:"randomGenerate,omitempty"`
	InternalMessage string `json:"internalMessage"`
	MoreInfo        string `json:"moreInfo"`
}

type Response struct {
	Param  int    `json:"param,omitempty"`
	Random int    `json:"random"`
	PiCalc string `json:"PiCalc"`
}

type ValidatePiRandom struct {
	InputNumber int `form:"input_number" binding:"required,numeric,excludesall=-."`
}

type ValidatePi struct {
	RandomNumber int `form:"random_number" binding:"required,numeric,excludesall=-."`
}

type GetPi struct{}

func (uc *GetPi) GetPiRandom(c *gin.Context) {
	var random ValidatePiRandom
	if err := c.ShouldBindQuery(&random); err != nil {
		er := ErrorResponse{}
		er.UserMessage = "MESSAGE"
		er.InternalMessage = "BAD_PARAMS"
		er.MoreInfo = err.Error()
		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}
	inputNumber := random.InputNumber
	randomCalculate := calculateRandom(inputNumber)
	pi, _ := calculatePI(float64(randomCalculate))

	response := Response{
		Param:  inputNumber,
		Random: randomCalculate,
		PiCalc: fmt.Sprint(pi),
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (uc *GetPi) GetPi(c *gin.Context) {
	var random ValidatePi

	if err := c.ShouldBindQuery(&random); err != nil {
		er := ErrorResponse{}
		er.UserMessage = "MESSAGE"
		er.InternalMessage = "BAD_PARAMS"
		er.MoreInfo = err.Error()
		c.IndentedJSON(http.StatusBadRequest, er)
		return
	}

	inputNumber := random.RandomNumber
	pi, _ := calculatePI(float64(inputNumber))

	response := Response{
		Random: inputNumber,
		PiCalc: fmt.Sprint(pi),
	}

	c.IndentedJSON(http.StatusOK, response)
}

func calculateRandom(number int) int {
	min := number / 2
	max := number + 1
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
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

// NewGetOrdersRepository  initialize repository.
func NewGetPi() *GetPi {
	return &GetPi{}
}
