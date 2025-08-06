# Shorty - A Robust URL Shortener in Go

Shorty is a high-performance, production-ready URL shortening service built with Go. It provides a clean RESTful API for creating, managing, and tracking short links, complete with user authentication and advanced features. This project demonstrates best practices in Go backend development, including clean architecture, secure authentication, and concurrent processing.

## ✨ Features

- **Secure User Authentication:** User registration and login using JWT (JSON Web Tokens).
    
- **URL Shortening:** Create short, unique, and easy-to-share links.
    
- **Custom Short Codes:** Users can propose their own custom aliases for links.
    
- **Click Analytics:** Track the number of clicks for each short link.
    
- **Link Expiration:** Set an optional expiration time for links.
    
- **User Dashboard:** An endpoint for authenticated users to view all the links they have created.
    
- **RESTful API:** A well-defined API for all interactions.
    
- **High Performance:** Built with Go for speed and concurrency.
    

## 🛠️ Tech Stack

- **Backend:** Go
    
- **Router:** `chi`
    
- **Database:** PostgreSQL
    
- **Authentication:** `golang-jwt` for JWT implementation
    
- **Password Hashing:** `bcrypt`
    
- **Configuration:** `viper`
    

## 🚀 Getting Started

### Prerequisites

- Go (version 1.21 or newer)
    
- PostgreSQL (running locally or in Docker)
    
- Docker (optional, for database)
    

### Installation

1. **Clone the repository:**
    
    ```
    git clone https://github.com/your-username/shorty.git
    cd shorty
    ```
    
2. **Set up environment variables:**
    
    - Create a `.env` file in the root of the project.
        
    - Add the following configuration:
        
    
    ```
    # Server configuration
    SERVER_PORT=8080
    
    # Database connection
    DB_URL="postgres://user:password@localhost:5432/shorty?sslmode=disable"
    
    # JWT Secret Key (change this to a long, random string)
    JWT_SECRET="your-super-secret-key-for-jwt"
    ```
    
3. **Run database migrations:**
    
    - Connect to your PostgreSQL instance and execute the SQL commands found in the `/db/migrations` directory to create the `users` and `urls` tables.
        
4. **Install dependencies and run the application:**
    
    ```
    go mod tidy
    go run ./cmd/server/main.go
    ```
    
    The server will start on `http://localhost:8080`.
    

## ⚙️ API Endpoints

### Authentication

|            |                    |               |                                         |
| ---------- | ------------------ | ------------- | --------------------------------------- |
| **Method** | **Path**           | **Protected** | **Description**                         |
| `POST`     | `/api/v1/register` | No            | Creates a new user account.             |
| `POST`     | `/api/v1/login`    | No            | Authenticates a user and returns a JWT. |

### URLs

|   |   |   |   |
|---|---|---|---|
|**Method**|**Path**|**Protected**|**Description**|
|`POST`|`/api/v1/shorten`|**Yes**|Creates a new short link.|
|`GET`|`/{shortCode}`|No|Redirects to the original long URL.|
|`GET`|`/api/v1/stats/{shortCode}`|No|Gets click analytics for a short link.|
|`GET`|`/api/v1/my-links`|**Yes**|Gets all links created by the logged-in user.|

### Example Usage (`curl`)

1. **Register a user:**
    
    ```
    curl -X POST http://localhost:8080/api/v1/register \
    -H "Content-Type: application/json" \
    -d '{"username": "testuser", "email": "test@example.com", "password": "strongpassword"}'
    ```
    
2. **Login to get a token:**
    
    ```
    curl -X POST http://localhost:8080/api/v1/login \
    -H "Content-Type: application/json" \
    -d '{"email": "test@example.com", "password": "strongpassword"}'
    ```
    
    _(Copy the token from the response for the next step)_
    
3. **Create a short link:**
    
    ```
    TOKEN="your-jwt-token-here"
    curl -X POST http://localhost:8080/api/v1/shorten \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d '{"url": "https://www.google.com/search?q=golang"}'
    ```
    

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://www.google.com/search?q=LICENSE "null") file for details.
