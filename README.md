# Products API PostgreSQL

REST API for managing products built with Go, PostgreSQL and Docker.

The project demonstrates a simple backend architecture with separated handlers, storage layer, database connection and JSON responses.

## Tech Stack

- Go
- PostgreSQL
- Docker
- Docker Compose
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
- Docker Compose support
- Environment variables configuration
- Basic request validation
- JSON responses

## Project Structure

```text
products-api-postgres
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
├── Dockerfile
├── docker-compose.yml
├── init.sql
├── solution.sql
├── .dockerignore
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Architecture

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

### Layers

**Handler layer**

- checks HTTP methods
- reads query parameters
- decodes JSON body
- validates request data
- returns JSON responses

**Storage layer**

- works with PostgreSQL
- executes SQL queries
- returns data or errors to handlers

**Response layer**

- formats JSON responses
- formats error responses

## Environment Variables

Create a `.env` file based on `.env.example`.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=products_api
DB_SSLMODE=disable
```

Do not commit your real `.env` file to GitHub.

## Run Locally

Create database:

```sql
CREATE DATABASE products_api;
```

Connect to database:

```sql
\c products_api
```

Create table:

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    price NUMERIC(10, 2) CHECK (price > 0),
    in_stock BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Install dependencies:

```bash
go mod tidy
```

Run the server:

```bash
go run main.go
```

Server will be available at:

```text
http://localhost:8080
```

## Run with Docker

Start API and PostgreSQL with Docker Compose:

```bash
docker compose up --build
```

Docker Compose starts:

- Go API container
- PostgreSQL container
- PostgreSQL volume
- database initialization from `init.sql`

API will be available at:

```text
http://localhost:8080
```

PostgreSQL will be available from host machine at:

```text
localhost:5433
```

Inside Docker network, API connects to PostgreSQL using:

```env
DB_HOST=postgres
DB_PORT=5432
```

Stop containers:

```bash
docker compose down
```

Stop containers and remove database volume:

```bash
docker compose down -v
```

`docker compose down -v` removes PostgreSQL data.

## API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/addProduct` | Create product |
| GET | `/getProducts` | Get all products |
| GET | `/getProduct?id=1` | Get product by ID |
| PUT | `/updateProduct?id=1` | Update product by ID |
| DELETE | `/deleteProduct?id=1` | Delete product by ID |
| GET | `/getProductsCount` | Get products count |
| DELETE | `/clearProducts` | Delete all products |

## Request Examples

### Create Product

```powershell
Invoke-RestMethod -Method POST http://localhost:8080/addProduct `
  -ContentType "application/json" `
  -Body '{"title":"Laptop","price":1500}'
```

Example response:

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

```powershell
Invoke-RestMethod -Method GET http://localhost:8080/getProducts
```

Example response:

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

```powershell
Invoke-RestMethod -Method GET "http://localhost:8080/getProduct?id=1"
```

### Update Product

```powershell
Invoke-RestMethod -Method PUT "http://localhost:8080/updateProduct?id=1" `
  -ContentType "application/json" `
  -Body '{"title":"Updated Laptop","price":1700,"in_stock":true}'
```

### Delete Product

```powershell
Invoke-RestMethod -Method DELETE "http://localhost:8080/deleteProduct?id=1"
```

### Get Products Count

```powershell
Invoke-RestMethod -Method GET http://localhost:8080/getProductsCount
```

### Clear Products

```powershell
Invoke-RestMethod -Method DELETE http://localhost:8080/clearProducts
```

## Validation

The API validates:

- product title must not be empty
- product price must be greater than zero
- product ID must be a valid integer
- request body must contain valid JSON
- HTTP method must match the endpoint

## Docker Notes

The application can work in two modes:

### Local mode

Environment variables are loaded from `.env`:

```go
_ = godotenv.Load()
```

### Docker mode

Environment variables are passed from `docker-compose.yml`:

```yaml
environment:
  DB_HOST: postgres
  DB_PORT: 5432
  DB_USER: postgres
  DB_PASSWORD: postgres
  DB_NAME: products_api
  DB_SSLMODE: disable
```

That is why `.env` is not copied into the Docker image.

## Next Improvements

- JWT authentication
- Auth middleware
- Protected routes
- Request logging middleware
- Unit tests
- Better error handling
- Swagger documentation