package main

import (
	conn "apiGo/connection"
	"apiGo/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := conn.GetEnv("PORT", ":3000")
	App := echo.New()
	App.Use(middleware.Logger())
	App.Use(middleware.Recover())

	router.Router(App)
	App.Logger.Fatal(App.Start(port))
}
