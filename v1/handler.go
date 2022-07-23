package v1

import "github.com/gin-gonic/gin"

type GetPiInterface interface {
	GetPiRandom(c *gin.Context)
	GetPi(c *gin.Context)
}

type Api struct {
	getPiInterface GetPiInterface
}

// Handler
func (a *Api) Handler() error {
	router := gin.Default()
	router.GET("/getPiRandom", a.getPiInterface.GetPiRandom)
	router.GET("/getPi", a.getPiInterface.GetPi)

	err := router.Run("localhost:8080")
	if err != nil {
		panic("un error inesperado")
	}

	return nil
}

func NewApi(
	getPiInterface GetPiInterface,
) *Api {
	return &Api{
		getPiInterface: getPiInterface,
	}
}
