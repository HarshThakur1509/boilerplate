# Use an official Golang image as a builder
FROM golang:1.23.5 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code 
COPY . .

# Build the main application and migration tool
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api cmd/api/main.go && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migrations cmd/migrations/main.go

# Use a minimal image for the final stage
FROM alpine:3.21

WORKDIR /app

# Copy the compiled binaries from builder
COPY --from=builder /app/api .
COPY --from=builder /app/migrations .

# Make files executable
RUN chmod +x ./api ./migrations

EXPOSE 3000

# Run migrations and start the application
CMD ["sh", "-c", "./migrations && ./api"]