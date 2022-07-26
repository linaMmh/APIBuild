package v1

import "github.com/gin-gonic/gin"

// GetPiInterface interface from main use Case
type GetPiInterface interface {
	GetPiRandom(c *gin.Context)
	GetPi(c *gin.Context)
	DeletePi(c *gin.Context)
}

// Api structure for main controller
type Api struct {
	getPiInterface GetPiInterface
}

// Handler main function
func (a *Api) Handler() error {
	router := gin.Default()
	router.GET("/getPiRandom", a.getPiInterface.GetPiRandom)
	router.GET("/getPi", a.getPiInterface.GetPi)
	router.DELETE("/deletePi", a.getPiInterface.DeletePi)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

	return nil
}

// NewApi main constructor
func NewApi(
	getPiInterface GetPiInterface,
) *Api {
	return &Api{
		getPiInterface: getPiInterface,
	}
}
