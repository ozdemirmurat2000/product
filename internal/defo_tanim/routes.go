package defotanim

import (
	"productApp/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterDefoTanimRoutes(api *gin.RouterGroup, controller IDefoTanimController) {

	defo := api.Group("/v1/defo")
	{
		defo.GET("/", middleware.AuthMiddleware(), controller.GetDefoTanimList)
	}

}
