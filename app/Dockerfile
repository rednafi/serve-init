# Build stage
FROM golang:1.22-alpine AS builder

# Install git (necessary to download dependencies)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY app/go.mod app/go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Run stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
