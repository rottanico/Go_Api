package router

import (
	"apiGo/controller"

	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	router := new(controller.Controller)
	e.GET("/Product", router.GetProduct)
	e.GET("/Products", router.GetProducts)
	e.GET("/Products-", router.GetProductsBy)
	e.POST("/Product", router.PostProduct)
	e.POST("/Products", router.InserMany)
	e.GET("/DB", router.Connection)
	e.PUT("/Product", router.UpdateStock)
	e.PUT("/Product-", router.UpdateProduct)
	e.DELETE("/Product", router.DeleteProduct)
}
