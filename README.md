# Transaction Tracker API

A simple REST API built with Go and Gin framework for managing financial transactions. The API allows you to create, retrieve, and filter transactions stored in a MySQL database.

## Features

- Create new transactions
- Retrieve all transactions
- Filter transactions by type
- MySQL database integration
- JSON API responses
- Error handling and validation

## Tech Stack

- **Language**: Go 1.24.1
- **Web Framework**: Gin v1.10.0
- **Database**: MySQL
- **Database Driver**: go-sql-driver/mysql v1.9.2

## Prerequisites

Before running this application, make sure you have:

- Go 1.24.1 or later installed
- MySQL server running
- A MySQL database named `transactions`

## Database Setup

1. Create a MySQL database named `transactions`
2. Create the transactions table with the following structure:

```sql
CREATE DATABASE transactions;

USE transactions;

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id INT AUTO_INCREMENT NOT NULL,
    type VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    date DATE NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    CONSTRAINT pk_transactions PRIMARY KEY (id)   
);

-- Insert sample data
INSERT INTO transactions (type, description, date, amount) 
VALUES
('income', 'Salary', '2025-01-01', 5000.00),
('expense', 'Groceries', '2025-01-02', -150.00),
('income', 'Rent from 627/D', '2025-01-03', 1200.00),
('expense', 'Utilities', '2025-01-04', -200.00),
('income', 'Investment Return', '2025-01-05', 300.00),
('expense', 'Dining Out', '2025-01-06', -100.00);
```

## Environment Variables
Set the following environment variables for database connection:

**On Linux or Mac:**
```bash
$ export DBUSER=username
$ export DBPASS=password
```

**On Windows:**
```cmd
C:\Users\you\data-access> set DBUSER=username
C:\Users\you\data-access> set DBPASS=password
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/MalshaPG/transaction-tracker.git
cd transaction-tracker
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your environment variables:
```bash
export DBUSER=your_mysql_username
export DBPASS=your_mysql_password
```

4. Run the application:
```bash
go run main.go
```

The server will start on `localhost:8080`

## API Endpoints

### Get All Transactions
- **URL**: `GET /transactions`
- **Description**: Retrieves all transactions from the database
- **Response**: Array of transaction objects

**Example Response:**
```json
[
    {
        "id": 1,
        "type": "income",
        "description": "Salary",
        "date": "2025-01-01",
        "amount": 5000.00
    },
    {
        "id": 2,
        "type": "expense",
        "description": "Groceries",
        "date": "2025-01-02",
        "amount": -150.00
    }
]
```

### Create New Transaction
- **URL**: `POST /transactions`
- **Description**: Creates a new transaction
- **Request Body**: JSON object with transaction details

**Example Request:**
```json
{
    "type": "expense",
    "description": "Coffee purchase",
    "date": "2025-01-17",
    "amount": -4.50
}
```

**Example Response:**
```json
{
    "id": 3
}
```

### Get Transactions by Type
- **URL**: `GET /transactions/:type`
- **Description**: Retrieves all transactions of a specific type
- **Parameters**: 
  - `type` (string): The transaction type to filter by

**Example Request:**
```
GET /transactions/income
```

**Example Response:**
```json
[
    {
        "id": 1,
        "type": "income",
        "description": "Salary",
        "date": "2025-01-01",
        "amount": 5000.00
    },
    {
        "id": 3,
        "type": "income",
        "description": "Rent from 627/D",
        "date": "2025-01-03",
        "amount": 1200.00
    }
]
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- `200 OK`: Successful GET requests
- `201 Created`: Successful POST requests
- `404 Not Found`: When no transactions match the filter criteria
- `500 Internal Server Error`: Database or server errors

**Error Response Format:**
```json
{
    "error": "Error description"
}
```

## Testing the API

You can test the API using curl commands:

### Create a transaction:
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "type": "income",
    "description": "Freelance work",
    "date": "2025-01-18",
    "amount": 750.00
  }'
```

### Get all transactions:
```bash
curl http://localhost:8080/transactions
```

### Get transactions by type:
```bash
curl http://localhost:8080/transactions/income
```

## Project Structure

```
transaction-tracker/
├── main.go           # Main application file
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
└── README.md        # This file
```

## Future Enhancements

- [ ] Add UPDATE endpoint for modifying transactions
- [ ] Add DELETE endpoint for removing transactions
- [ ] Implement authentication and authorization
- [ ] Add input validation and sanitization
- [ ] Add pagination for large datasets
- [ ] Add transaction categories and tags
- [ ] Implement transaction search functionality
- [ ] Add API documentation with Swagger

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This is a personal project made available for educational and reference purposes. Feel free to fork, modify, and use as you like.

## Contact

For questions or support, please contact me or create an issue in the GitHub repository.
 