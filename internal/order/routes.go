package order

import (
	"productApp/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.RouterGroup, ctx IOrderController) {

	order := router.Group("/v1/order")
	summary := order.Group("/summary")
	{
		order.GET("/", middleware.AuthMiddleware(), ctx.GetOrderBySiparisID)
		summary.GET("/", middleware.AuthMiddleware(), ctx.GetOrderSummaryList)
		summary.GET("/uretim", middleware.AuthMiddleware(), ctx.GetOrderUretimBilgileriBySiparisID)
		order.POST("/uretim", middleware.AuthMiddleware(), ctx.AddNewUretim)
		order.DELETE("/uretim", middleware.AdminMiddleware(), ctx.DeleteUretim)
	}

	order.GET("/customerOrders", middleware.AuthMiddleware(), ctx.GetCustomerOrdersByIslemAdi)
	order.POST("/modelResim", ctx.AddNewModelResim)

}
