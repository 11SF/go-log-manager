package main

import (
	"go_log_manager/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
)

func init() {
	godotenv.Load()
}

func main() {
	e := router.New()

	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://family-pay-14de8.web.app", "http://localhost:3000", "https://familypay.11sf.site"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))
}
