# Ainyx Solutions - Go Backend Intern Task

![Go Version](https://img.shields.io/badge/go-1.25.1-blue.svg) ![PostgreSQL](https://img.shields.io/badge/postgres-15-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg)

Welcome to my submission for the **Software Engineering Intern** role at Ainyx Solutions. This project is a robust, high-performance RESTful API designed to manage users, demonstrating production-grade Go practices, type-safe database interactions, and containerized deployment.

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
go test -v ./internal/...
```

## 📝 Reasoning & Design

For a detailed explanation of the architectural choices, trade-offs, and implementation details, please refer to [reasoning.md](./reasoning.md).
