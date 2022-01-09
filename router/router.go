package router

import (
	"assignment-2/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	orderRouter := r.Group("/orders")
	{
		orderRouter.POST("/", controller.CreateOrder)
		orderRouter.GET("/", controller.GetOrders)
		orderRouter.PUT("/:orderId", controller.UpdateOrder)
		orderRouter.DELETE("/:orderId", controller.DeleteOrder)
	}
	return r
}
