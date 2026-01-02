# Go URL Shortener Service

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A high-performance, production-ready URL Shortener backend engineered with **Go** and **Gin**. This project demonstrates industry best practices for building scalable microservices, featuring a **Clean Layered Architecture** (Handler → Service → Repository) to ensure maintainability, testability, and clear separation of concerns.

Designed to handle high-throughput link generation and resolution, it utilizes **Base62 encoding** for efficient short codes and **PostgreSQL** for reliable persistent storage.

## 🚀 Key Features

- **⚡️ High Performance**: Built on the Gin framework for minimal latency.
- **🏗 Clean Architecture**: Strickland separation of concerns (Handler, Service, Repository, Model).
- **🐘 Persistent Storage**: Robust data management using PostgreSQL.
- **🔢 Base62 Shortening**: Efficient, collision-resistant short code generation.
- **📊 Analytics Tracking**: Automatic click tracking for every redirected URL.
- **🛡 Error Handling**: Domain-level error handling mapped to semantic HTTP status codes.
- **💓 Health Monitoring**: Built-in endpoints for container orchestration health checks.

## 🛠️ Technology Stack

- **Language**: [Go (Golang)](https://golang.org/) 1.21+
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Driver**: [lib/pq](https://github.com/lib/pq)
- **Configuration**: Environment-based (12-Factor App compliant)

## 📂 Project Structure

```bash
.
├── cmd
│   └── server          # Application entry point (main.go)
├── internal
│   ├── config          # Configuration loading (env vars)
│   ├── handler         # HTTP Transport layer (REST Controllers)
│   ├── service         # Business Logic layer (Validation, Algorithm)
│   ├── repository      # Data Access layer (SQL queries)
│   └── model           # Core domain entities
├── migrations          # SQL Database migration scripts
└── go.mod              # Module dependencies
```

## 🏁 Getting Started

### Prerequisites

- **Go**: v1.21 or higher
- **PostgreSQL**: Running instance (local or remote)
- **Git**

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/pranavgawaii/go-url-shortener.git
    cd go-url-shortener
    ```

2.  **Install dependencies**
    ```bash
    go mod tidy
    ```

3.  **Setup Configuration**
    Set the required environment variables:
    ```bash
    export PORT=8080
    export DATABASE_URL="postgres://user:password@localhost:5432/shortener?sslmode=disable"
    ```

4.  **Database Migration**
    Run the SQL script in `migrations/001_create_urls_table.sql` on your PostgreSQL database to create the required schema.

5.  **Run the application**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on port `8080`.

## 🔌 API Documentation

### 1. Shorten URL
Generates a unique 6-character short code for a given URL.

- **Endpoint**: `POST /api/shorten`
- **Content-Type**: `application/json`

**Request Body**
```json
{
  "url": "https://www.google.com/search?q=golang"
}
```

**Success Response (201 Created)**
```json
{
  "short_url": "http://localhost:8080/aX9d21"
}
```

**Error Response (400 Bad Request)**
```json
{
  "error": "original URL is required"
}
```

---

### 2. Redirect to Original URL
Redirects the client to the original URL and increments the click count.

- **Endpoint**: `GET /:shortCode`
- **Example**: `GET /aX9d21`

**Behavior**
- **Success**: redirects to original URL (HTTP 302 Found).
- **Not Found**: returns HTTP 404 (if code does not exist).

---

### 3. Health Check
Liveness probe for load balancers.

- **Endpoint**: `GET /health`
- **Response**: `200 OK`
```json
{
  "status": "ok"
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an issue for discussion.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.