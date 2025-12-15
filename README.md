# Ainyx Solutions - Go Backend Intern Task

![Go Version](https://img.shields.io/badge/go-1.25.1-blue.svg) ![PostgreSQL](https://img.shields.io/badge/postgres-15-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg) ![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)

Welcome to the **Ultimate Submission** for the **Software Engineering Intern** role at Ainyx Solutions. This project is a production-grade, highly scalable RESTful API designed to demonstrate advanced Go patterns, including the **Repository Pattern**, **Dependency Injection**, and **Graceful Shutdown**.

## 🚀 Key Features & Bonuses

*   **RESTful API**: Clean, standard-compliant endpoints for User CRUD operations.
*   **Dynamic Age Calculation**: Real-time age computation based on Date of Birth (DoB), handling edge cases like leap years.
*   **Repository Pattern**: Decoupled data layer using Interfaces, allowing for easy mocking and testing.
*   **Dependency Injection**: Service and Handler layers are injected with their dependencies, strictly avoiding global state.
*   **Type-Safe Database**: Powered by **SQLC** to generate Go code from SQL, preventing runtime queries errors.
*   **Graceful Shutdown**: Handles `SIGINT`/`SIGTERM` to allow active requests to complete before shutting down.
*   **Safety Middleware**: Panic recovery middleware ensures the server never crashes due to handler errors.
*   **Pagination**: Efficient list retrieval with `page` and `limit` support.
*   **Dockerized**: Production-ready `Dockerfile` (Multi-stage) and `docker-compose`.

## 🛠️ Tech Stack

*   **Language**: Go (Golang) 1.25+
*   **Web Framework**: [Fiber](https://gofiber.io/) (High-performance, Express-style)
*   **Database**: PostgreSQL 15
*   **Data Access**: [SQLC](https://sqlc.dev/) (Type-safe SQL compiler)
*   **Logging**: [Uber Zap](https://github.com/uber-go/zap) (Structured, zero-allocation)
*   **Validation**: [Validator v10](https://github.com/go-playground/validator)

## 📂 Project Structure

Verified "Standard Go Project Layout":

```
├── cmd/
│   └── server/       # Application entry point (Wires DI containers)
├── config/           # Database connection logic
├── db/
│   ├── migrations/   # SQL Schema and Queries
│   └── sqlc/         # Auto-generated Go code (Models & DBTX)
├── internal/
│   ├── handler/      # HTTP Controllers (Validates input, calls Service)
│   ├── middleware/   # Request Logger, Panic Recovery, CORS
│   ├── models/       # Domain objects & Logic (CalculateAge)
│   ├── repository/   # Data Access Layer (Implements UserRepository Interface)
│   ├── routes/       # API Route setup
│   └── service/      # Business Logic (Implements UserService Interface)
├── .env              # Configuration
├── docker-compose.yml # Orchestration
└── Dockerfile        # Multi-stage build
```

## ⚡ Getting Started

### Option 1: Docker (Fastest)

Run the entire stack without installing Go or Postgres locally.

1.  **Clone & Run:**
    ```bash
    git clone https://github.com/Abhishekkjainn/ainyx-backend-task.git
    cd ainyx-backend-task
    docker-compose up --build
    ```
    API will be available at `http://localhost:8080`.

### Option 2: Local Development

1.  **Prerequisites:** Go 1.25+, PostgreSQL running globally.
2.  **Setup DB:** Create a DB named `ainyx` and run `db/migrations/schema.sql`.
3.  **Environment:**
    Create a `.env` file:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=ainyx
    ```
4.  **Run:**
    ```bash
    go mod tidy
    go run ./cmd/server/main.go
    ```

## 🔌 API Documentation

### User Object
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1990-01-01",
  "age": 34
}
```

### Endpoints

| Method | Endpoint | Description | Params |
| :--- | :--- | :--- | :--- |
| `POST` | `/users` | Create a new user | Body: `{ "name": "...", "dob": "YYYY-MM-DD" }` |
| `GET` | `/users/:id` | Get user by ID (Calc Age) | |
| `GET` | `/users` | List users (Paginated) | `?page=1&limit=10` |
| `PUT` | `/users/:id` | Update user details | Body: `{ "name": "...", "dob": "YYYY-MM-DD" }` |
| `DELETE` | `/users/:id` | Soft/Hard delete user | |

## 🧪 Testing

Unit tests are included for critical domain logic (Age Calculation).

```bash
# Run all tests
go test -v ./internal/...
```

## � Architectural Reasoning

This project uses a **Layered Architecture**:

1.  **Handler**: Knows about HTTP (Status Codes, JSON). Calls Service.
2.  **Service**: Knows about Business Logic (Validation, Calc Age). Calls Repository.
3.  **Repository**: Knows about Database (SQL). Returns Entities.

For a deep dive into why `SQLC` was chosen over `GORM`, or why `Fiber` was selected, please read **[reasoning.md](./reasoning.md)**.
