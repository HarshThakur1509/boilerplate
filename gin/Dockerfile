
# Use an official Golang image as a builder
FROM golang:1.23.4 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application as a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .


# Use a minimal image to run the application
FROM alpine:3.21

# Install necessary runtime tools
RUN apk add --no-cache bash curl

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./main"]



