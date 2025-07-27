
# Production-Ready REST API: "Shorty" - A URL Shortener

This is a classic project that touches on all the fundamentals of backend development in a clean, understandable package.

### 🎯 Objective

Build a robust, testable, and configurable RESTful API that takes a long URL and returns a short, unique code that redirects back to the original URL.

### 🛠️ Core Concepts Demonstrated

- REST API Design (CRUD operations)
    
- Database Interaction (SQL)
    
- Configuration Management (from files or environment variables)
    
- Structured Logging
    
- Routing & Middleware
    
- Unit & Integration Testing
    
- Graceful Shutdown
    

### 💻 Technology Stack

- **Go:** For the core logic.
    
- **Router:** `chi` (A great, lightweight router that embraces standard library concepts).
    
- **Database:** PostgreSQL (A powerful, open-source SQL database).
    
- **DB Driver/Toolkit:** `sqlx` (An excellent extension to the standard `database/sql` package).
    
- **Configuration:** `viper` (For handling configuration from files and environment variables).
    
- **Logging:** `slog` (Go's new structured logging library).
    

---

### 📋 Step-by-Step Guide

1. **Project Setup & Configuration**
    
    - Initialize your Go module: `go mod init github.com/your-username/shorty`
        
    - Install dependencies: `go get github.com/go-chi/chi/v5 github.com/jmoiron/sqlx github.com/lib/pq github.com/spf13/viper`
        
    - Set up `viper` to read a `config.yaml` file. This file should contain the server port and database connection string.
        
2. **Database Layer**
    
    - Define your database schema in a `.sql` file. You'll need one table, `urls`, with columns like `id` (auto-incrementing integer, primary key), `short_code` (text, unique index), `original_url` (text), and `created_at` (timestamp).
        
    - In a new `storage` package, create a `PostgresStore` struct that holds your `*sqlx.DB` connection.
        
    - Write the methods for this struct:
        
        - `SaveURL(originalURL string) (string, error)`: Inserts the URL, generates a `short_code` (e.g., by base62-encoding the `id`), and returns the code.
            
        - `GetURL(shortCode string) (string, error)`: Queries the database for a `short_code` and returns the `original_url`.
            
3. **HTTP Handler Layer**
    
    - In a new `handler` package, create your HTTP handlers. They should accept the storage layer as a dependency.
        
    - `ShortenHandler(w http.ResponseWriter, r *http.Request)`: Decodes a JSON request, calls `storage.SaveURL()`, and writes a JSON response with the short code.
        
    - `RedirectHandler(w http.ResponseWriter, r *http.Request)`: Gets the `shortCode` from the URL path, calls `storage.GetURL()`, and performs a `301 Moved Permanently` redirect.
        
4. **Main Application (`cmd/server/main.go`)**
    
    - Load configuration, initialize the logger, and connect to the database.
        
    - Instantiate your `storage` and `handler` layers.
        
    - Set up your `chi` router with routes:
        
        - `router.Post("/shorten", handler.ShortenHandler)`
            
        - `router.Get("/{shortCode}", handler.RedirectHandler)`
            
    - Start the HTTP server and implement **graceful shutdown**.
        
5. **Testing**
    
    - Write **unit tests** for your storage functions.
        
    - Write **integration tests** for your HTTP handlers using the `net/http/httptest` package.
        

---

### 📂 Suggested Project Structure

Plaintext

```
/shorty
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── url_handler.go
│   ├── storage/
│   │   └── postgres.go
│   └── shortener/
│       └── shortener.go
├── go.mod
├── go.sum
└── config.yaml
```

---

### ✨ Going Further

- **Add a Rate Limiter:** Use middleware to limit requests.
    
- **Add Analytics:** Track how many times each short URL is clicked.
    
- **Swagger/OpenAPI Docs:** Generate API documentation for your endpoints.
