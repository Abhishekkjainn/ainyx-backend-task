# Architectural Reasoning & Design Decisions

## 📖 Introduction
This document outlines the thought process, architectural decisions, and technical trade-offs made during the development of the Ainyx Go Backend Task. The goal was not just to satisfy the requirements, but to build a foundation that is maintainable, scalable, and adheres to Go best practices.

## 🏗️ Architecture: Standard Go Layout

I adopted the **Standard Go Project Layout** (`cmd/`, `internal/`) to separate concerns effectively.

*   **`cmd/server/main.go`**: This is the connection point. It strictly handles initialization (loading config, connecting to DB, wiring dependencies) and starting the server. It contains no business logic.
*   **`internal/`**: This directory ensures that the application code cannot be imported by external projects, enforcing encapsulation.
    *   **Layered Architecture**:
        1.  **Handler/Controller**: Parses HTTP requests, validates input, and sends responses. It knows *nothing* about the database.
        2.  **Service/Business Logic**: Contains the core domain logic (e.g., orchestrating creating a user). It sits between the handler and the data layer.
        3.  **Repository/Data Layer (SQLC)**: Handles raw database interactions.

**Why this approach?**
This separation allows for easier testing and refactoring. For example, if we switched from Postgres to MySQL, we would mostly touch the Repository layer. If we switched from REST to gRPC, we would mostly touch the Handler layer.

## 🛠️ Technology Stack Recommendations

### 1. Web Framework: Fiber vs. Gin vs. Stdlib
*   **Decision**: **Fiber**.
*   **Reasoning**: Fiber was chosen for its extreme performance (based on Fasthttp) and its ergonomic, Express.js-like API. For a task requiring rapid development and clean middleware handling (like logging and CORS), Fiber allows for very concise code compared to the standard `net/http` library.

### 2. Database Interaction: SQLC vs. GROM
*   **Decision**: **SQLC**.
*   **Reasoning**:
    *   **Type Safety**: SQLC compiles raw SQL queries into type-safe Go code. If a query is invalid or column names change, the build fails *immediately*, not at runtime.
    *   **Performance**: Unlike GORM, which relies heavily on reflection, SQLC generates pure struct-scanning code, making it significantly faster and lower overhead.
    *   **Clarity**: Writing raw SQL (in `db/migrations`) is often clearer/more powerful than learning a specific ORM DSL.

### 3. Logging: Uber Zap
*   **Decision**: **Zap**.
*   **Reasoning**: Standard `log` package is insufficient for structured logging. Zap provides structured, leveled logging with zero allocation overhead in critical paths. Combined with a `RequestID` middleware, this enables powerful traceability in production logs.

## 🌟 Implementation Highlights

### Dynamic Age Calculation
Instead of storing "Age" (which becomes stale daily), I store `DOB`. The `CalculateAge` function in `pkg/models` uses `time.Now()` to compute the age on-the-fly.
*   **Logic**: It compares years, then checks if the current day/month is before the birthday to subtract one if necessary.
*   **Testing**: This critical logic is backed by Unit Tests covering edge cases (birthday today, birthday tomorrow, leap years).

### Pagination (Bonus)
I implemented offset-based pagination (`Limit` and `Offset`) in the `ListUsers` endpoint.
*   The SQL query was updated to accept `$1` (limit) and `$2` (offset).
*   The Service layer calculates `offset = (page - 1) * limit`.
*   This prevents the API from crashing when the dataset grows to thousands of users.

### Docker Support (Bonus)
A Multi-Stage Dockerfile was used to minimize image size.
1.  **Builder Stage**: Compiles the Go binary.
2.  **Final Stage**: Uses `alpine` (approx 5MB) and copies *only* the binary and .env file.
This results in a tiny, secure production image.

## ⚖️ Trade-offs & Future Improvements

1.  **Configuration**: Currently using `godotenv`. In a larger system, I would use simpler tooling like `kelseyhightower/envconfig` or `viper` for hierarchical config (file vs env vs flags).
2.  **Migrations**: Schema is defined in `schema.sql`. For a real production app, I would add a migration tool like `golang-migrate` to manage versioned database changes up/down.
3.  **Testing**: Added unit tests for logic. Integration tests spinning up a test container DB would be the next step for 100% reliability.
