package router

import (
	"11sf/go_log_manager/api"
	// "11sf/go_log_manager/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	// adminGroup := e.Group("/admin")

	//set all middlewares
	// middlewares.SetMainMiddleWares(e)
	// middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	// e.Use(middleware.CORS())
	fpGroup := e.Group("/fp")
	api.TransactionGroup(fpGroup)

	//set groupRoutes
	// api.AdminGroup(adminGroup)
	return e
}
