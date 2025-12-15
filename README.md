# Ainyx User Management API

A RESTful API built with Go, Fiber, PostgreSQL, and SQLC.

## Features
- User CRUD operations
- Dynamic Age Calculation
- Input Validation
- Structured Logging (Zap)
- Request Tracing (X-Request-ID)

## Prerequisites
- Go 1.25+
- PostgreSQL

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd project
   ```

2. **Database Setup**:
   Ensure Postgres is running and creating a database. Update `.env` or environment variables with connection string.

3. **Run Migrations** (if applicable) or load schema:
   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       dob DATE NOT NULL
   );
   ```

4. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

5. **Run the Server**:
   ```bash
   go run ./cmd/server/main.go
   ```
   Server starts at `http://localhost:8080`

## API Endpoints

- **POST /users**: Create a user
- **GET /users/:id**: Get user details (incl. Age)
- **PUT /users/:id**: Update user
- **DELETE /users/:id**: Delete user
- **GET /users**: List all users

## Testing
Run unit tests:
```bash
go test ./internal/...
```
