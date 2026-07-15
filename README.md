# Products API PostgreSQL

REST API for product management built with Go, PostgreSQL and Docker.

The project implements user authentication using JWT, password hashing with bcrypt and role-based authorization through HTTP middleware. It demonstrates a layered backend architecture with handlers, middleware, storage layer and PostgreSQL.

---

# Tech Stack

- Go
- PostgreSQL
- Docker
- Docker Compose
- JWT (golang-jwt)
- bcrypt
- net/http
- database/sql
- godotenv
- lib/pq

---

# Features

### Authentication

- User registration
- User login
- Password hashing (bcrypt)
- JWT authentication
- JWT validation
- Authentication middleware
- Role-based authorization (user/admin)
- Protected routes

### Products

- Create product
- Get all products
- Get product by ID
- Update product
- Delete product
- Get products count
- Clear all products

### General

- PostgreSQL storage layer
- Docker Compose support
- Environment variables configuration
- JSON API
- Request validation

---

# Project Structure

```text
products-api-postgres
├── internal
│
├── auth
│   ├── jwt.go
│   └── password.go
│
├── database
│   └── database.go
│
├── handlers
│   ├── auth_handlers.go
│   └── products_handler.go
│
├── middleware
│   └── auth_middleware.go
│
├── models
│   ├── product.go
│   ├── user.go
│   └── login.go
│
├── response
│   └── response.go
│
├── storage
│   ├── product_storage.go
│   └── user_storage.go
│
├── Dockerfile
├── docker-compose.yml
├── init.sql
├── solution.sql
├── .env.example
├── main.go
└── README.md
```

---

# Architecture

```text
                Client
                   │
                   ▼
        HTTP Request
                   │
                   ▼
        Auth Middleware
        (JWT Validation)
                   │
                   ▼
      Admin Middleware
      (Role Validation)
                   │
                   ▼
              Handler
                   │
                   ▼
              Storage
                   │
                   ▼
            PostgreSQL
                   │
                   ▼
              Storage
                   │
                   ▼
              Handler
                   │
                   ▼
          HTTP JSON Response
```

---

# Layers

## Middleware

Responsible for:

- reading Authorization header
- validating JWT
- parsing token
- extracting user claims
- checking user role
- protecting API endpoints

---

## Handler

Responsible for:

- HTTP methods validation
- request validation
- JSON decoding
- calling storage methods
- returning JSON responses

---

## Storage

Responsible for:

- executing SQL queries
- communicating with PostgreSQL
- returning models and errors

---

## Response

Responsible for:

- JSON formatting
- unified error responses

---

# Authentication

The API uses JWT (JSON Web Token).

Authentication flow:

1. User registers.
2. Password is hashed using bcrypt.
3. User logs in.
4. Server validates credentials.
5. JWT token is generated.
6. Client stores the token.
7. Client sends

```
Authorization: Bearer <JWT_TOKEN>
```

for every protected request.

---

# Roles

## user

Can:

- Get products
- Get product by ID
- Get products count

---

## admin

Can additionally:

- Create product
- Update product
- Delete product
- Clear all products

---

# Environment Variables

Create a `.env` file based on `.env.example`.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=products_api
DB_SSLMODE=disable

JWT_SECRET=your_secret_key
```

---

# Run with Docker

```bash
docker compose up --build
```

Docker Compose starts:

- Go API
- PostgreSQL
- Docker volume
- Database initialization

API:

```
http://localhost:8080
```

Stop:

```bash
docker compose down
```

Remove database volume:

```bash
docker compose down -v
```

---

# API Endpoints

| Method | Endpoint | Access |
|---------|----------|--------|
| POST | /register | Public |
| POST | /login | Public |
| GET | /getProducts | Authenticated |
| GET | /getProduct?id=1 | Authenticated |
| GET | /getProductsCount | Authenticated |
| POST | /addProduct | Admin |
| PUT | /updateProduct?id=1 | Admin |
| DELETE | /deleteProduct?id=1 | Admin |
| DELETE | /clearProducts | Admin |

---

# Request Examples

## Register

```http
POST /register
```

```json
{
    "email":"admin@mail.com",
    "password":"12345678"
}
```

---

## Login

```http
POST /login
```

```json
{
    "email":"admin@mail.com",
    "password":"12345678"
}
```

Response

```json
{
    "token":"eyJhbGc..."
}
```

---

## Authorized Request

```http
Authorization: Bearer eyJhbGc...
```

---

## Create Product

```json
{
    "title":"Laptop",
    "price":1500,
    "in_stock":true
}
```

---

# Validation

The API validates:

## Product

- title cannot be empty
- price must be positive
- ID must be valid
- JSON must be valid

## Authentication

- Authorization header is required
- JWT token must be valid
- JWT token must not be expired
- Password must match bcrypt hash

## Authorization

- Admin routes require admin role

---

# Docker Notes

Local mode

Environment variables are loaded from `.env`.

Docker mode

Variables are passed through `docker-compose.yml`.

`.env` is not copied into the Docker image.

---

# Future Improvements

- Database migrations
- Unit tests
- Integration tests
- Redis caching
- Swagger / OpenAPI
- Request logging middleware
- Graceful shutdown
- Configuration package
- Refresh tokens

---

# Learning Goals

This project demonstrates practical implementation of:

- Layered Architecture
- PostgreSQL
- Docker
- JWT Authentication
- HTTP Middleware
- Role-Based Authorization
- Password Hashing (bcrypt)
- Environment Variables
- REST API Design
- JSON APIs