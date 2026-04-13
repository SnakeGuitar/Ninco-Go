# Database Documentation

This document describes the database connection and configuration for the Ninco-Go project.

## Overview

The project uses **MySQL** as its primary database. The database logic is located in `internal/platform/db/`.

- **Driver:** `github.com/go-sql-driver/mysql`
- **Package:** `db`

## Connection Configuration

The `Config` struct manages the connection settings:

```go
type Config struct {
    User            string        // Database username
    Password        string        // Database password
    Host            string        // IP or hostname (e.g., 127.0.0.1)
    Port            string        // Port number (default: 3306)
    Database        string        // Name of the database schema
    MaxIdleConns    int           // Maximum idle connections in the pool
    MaxOpenConns    int           // Maximum open connections in the pool
    ConnMaxLifetime time.Duration // Maximum amount of time a connection may be reused
}
```

## How to Connect

Use the `CreateConnection` function to initialize a new pool:

```go
db, err := db.CreateConnection(myConfig)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

## Environment Variables

For security and development ease, we use environment variables for credentials. You can find a template in `.env.example`.

| Variable | Description |
| :--- | :--- |
| `DB_USER` | MySQL Username |
| `DB_PASSWORD` | MySQL Password |
| `DB_HOST` | Host address |
| `DB_PORT` | Port (3306) |
| `DB_NAME` | Database Name |

## Testing

Tests are located in `internal/platform/db/mysql_test.go`.

### Running Tests (Windows PowerShell)
```powershell
$env:DB_USER="root"; $env:DB_HOST="127.0.0.1"; go test -v ./internal/platform/db/...
```

The tests will automatically **skip** if the mandatory environment variables are not set, preventing errors in non-db environments.

## Connection Pool Best Practices

1.  **Ping Early:** Always verify the connection using `db.Ping()` after opening (already handled by `CreateConnection`).
2.  **Close Properly:** Use `defer db.Close()` to ensure connections are returned to the system when the program shuts down.
3.  **Set Limits:** Use `MaxOpenConns` and `MaxIdleConns` to prevent the application from overwhelming the database.
