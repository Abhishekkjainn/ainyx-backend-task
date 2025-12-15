# Reasoning and Approach

## Overview
This repository contains the solution for the Ainyx Solutions Go Backend Development Task. The goal was to build a RESTful API for user management with dynamic age calculation, robust validation, and structured logging.

## Architectural Decisions

### 1. Project Structure
I adopted a standard Go project layout (`cmd`, `internal`, `pkg`) to separate concerns:
- **`cmd/server`**: Entry point, keeping main logic minimal.
- **`internal/`**: Private application code.
    - **`handler`**: HTTP transport layer, handles request/response.
    - **`service`**: Business logic (e.g., date parsing, age calculation).
    - **`repository` (via SQLC)**: Type-safe database access.
    - **`models`**: Data structures and validation rules.

### 2. Dependency Choices
- **Fiber**: Chosen for its performance and Expressjs-like ease of use.
- **SQLC**: Preferred over ORMs (like GORM) for performance and type-safety while writing raw SQL.
- **Uber Zap**: Low-allocation, high-performance structured logging.
- **go-playground/validator**: Industry standard for struct validation.

### 3. Implementation Details
- **Dynamic Age Calculation**: Implemented in the `models` package using `time.Now()` and specific date logic to account for years and days, ensuring accuracy.
- **Middleware**: Added a request logger middleware that also injects a UUID `X-Request-ID` for traceability.
- **Validation**: Struct tags are used to enforce `required` fields and date formats (`2006-01-02`).

## Key Trade-offs
- **SQLC vs ORM**: SQLC requires a compiled step (`sqlc generate`), but guarantees that queries match the schema.
- **Fiber vs Stdlib**: Fiber is not compatible with `net/http` interfaces out of the box, but offers better developer experience for routing and middleware.

## Future Improvements
- **Docker**: Containerization for consistent deployment.
- **Pagination**: Adding offset/limit to the List endpoint.
