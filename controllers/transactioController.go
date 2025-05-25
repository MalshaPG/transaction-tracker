package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MalshaPG/transaction-tracker/database"
	"github.com/MalshaPG/transaction-tracker/models"
	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all transactions as JSON.
// gin.Context is the most important part of Gin.
// It carries request details, validates and serializes JSON, and more.
func GetTransactions(c *gin.Context) {
	//Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.

	var transactions []models.Transaction

	rows, err := database.DB.Query("SELECT * FROM transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Query error: %v", err)})
		return
	}
	defer rows.Close()

	//Loop through the rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var trans models.Transaction

		if err := rows.Scan(&trans.ID, &trans.Type, &trans.Description, &trans.Date, &trans.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Scan error: %v", err)})
			return
		}

		transactions = append(transactions, trans)
	}
	// return http.status code 500
	c.JSON(http.StatusOK, transactions)
}

// saveTransaction adds a transaction from JSON received in the request body.
func SaveTransaction(c *gin.Context) {

	// Call BindJSON to bind the received JSON to newTransaction.
	var newTransaction models.Transaction

	//Deserialize JSON into a Go struct using BindJSON.
	// If the JSON is invalid, BindJSON returns an error.
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid JSON: %v", err)})
		return
	}

	// Validate the transaction data
	if newTransaction.Type == "" || newTransaction.Amount <= 0 || newTransaction.Date == "" || newTransaction.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required and amount must be greater than 0"})
		return
	}

	// Validate the transaction type
	// Convert the transaction type to lowercase to ensure case-insensitivity
	transactionType := strings.ToLower(newTransaction.Type)
	if transactionType != "income" && transactionType != "expense" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction type must be 'income' or 'expense'"})
		return
	}

	//Overwrite the type
	newTransaction.Type = transactionType

	// Insert the transaction into the database
	result, err := database.DB.Exec("INSERT INTO transactions (type, description, date, amount) VALUES (?, ?, ?, ?)", newTransaction.Type, newTransaction.Description, newTransaction.Date, newTransaction.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Query error: %v", err)})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error accessing the ID of the transaction: %v", err)})
		return
	}
	//Return 201 http status code
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getTransactionByType located the transaction whose type value matches the id
// paramenter sent by the client, then returns thst tarnsactions as a response
func GetTransactionByType(c *gin.Context) {
	transactionType := strings.ToLower(c.Param("type"))

	var results []models.Transaction

	rows, err := database.DB.Query("SELECT * FROM transactions WHERE type = ?", transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error: %v", err)})
		return
	}
	defer rows.Close()

	//Loop through the rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var trans models.Transaction
		if err := rows.Scan(&trans.ID, &trans.Type, &trans.Description, &trans.Date, &trans.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Scan error: %v", err)})
			return
		}
		results = append(results, trans)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Row error: %v", err)})
	}
	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No transactions found for this type"})
		return
	}

	//send 200 OK http status code with the results
	c.JSON(http.StatusOK, results)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")

	var updatedTransaction models.Transaction
	if err := c.BindJSON(&updatedTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid JSON: %v", err)})
		return
	}

	// Validate the inputs
	if updatedTransaction.Type == "" || updatedTransaction.Amount <= 0 || updatedTransaction.Date == "" || updatedTransaction.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required and amount must be greater than 0"})
		return
	}

	transactionType := strings.ToLower(updatedTransaction.Type)
	if transactionType != "income" && transactionType != "expense" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction type must be 'income' or 'expense'"})
		return
	}
	// Overwrite the type
	updatedTransaction.Type = transactionType

	//Update the transaction in the database
	query := "UPDATE transactions SET type = ?, description = ?, date = ?, amount = ? WHERE id = ?"

	result, err := database.DB.Exec(query, updatedTransaction.Type, updatedTransaction.Description, updatedTransaction.Date, updatedTransaction.Amount, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Update failed: %v", err)})
		return
	}

	// Check if any rows were affected and return an appropriate response
	affectedRowCount, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error checking affected rows: %v", err)})
		return
	}

	if affectedRowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	}

	//return 200 OK status code with a message
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully", "id": id})

}

// DeleteTransaction removes a transaction from the database using the ID provided in the URL parameter.
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	result, err := database.DB.Exec("DELETE FROM transactions WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Delete failed: %v", err)})
		return
	}

	// Check if any rows were affected and return an appropriate response
	affectedRowCount, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if affectedRowCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	}

	//return a 204 No Content status code
	c.Status(http.StatusNoContent)
}
