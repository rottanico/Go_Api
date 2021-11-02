package main

import (
	"apiGo/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	App := echo.New()
	App.Use(middleware.Logger())
	App.Use(middleware.Recover())
	router.Router(App)
	App.Logger.Fatal(App.Start(":3000"))
}
