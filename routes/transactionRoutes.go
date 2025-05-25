package routes

import (
	"github.com/MalshaPG/transaction-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(router *gin.Engine) {
	// With Gin, you can associate a handler with an HTTP method-and-path combination.
	// In this way, you can separately route requests sent to a single path
	// based on the method the client is using.
	router.POST("/transactions", controllers.SaveTransaction)
	router.GET("/transactions", controllers.GetTransactions)
	router.GET("/transactions/:type", controllers.GetTransactionByType)
	router.PUT("/transactions/:id", controllers.UpdateTransaction)
	router.DELETE("/transactions/:id", controllers.DeleteTransaction)
}
