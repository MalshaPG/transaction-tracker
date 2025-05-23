package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

type transaction struct {
	ID          int     `json:"id"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
}

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "transactions"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

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
// gin.Context is the most important part of Gin.
// It carries request details, validates and serializes JSON, and more.
func getTransactions(c *gin.Context) {
	//Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.

	var transactions []transaction

	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Query error: %v", err)})
		return
	}
	defer rows.Close()

	//Loop through the rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var trans transaction

		if err := rows.Scan(&trans.ID, &trans.Type, &trans.Description, &trans.Date, &trans.Amount); err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Scan error: %v", err)})
			return
		}

		transactions = append(transactions, trans)
	}
	// c.IndentedJSON(http.StatusOK, transactions)
	c.JSON(200, transactions)
}

// saveTransaction adds a transaction from JSON received in the request body.
func saveTransaction(c *gin.Context) {

	// Call BindJSON to bind the received JSON to newTransaction.
	var newTransaction transaction

	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO transactions (type, description, date, amount) VALUES (?, ?, ?, ?)", newTransaction.Type, newTransaction.Description, newTransaction.Date, newTransaction.Amount)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Query error: %v", err)})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error accessing the ID of the transaction: %v", err)})
		return
	}
	c.JSON(201, gin.H{"id": id})
}

// getTransactionByType located the transaction whose type value matches the id
// paramenter sent by the client, then returns thst tarnsactions as a response
func getTransactionByType(c *gin.Context) {
	transactionType := c.Param("type")

	var results []transaction

	rows, err := db.Query("SELECT * FROM transactions WHERE type = ?", transactionType)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Query error: %v", err)})
		return
	}
	defer rows.Close()

	//Loop through the rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var trans transaction
		if err := rows.Scan(&trans.ID, &trans.Type, &trans.Description, &trans.Date, &trans.Amount); err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Scan error: %v", err)})
			return
		}
		results = append(results, trans)
	}
	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Row error: %v", err)})
	}
	if len(results) == 0 {
		c.JSON(404, gin.H{"message": "No transactions found for this type"})
		return
	}

	c.JSON(200, results)
}
