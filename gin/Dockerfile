# Use an official Golang image as a builder
FROM golang:1.23.5 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and entrypoint script
COPY . .

# Build the main application and migration tool
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main . && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migrate migrate/migrate.go

# Use a minimal image for the final stage
FROM alpine:3.21

WORKDIR /app

# Copy the compiled binaries from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Make files executable
RUN chmod +x ./main ./migrate

EXPOSE 3000

# Run migrations and start the application
CMD ./migrate && ./main