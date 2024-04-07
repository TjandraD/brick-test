# Use the official Golang image as base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o money_transfer.backend .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./money_transfer.backend"]
