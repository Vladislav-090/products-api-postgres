# Products API PostgreSQL

REST API for managing products, built with Go and PostgreSQL.

This project is a practical backend API created to improve Go, PostgreSQL, HTTP handlers, clean project structure, and database interaction skills.

## Tech Stack

- Go
- PostgreSQL
- net/http
- database/sql
- godotenv
- lib/pq

## Features

- Create product
- Get all products
- Get product by ID
- Update product by ID
- Delete product by ID
- Get products count
- Clear all products
- PostgreSQL storage layer
- JSON responses
- Basic request validation
- Environment variables support

## Project Structure

```text
products-api-postgres
├── main.go
├── internal
│   ├── database
│   │   └── database.go
│   ├── handlers
│   │   └── products_handler.go
│   ├── models
│   │   └── product.go
│   ├── response
│   │   └── response.go
│   └── storage
│       └── product_storage.go
├── solution.sql
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Environment Variables

Create a `.env` file in the project root based on `.env.example`.

Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=products_api
DB_SSLMODE=disable
```

Important: do not commit your real `.env` file to GitHub.

## Database Setup

Create database:

```sql
CREATE DATABASE products_api;
```

Connect to the database:

```sql
\c products_api
```

Create products table:

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    price NUMERIC(10, 2) CHECK (price > 0),
    in_stock BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Check table:

```sql
\dt
```

View products:

```sql
SELECT * FROM products;
```

## Run Project

Install dependencies:

```bash
go mod tidy
```

Run the server:

```bash
go run main.go
```

Server runs on:

```text
http://localhost:8080
```

## API Endpoints

### Create Product

```http
POST /addProduct
```

Request body:

```json
{
  "title": "Laptop",
  "price": 1500
}
```

Successful response example:

```json
{
  "message": "Product Created Succsessfully!",
  "product": {
    "id": 1,
    "title": "Laptop",
    "price": 1500,
    "in_stock": true,
    "created_at": "2026-07-12T09:30:18.42289Z"
  }
}
```

### Get All Products

```http
GET /getProducts
```

Successful response example:

```json
[
  {
    "id": 1,
    "title": "Laptop",
    "price": 1500,
    "in_stock": true,
    "created_at": "2026-07-12T09:30:18.42289Z"
  }
]
```

### Get Product By ID

```http
GET /getProduct?id=1
```

Successful response example:

```json
{
  "id": 1,
  "title": "Laptop",
  "price": 1500,
  "in_stock": true,
  "created_at": "2026-07-12T09:30:18.42289Z"
}
```

Error response example:

```json
{
  "error": "product not found!"
}
```

### Update Product By ID

```http
PUT /updateProduct?id=1
```

Request body:

```json
{
  "title": "Updated Laptop",
  "price": 1700,
  "in_stock": true
}
```

Successful response example:

```json
{
  "id": 1,
  "title": "Updated Laptop",
  "price": 1700,
  "in_stock": true,
  "created_at": "2026-07-12T09:30:18.42289Z"
}
```

### Delete Product By ID

```http
DELETE /deleteProduct?id=1
```

Successful response example:

```json
{
  "message": "Product deleted successfully!"
}
```

### Get Products Count

```http
GET /getProductsCount
```

Successful response example:

```json
{
  "count": 3
}
```

### Clear All Products

```http
DELETE /clearProducts
```

Successful response example:

```json
{
  "message": "All Products cleared successfully!"
}
```

## Example PowerShell Requests

Create product:

```powershell
Invoke-RestMethod -Method POST http://localhost:8080/addProduct `
  -ContentType "application/json" `
  -Body '{"title":"Laptop","price":1500}'
```

Get all products:

```powershell
Invoke-RestMethod -Method GET http://localhost:8080/getProducts
```

Get product by ID:

```powershell
Invoke-RestMethod -Method GET "http://localhost:8080/getProduct?id=1"
```

Update product:

```powershell
Invoke-RestMethod -Method PUT "http://localhost:8080/updateProduct?id=1" `
  -ContentType "application/json" `
  -Body '{"title":"Updated Laptop","price":1700,"in_stock":true}'
```

Delete product:

```powershell
Invoke-RestMethod -Method DELETE "http://localhost:8080/deleteProduct?id=1"
```

Get products count:

```powershell
Invoke-RestMethod -Method GET http://localhost:8080/getProductsCount
```

Clear all products:

```powershell
Invoke-RestMethod -Method DELETE http://localhost:8080/clearProducts
```

## Architecture

The project uses a simple layered structure:

```text
HTTP Request
    ↓
Handler
    ↓
Storage
    ↓
PostgreSQL
    ↓
Storage
    ↓
Handler
    ↓
HTTP Response
```

### Handler Layer

Responsible for:

- checking HTTP method
- reading URL query parameters
- decoding JSON body
- validating request data
- returning JSON responses

### Storage Layer

Responsible for:

- SQL queries
- working with PostgreSQL
- returning data or errors to handlers

### Response Layer

Responsible for:

- JSON response formatting
- error response formatting

## Current Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/addProduct` | Create new product |
| GET | `/getProducts` | Get all products |
| GET | `/getProduct?id=1` | Get product by ID |
| PUT | `/updateProduct?id=1` | Update product by ID |
| DELETE | `/deleteProduct?id=1` | Delete product by ID |
| GET | `/getProductsCount` | Get products count |
| DELETE | `/clearProducts` | Delete all products |

## Validation

The API validates:

- product title should not be empty
- product price should be greater than zero
- product ID should be valid integer
- request body should be valid JSON
- HTTP method should match the endpoint

## Next Improvements

Planned improvements:

- Docker support
- Docker Compose with PostgreSQL
- JWT authentication
- Auth middleware
- Protected routes
- Request logging middleware
- Unit tests
- Better error handling
- Swagger documentation