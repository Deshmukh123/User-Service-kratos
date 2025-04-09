# FROM golang:1.21

# WORKDIR /app

# COPY . .

# RUN go mod download

# RUN go build -o user-services ./main.go

# EXPOSE 8080

# CMD ["./user-services"]
# Use official Golang image
FROM golang:1.21

# Set working directory inside the container
WORKDIR /app

# Copy Go modules and source files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the Go binary
RUN go build -o user-services ./main.go

# Expose app port
EXPOSE 8080

# Run the app
CMD ["./user-services"]
