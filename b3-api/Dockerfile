# Dockerfile for Go API application

# Start from a Go runtime base image
FROM golang:alpine3.19 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Copy .env file to the working directory (adjust the path if needed)
COPY .env ./

# Build the Go application
RUN go build -o b3-api ./cmd

# Start a new stage from a smaller base image
FROM alpine:latest

# Install necessary packages (PostgreSQL client)
RUN apk --no-cache add ca-certificates postgresql-client

# Set the current working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/b3-api .

# Copy .env file from build stage to runtime stage
COPY --from=build /app/.env ./

# Command to run the Go application
CMD ["./b3-api"]
