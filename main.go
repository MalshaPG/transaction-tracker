package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
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

	// With Gin, you can associate a handler with an HTTP method-and-path combination. 
	// In this way, you can separately route requests sent to a single path 
	// based on the method the client is using.
    router.POST("/transactions", saveTransaction)
    router.GET("/transactions", getTransactions)
    router.GET("/transactions/:type", getTransactionByType)
    // r.PUT("/transactions/:id", updateTransaction)

    //Use the Run function to attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}

// getAlbums responds with the list of all transactions as JSON.
//gin.Context is the most important part of Gin. 
//It carries request details, validates and serializes JSON, and more. 
func getTransactions(c *gin.Context) {
    //Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
    c.IndentedJSON(http.StatusOK, transactions)
}

// saveTransaction adds a transaction from JSON received in the request body.
func saveTransaction(c *gin.Context) {
    var newTransaction transaction

    // Call BindJSON to bind the received JSON to newTransaction.
    if err := c.BindJSON(&newTransaction); err != nil {
        return
    }

    // Add the new transaction to the slice.
    transactions = append(transactions, newTransaction)

	// Add a 201 status code to the response, along with JSON representing the transaction added.
    c.IndentedJSON(http.StatusCreated, newTransaction)
}

//getTransactionByType located the transaction whose type value matches the id
// paramenter sent by the client, then returns thst tarnsactions as a response
func getTransactionByType(c *gin.Context) {
    transactionType := c.Param("type")

    var results []transaction 

    for _, a := range transactions {
        if strings.ToLower(a.Type) == strings.ToLower(transactionType) {
            results = append(results, a)    
        }
    }

    if len(results) == 0 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, results)
}







