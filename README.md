# URL Shortener API

A production-grade URL shortening service built with Go.

## Features

- Shorten long URLs to 6-character codes
- Redirect users to the original URL
- Concurrent-safe in-memory storage
- Full test coverage
- RESTful API design

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── shortener.go     # HTTP handlers
│   │   └── shortener_test.go
│   └── storage/
│       ├── memory.go        # In-memory storage
│       └── memory_test.go
├── Makefile                 # Common commands
├── go.mod
└── README.md
```

## API Endpoints

### Create a Short URL

```
POST /shorten
Content-Type: application/json

{
    "url": "https://example.com"
}
```

**Response:**

```
201 Created

{
    "short_code": "aB3xYz",
    "short_url": "http://localhost:8080/aB3xYz"
}
```

### Redirect to Original URL

```
GET /{shortCode}
```

Returns a 302 redirect to the original URL.

## Error Responses

| Status Code | Description |
|-------------|-------------|
| 400 | Invalid URL or JSON body |
| 404 | Short code not found |
| 405 | Method not allowed |
| 409 | Short code already exists (collision) |

## Running the Application

### Development

```bash
# Run the server
make run

# Run tests
make test

# Build the binary
make build
```

### Manual

```bash
# Install dependencies
go mod tidy

# Run the server
go run cmd/api/main.go
```

## Testing

```bash
# Run all tests
go test ./... -v

# Run tests with coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```