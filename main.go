package main

import (
	"github.com/MalshaPG/transaction-tracker/database"
	"github.com/MalshaPG/transaction-tracker/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	database.Connect()

	//Close the database connection when the application exits
	defer database.DB.Close()

	// Initialize Gin router
	router := gin.Default()

	routes.RegisterTransactionRoutes(router)

	// Run the server on localhost:8080
	router.Run("localhost:8080")
}
