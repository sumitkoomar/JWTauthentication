# Use the official golang image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install Go modules
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
