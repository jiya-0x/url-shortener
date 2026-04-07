.PHONY: run build test clean

# Run the server
run:
	go run cmd/api/main.go

# Build the binary
build:
	go build -o bin/url-shortener cmd/api/main.go

# Run all tests
test:
	go test ./... -v

# Clean the binary
clean:
	rm -rf bin/

# Format code
fmt:
	go fmt ./...

# Tidy dependencies
tidy:
	go mod tidy

# Run tests with coverage
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out