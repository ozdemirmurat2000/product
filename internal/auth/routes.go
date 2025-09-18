package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(router *gin.RouterGroup, ctx AuthController) {

	announcement := router.Group("/v1")
	{
		announcement.POST("/auth/login", ctx.Login)
	}

}
