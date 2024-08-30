# Use the official Go image as the base
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o blog cmd/blog/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application
CMD ["./blog"]
