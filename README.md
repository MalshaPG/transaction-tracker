# Transaction Tracker API

A REST API built with Go and Gin framework for managing financial transactions with full CRUD operations and MySQL database integration.

## Features

- **Full CRUD Operations**: Create, Read, Update, and Delete transactions
- **Filter by Type**: Get transactions by income/expense type
- **Input Validation**: Data validation with proper error handling
- **Modular Architecture**: Clean separation with controllers, models, routes
- **JSON API**: RESTful endpoints with JSON responses

## Tech Stack

- **Language**: Go 1.24.1
- **Framework**: Gin v1.10.0
- **Database**: MySQL
- **Driver**: go-sql-driver/mysql v1.9.2

## Project Structure

```
transaction-tracker/
├── controllers/           # HTTP handlers
├── database/             # DB connection
├── models/               # Data structures
├── routes/               # Route definitions
├── create-tables.sql     # Database schema
├── main.go              # Entry point
└── README.md
```

## Quick Start

1. **Clone and install:**
```bash
git clone https://github.com/MalshaPG/transaction-tracker.git
cd transaction-tracker
go mod tidy
```

2. **Set environment variables:**
```bash
# Linux/Mac
export DBUSER=username
export DBPASS=password

# Windows
set DBUSER=username
set DBPASS=password
```

3. **Setup database:**
Run the SQL from `create-tables.sql` to create the database and sample data.

4. **Run the application:**
```bash
go run main.go
```

Server starts on `localhost:8080`

## API Endpoints

| Method | Endpoint              | Description              |
| ------ | --------------------- | ------------------------ |
| GET    | `/transactions`       | Get all transactions     |
| GET    | `/transactions/:type` | Get transactions by type |
| POST   | `/transactions`       | Create a new transaction |
| PUT    | `/transactions/:id`   | Update a transaction     |
| DELETE | `/transactions/:id`   | Delete a transaction     |


## Usage Examples

**Get all transactions:**
```bash
curl http://localhost:8080/transactions
```

**Create transaction:**
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "type": "income",
    "description": "Salary",
    "date": "2025-01-20",
    "amount": 5000.00
  }'
```

**Update transaction:**
```bash
curl -X PUT http://localhost:8080/transactions/1 \
  -H "Content-Type: application/json" \
  -d '{
    "type": "expense",
    "description": "Updated expense",
    "date": "2025-01-20",
    "amount": 100.00
  }'
```

**Delete transaction:**
```bash
curl -X DELETE http://localhost:8080/transactions/1
```

## Transaction Object

```json
{
  "id": 1,
  "type": "income",           // "income" or "expense"
  "description": "Salary",    // required, max 200 chars
  "date": "2025-01-01",      // YYYY-MM-DD format
  "amount": 5000.00          // must be > 0
}
```

## Validation Rules

- **Type**: Must be "income" or "expense" (case-insensitive)
- **Description**: Required, non-empty, max 200 characters
- **Date**: Required, YYYY-MM-DD format
- **Amount**: Required, must be greater than 0

## Database Schema

```sql
CREATE TABLE transactions (
    id INT AUTO_INCREMENT NOT NULL,
    type VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    date DATE NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    CONSTRAINT pk_transactions PRIMARY KEY (id)   
);
```

## HTTP Status Codes

- `200 OK` - Successful GET/PUT
- `201 Created` - Successful POST
- `204 No Content` - Successful DELETE
- `400 Bad Request` - Invalid input
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server/DB error

## License

This is a personal project made available for educational purposes. Feel free to fork and use as you like.