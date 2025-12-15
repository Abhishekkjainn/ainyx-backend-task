# Ainyx Solutions - Go Backend Intern Task

![Go Version](https://img.shields.io/badge/go-1.25.1-blue.svg) ![PostgreSQL](https://img.shields.io/badge/postgres-15-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg)

Welcome to my submission for the **Software Engineering Intern** role at Ainyx Solutions. This project is a robust, high-performance RESTful API designed to manage users, demonstrating production-grade Go practices, type-safe database interactions, and containerized deployment.

## 🚀 Key Features

*   **RESTful API**: Clean, standard-compliant endpoints for CRUD operations.
*   **Dynamic Age Calculation**: Real-time age computation based on Date of Birth (DoB) handling edge cases (leap years, exact dates).
*   **Type-Safe Database Access**: Powered by **SQLC** to generate Go code from SQL queries, preventing runtime SQL errors.
*   **Robust Validation**: Request payload validation using `go-playground/validator`.
*   **Structured Logging**: High-performance logging with **Uber Zap**, including request duration and unique Request IDs for tracing.
*   **Pagination**: Efficient list retrieval with configurable `page` and `limit` parameters.
*   **Dockerized**: Fully containerized with `Dockerfile` and `docker-compose` for one-command setup.

## 🛠️ Tech Stack

*   **Language**: Go (Golang) 1.25+
*   **Web Framework**: [Fiber](https://gofiber.io/) (Fast, Express-inspired)
*   **Database**: PostgreSQL 15
*   **ORM/DAO**: [SQLC](https://sqlc.dev/) (Compile SQL to Go)
*   **Logging**: [Uber Zap](https://github.com/uber-go/zap)
*   **Validation**: [Validator v10](https://github.com/go-playground/validator)

## 📂 Project Structure

The project follows the standard Golang project layout:

```
├── cmd/
│   └── server/       # Application entry point
├── config/           # Database connection and configuration
├── db/
│   ├── migrations/   # SQL Schema and Queries
│   └── sqlc/         # Auto-generated Go code for DB interaction
├── internal/
│   ├── handler/      # HTTP Request Handlers (Controllers)
│   ├── middleware/   # Request Logger, Request ID injection
│   ├── models/       # Data structures and domain logic (e.g., CalculateAge)
│   ├── routes/       # API Route definitions
│   └── service/      # Business logic layer
├── .env              # Environment variables
├── docker-compose.yml # Docker orchestration
└── Dockerfile        # Container definition
```

## ⚡ Getting Started

### Option 1: Docker (Recommended)

Run the entire application (API + Postgres) with a single command.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Abhishekkjainn/ainyx-backend-task.git
    cd ainyx-backend-task
    ```

2.  **Start services:**
    ```bash
    docker-compose up --build
    ```
    The server will start at `http://localhost:8080`.

### Option 2: Local Development

1.  **Prerequisites:** Ensure Go 1.25+ and PostgreSQL are installed.
2.  **Setup Database:**
    Create a database named `ainyx` and run the schema in `db/migrations/schema.sql`.
3.  **Configure:**
    Set environment variables (or create a `.env` file):
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=ainyx
    ```
4.  **Run:**
    ```bash
    go mod tidy
    go run ./cmd/server/main.go
    ```

## 🔌 API Documentation

### 1. Create User
**POST** `/users`

*   **Request Body:**
    ```json
    {
      "name": "Alice Wonderland",
      "dob": "1998-05-20"
    }
    ```
*   **Response (201 Created):**
    ```json
    {
      "id": 1,
      "name": "Alice Wonderland",
      "dob": "1998-05-20",
      "age": 27
    }
    ```

### 2. Get User
**GET** `/users/:id`

Retrieves a user and **dynamically calculates their age**.

*   **Response (200 OK):**
    ```json
    {
      "id": 1,
      "name": "Alice Wonderland",
      "dob": "1998-05-20",
      "age": 27
    }
    ```

### 3. List Users (Paginated)
**GET** `/users?page=1&limit=10`

*   **Response (200 OK):**
    ```json
    [
      { "id": 1, "name": "Alice...", "dob": "1998-05-20", "age": 27 },
      { "id": 2, "name": "Bob...", "dob": "1995-01-15", "age": 30 }
    ]
    ```

### 4. Update User
**PUT** `/users/:id`

*   **Request Body:**
    ```json
    {
      "name": "Alice Cooper",
      "dob": "1998-05-21"
    }
    ```

### 5. Delete User
**DELETE** `/users/:id`

*   **Response (204 No Content)**

## 🧪 Testing

The project includes unit tests for the core business logic (Age Calculation).

Run tests with:
```bash
go test -v ./internal/...
```

## 📝 Reasoning & Design

For a detailed explanation of the architectural choices, trade-offs, and implementation details, please refer to [reasoning.md](./reasoning.md).
