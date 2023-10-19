package routes

import (
	"github.com/gin-gonic/gin"
	"secondTask/internal/controllers"
)

type Routes struct {
	Controller controllers.ApiController
}

func NewRoutes(controller controllers.ApiController) *Routes {
	return &Routes{Controller: controller}
}

func (r *Routes) SetupAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.GET("/getCryptoCurrencyPrice/:symbol", r.Controller.GetCryptoCurrencyPrice)
	}
}
