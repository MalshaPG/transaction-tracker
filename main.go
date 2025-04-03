package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type transaction struct {
    ID     int     `json:"id"`
    Type  string   `json:"type"`
    Name   string  `json:"name"`
    Amount float64 `json:"amount"`
}

// transactions slice to seed record trancation data.
// var is used for global (package-level) variables.
// := cannot be used at the package level.
var transactions = []transaction{
    {ID: 1, Type: "Credit", Name: "Salary", Amount: 100},
    {ID: 2, Type: "Debit", Name: "Buy groceries", Amount: 10.45},
    {ID: 3, Type: "Debit", Name: "Movie ticket for Moana", Amount: 6.32}}

func main() {
    // Initialize a Gin router using Default.
    router := gin.Default()
    // r.POST("/transactions", createTransaction)
    router.GET("/transactions", getTransactions)
    // r.PUT("/transactions/:id", updateTransaction)

    //Use the Run function to attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
//gin.Context is the most important part of Gin. 
//It carries request details, validates and serializes JSON, and more. 
func getTransactions(c *gin.Context) {
    //Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
    c.IndentedJSON(http.StatusOK, transactions)
}





