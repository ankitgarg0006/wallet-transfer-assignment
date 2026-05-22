# Wallet Transfer Service

A robust, highly concurrent wallet-to-wallet transfer service built in Golang. This service is designed to handle strict transactional guarantees, prevent double-spending, maintain a perfect double-entry ledger, and provide exactly-once semantics via a two-tiered idempotency system.

For a deep dive into the architecture, database schema, concurrency strategy, and tradeoffs, please read the [Architecture & Design Document](DESIGN.md).

## Core Features
*   **Exactly-Once Semantics:** Dual-tier idempotency using a fast-path in-memory cache and a strict PostgreSQL `UNIQUE` constraint fallback.
*   **Concurrency Safe:** Utilizes Pessimistic Row-Level Locking (`SELECT ... FOR UPDATE`) with Lexicographical Lock Ordering to prevent deadlocks and race conditions.
*   **Double-Entry Ledger:** Enforces a perfectly balanced, insert-only ledger using PostgreSQL triggers to prevent modifications or deletions.
*   **Clean Architecture:** Strict separation of concerns (Handlers -> Controllers -> Repositories).

## Getting Started

### Prerequisites
*   **Go** (1.23+)
*   **Docker & Docker Desktop** (Required for the local database and Testcontainers).

### 1. Environment Configuration
Navigate to the `config/` directory and create your environment file:
```bash
cp config/.env.example config/.env
```
*Open `config/.env` and fill in your local PostgreSQL credentials.*

### 2. Infrastructure Setup
Start the local PostgreSQL database using the provided `Makefile`:
```bash
make db-up
```

### 3. Run the Application
Start the main Gin server:
```bash
make run
```
*(Optional) Seed the database with initial test wallets:*
```bash
make seed
```

## Testing Strategy
The project employs a rigorous dual-tier testing strategy. **Docker must be running** for the integration tests.

**Run Isolated Unit Tests (Mocks, No DB):**
```bash
make test-unit
```

**Run Integration Tests (Real Postgres via Testcontainers):**
```bash
make test-integration
```

**Run All Tests:**
```bash
make test
```

## Future Scope
*   **Database Migrations:** Currently, the application uses GORM's `AutoMigrate` for schema initialization. A future enhancement will introduce a formal migration tool (like `golang-migrate/migrate` or `goose`) to manage versioned, raw SQL schema changes safely across environments.
*   **Authentication:** Adding centralized JWT/API key validation.
*   **Distributed Cache:** Transitioning the Tier 1 in-memory cache to Redis for multi-instance deployments.
