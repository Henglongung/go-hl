# Use the official Go image as a build stage
FROM golang:1.20-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o go-hl

# Use a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=builder /app/go-hl .

# Command to run the executable
CMD ["./go-hl"]
