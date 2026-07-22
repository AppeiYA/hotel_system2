# Go Hotel Management System

This project is the backend for a modern Hotel Management System, built with Go. It follows the principles of Clean Architecture and Domain-Driven Design (DDD) to create a system that is scalable, maintainable, and robust.

## Core Concepts & Architecture

The system is designed around a modular, domain-centric architecture that separates business logic from infrastructure concerns. This makes the core logic independent of the database, web framework, or any external services.

### Key Architectural Principles

- **Clean Architecture:** The dependencies flow inwards. The core `domain` and `use_case` layers have no knowledge of outer layers like `adapters` (HTTP, database).
- **Domain-Driven Design (DDD):** The code is structured around the business domains of a hotel, such as `Reservation`, `Room`, `Payment`, and `Ledger`.
- **Ports & Adapters (Hexagonal Architecture):** The application communicates with the outside world through `ports` (interfaces) which are implemented by `adapters` (e.g., a PostgreSQL repository adapter for a `RoomRepository` port).

### Core Domains

- **Reservation:** Manages booking lifecycle, including creation, confirmation, check-in, and check-out. It prevents booking overlaps to ensure room availability.
- **Room:** Manages room inventory, status (e.g., `Available`, `Occupied`, `Cleaning`), types, and rates.
- **Payment:** Handles payment processing for reservations.
- **Ledger:** A sophisticated double-entry accounting system that provides a complete, immutable financial record of all transactions within the hotel. Every financial event (payment, refund, charge) is recorded as a balanced transaction, ensuring financial integrity.

## Features

- **Reservation Management:** Create, view, list, check-in, and check-out reservations.
- **Room Status Management:** A state machine enforces valid transitions for room statuses (e.g., an `Occupied` room must be marked for `Cleaning` before becoming `Available` again).
- **Double-Entry Accounting:** All financial operations are recorded in a ledger, ensuring that debits and credits are always balanced.
- **Transactional Integrity:** Business operations that span multiple database tables (e.g., creating a reservation and posting to the ledger) are executed within a single atomic transaction to prevent data corruption.
- **RESTful API:** A clear and documented API for interacting with the system.
- **Structured Logging:** Configurable, structured logging for clear and effective monitoring in both development and production environments.
- **API Documentation:** API endpoints are documented using Swagger/OpenAPI specifications.

## Tech Stack

- **Language:** Go
- **Web Framework:** Fiber
- **Database:** PostgreSQL
- **Database Driver/Toolkit:** sqlx
- **Logging:** Zap
- **API Documentation:** Swag

## Project Structure

The project follows a standard layout for Go applications, with a clear separation of concerns.

```
├── cmd/
│   └── api/              # Main application entrypoint
├── internal/
│   ├── ledger/           # Ledger domain, use cases, and adapters
│   ├── payment/          # Payment domain, use cases, and adapters
│   ├── reservation/      # Reservation domain, use cases, and adapters
│   ├── room/             # Room domain, use cases, and adapters
│   └── shared/           # Shared components (config, db, http, logger, etc.)
│       ├── adapters/
│       ├── config/
│       ├── db/
│       ├── domain/
│       ├── errors/
│       ├── http/
│       ├── logger/
│       └── ports/
├── pkg/                  # Reusable libraries safe for external use (if any)
├── go.mod
└── README.md
```

- `cmd/api`: Contains the `main` function to start the web server. It's responsible for wiring together all the components of the application (repositories, use cases, handlers, router).
- `internal/`: Contains all the core application code, separated by domain.
  - `<domain>/domain`: The core entities, value objects, and domain logic. Has zero external dependencies.
  - `<domain>/ports`: Defines the interfaces (ports) that the application layer uses to talk to the infrastructure layer.
  - `<domain>/use_case`: Implements the application-specific business rules, orchestrating domain objects.
  - `<domain>/adapters`: Contains the concrete implementations (adapters) of the ports, such as PostgreSQL repositories and Fiber HTTP handlers.
- `internal/shared`: Contains code shared across multiple domains, such as database transaction management, custom error types, and logging setup.

## Getting Started

### Prerequisites

- Go (version 1.21 or later)
- PostgreSQL
- Docker (optional, for running a local database)

### Installation

1.  **Clone the repository:**

    ```sh
    git clone <your-repo-url>
    cd hotel_system2
    go mod tidy
    ```

2.  **Set up environment variables:**

    Create a `.env` file in the root directory and add the necessary configuration.

    ```env
    # Application
    APP_ENV=development
    LOG_LEVEL=debug
    HTTP_PORT=8080

    # Database
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=hotel
    DB_SSL_MODE=disable
    ```

3.  **Run database migrations:**

    ```sh
    go run ./migrations --up/--down/--seed/--force value
    ```

4.  **Run the application:**
    ```sh
    go run ./cmd/api
    ```
    The server will start on the port specified in your `.env` file (e.g., `http://localhost:8080`).

## API Documentation

Run 
```bash
swag init -g ./cmd/api/main.go --parseInternal --parseDependency
```

Once the server is running, the Swagger API documentation is available at `/swagger/index.html`.

Example: `http://localhost:8080/swagger/index.html`
