package api

import (
	"11sf/go_log_manager/handlers"

	"github.com/labstack/echo"
)

// e *echo.Echo for Main Group
func TransactionGroup(e *echo.Group) {
	// Route / to handler function
	trs := handlers.TransactionHandler{}
	trs.Initialize()
	e.GET("/transaction", trs.GetAllTransaction)
	e.GET("/transaction/:familyId", trs.GetAllTransactionByFamilyId)
	e.POST("/transaction", trs.SaveTransaction)
	e.DELETE("/transaction", trs.DeleteTransaction)
	e.PUT("/transaction/:familyId", trs.UpdateTransaction)
}

// func AdminGroup(g *echo.Group) {
// 	g.GET("/main", handlers.MainAdmin)
// }
