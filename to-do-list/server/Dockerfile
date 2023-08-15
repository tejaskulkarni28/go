# Use the official Go image for ARM64 architecture (M1)
FROM golang:1.20.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the workspace
COPY server/go.mod ./
COPY server/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY server/ ./

# Build the Go application
RUN go build -o app

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./app"]
