# Go URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A production-grade URL Shortener backend built with Go and Gin. Designed with clean architecture principles to ensure scalability, maintainability, and testability.

## 🚀 Features

- **High Performance**: Built on top of Gin, one of the fastest HTTP web frameworks for Go.
- **Clean Architecture**: Separation of concerns with Handler, Service, Repository, and Model layers.
- **Configuration Management**: flexible environment-based configuration.
- **Health Checks**: Built-in health monitoring endpoints.
- **Scalable**: Ready for PostgreSQL integration (upcoming).

## 🛠️ Tech Stack

- **Language**: [Go](https://golang.org/)
- **Framework**: [Gin Web Framework](https://gin-gonic.com/)
- **Database**: PostgreSQL (Planned)
- **Configuration**: Environment Variables

## 📂 Project Structure

```bash
.
├── cmd
│   └── server          # Application entry point
├── internal
│   ├── config          # Configuration management
│   ├── handler         # HTTP handlers (Controllers)
│   ├── model           # Domain models
│   ├── repository      # Data access layer
│   └── service         # Business logic layer
└── go.mod              # Go module definition
```

## 🏁 Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

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

3.  **Run the application**
    ```bash
    go run cmd/server/main.go
    ```

    The server will start on port `8080`.

## 🔌 API Documentation

### Health Check

**Endpoint**
`GET /health`

**Response**
```json
{
  "status": "ok"
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.