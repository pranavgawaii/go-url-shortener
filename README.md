<div align="center">

# Go URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/pranavgawaii/go-url-shortener/graphs/commit-activity)

**A production-grade, high-performance URL shortening microservice built with Go.**

[Key Features](#-key-features) • [Architecture](#-architecture) • [Getting Started](#-getting-started) • [API Reference](#-api-documentation)

</div>

---

## 📖 Project Overview

The **Go URL Shortener** is designed to demonstrate industry-standard practices for building scalable backend systems. Unlike simple tutorials, this project prioritizes **maintainability, testability, and performance** from day one.

It implements a **Clean Layered Architecture**, ensuring that business logic is decoupled from transport layers (HTTP) and data persistence (SQL). This separation allows for easier unit testing, refactoring, and potentially swapping out components (e.g., changing databases or web frameworks) with minimal friction.

### Why Go?
Go was chosen for its exceptional performance, low memory footprint, and high concurrency capabilities (goroutines), making it the ideal choice for a high-traffic service like a URL shortener that needs to handle thousands of redirects with sub-millisecond latency.

## 🏛️ Architecture

The application follows a strict unidirectional data flow:

```mermaid
graph LR
    A[HTTP Handler] -->|DTOs| B[Service Layer]
    B -->|Domain Models| C[Repository Layer]
    C -->|SQL| D[(PostgreSQL)]
```

*   **Handler Layer**: Manages HTTP requests/responses, payload validation, and status codes.
*   **Service Layer**: Contains pure business logic (e.g., Base62 encoding, collision handling, rules).
*   **Repository Layer**: Handles all direct data access, abstracting the underlying database technology.

## 🚀 Key Features

- **⚡️ High Performance**: Powered by **Gin**, optimized for speed and low allocation.
- **🔢 Efficient Shortening**: Uses **Base62 encoding** (a-zA-Z0-9) to generate short, URL-safe codes.
- **🐘 Robust Storage**: Data is persisted in **PostgreSQL** with proper indexing for fast lookups.
- **📊 Analytics Ready**: Built-in click tracking increments counters atomically on every redirect.
- **🛡 Enterprise Error Handling**: Semantic domain errors mapped to appropriate HTTP status codes (400 vs 404 vs 500).
- **💓 Production Ready**: Includes health check endpoints (`/health`) for container orchestration (Kubernetes/ECS).

## 🛠️ Tech Stack

- **Language**: [Go (Golang)](https://golang.org/) 1.21+
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Configuration**: 12-Factor App (Environment Variables)

## 📂 Project Structure

```bash
.
├── cmd
│   └── server          # Entry point (main.go)
├── internal
│   ├── config          # Configuration loader
│   ├── handler         # HTTP Controllers
│   ├── service         # Business Logic
│   ├── repository      # Data Access Objects (DAO)
│   └── model           # Domain Entities
├── migrations          # Database Schemas
└── go.mod              # Dependency Manager
```

## 🏁 Getting Started

### Prerequisites

- **Go**: v1.21 or higher
- **PostgreSQL**: Running instance
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
    ```bash
    export PORT=8080
    export DATABASE_URL="postgres://user:password@localhost:5432/shortener?sslmode=disable"
    ```

4.  **Database Migration**
    Execute the SQL in `migrations/001_create_urls_table.sql` on your database.

5.  **Run the application**
    ```bash
    go run cmd/server/main.go
    ```

## 🔌 API Documentation

### 1. Shorten URL
`POST /api/shorten`

Creates a new short link for the provided URL.

**Request:**
```json
{
  "url": "https://www.example.com/very/long/url"
}
```

**Response (201 Created):**
```json
{
  "short_url": "http://localhost:8080/Ax79bZ"
}
```

### 2. Redirect
`GET /:shortCode`

Redirects the user to the original URL (HTTP 302) and tracks the click.

### 3. Health Check
`GET /health`

**Response (200 OK):**
```json
{ "status": "ok" }
```

## 🗺️ Roadmap
- [x] Core Shortening Engine
- [x] PostgreSQL Persistence
- [x] Click Analytics
- [ ] Redis Caching Layer
- [ ] Docker & Docker Compose Support
- [ ] User Authentication & API Keys

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.