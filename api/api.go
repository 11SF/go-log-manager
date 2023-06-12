package api

import (
	"go_log_manager/handlers"
	"go_log_manager/utils"

	"github.com/11SF/go-common/postgres"
	"github.com/labstack/echo"
)

// e *echo.Echo for Main Group
func TransactionGroup(e *echo.Group) {
	pgConfig := postgres.Config{
		Host:     utils.GetEnv("DB_HOST"),
		Username: utils.GetEnv("DB_USERNAME"),
		Password: utils.GetEnv("DB_PASSWORD"),
		Port:     utils.GetEnvInt("DB_PORT"),
		DBName:   utils.GetEnv("DB_NAME"),
		SSLMode:  "disable",
	}

	// Route / to handler function
	trs := handlers.TransactionHandler{}
	trs.Initialize(pgConfig)
	e.GET("/transaction", trs.GetAllTransaction)
	e.GET("/transaction/:familyId", trs.GetAllTransactionByFamilyId)
	e.POST("/transaction", trs.SaveTransaction)
	e.DELETE("/transaction", trs.DeleteTransaction)
	e.PUT("/transaction/:familyId", trs.UpdateTransaction)
}

// func AdminGroup(g *echo.Group) {
// 	g.GET("/main", handlers.MainAdmin)
// }
