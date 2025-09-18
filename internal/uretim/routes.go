package uretim

import (
	"productApp/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUretimRoutes(router *gin.RouterGroup, service IUretimController) {

	uretim := router.Group("/v1/uretim")
	{
		uretim.GET("/", middleware.AuthMiddleware(), service.GetUretimList)
		uretim.DELETE("/", middleware.AdminMiddleware(), service.DeleteUretim)
		uretim.POST("/", middleware.AuthMiddleware(), service.AddUretim)
	}

}
