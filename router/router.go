package router

import (
	"apiGo/controller"

	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	ndea := new(controller.Controller)
	e.GET("/Products-", ndea.GetProductsBy)
	e.GET("/Products", ndea.GetProducts)
	e.DELETE("/Product", ndea.DeleteProduct)
	e.GET("/Product", ndea.GetProduct)
	e.POST("/Product", ndea.PostProduct)
	e.POST("/Products", ndea.InserMany)
	e.GET("/DB", ndea.Connection)
}
