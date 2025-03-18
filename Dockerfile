# Stage 1: Build the Go application
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o todo-api cmd/server/main.go


# Stage 2: Create a lightweight image for running the application
FROM alpine:latest

# Install necessary dependencies (if required)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/todo-api .

# Copy the .env file (optional, only if needed at build time)
COPY .env .env

# Expose API port
EXPOSE 8080

# Run the application
CMD ["./todo-api"]
