# Use official Golang image
FROM golang:1.21

# Set working directory inside the container
WORKDIR /app

# Copy Go modules and source files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y curl

# Build the Go binary
RUN go build -o user-service ./main.go

# Expose app port
EXPOSE 8081

# Run the app
CMD ["./user-service"]
