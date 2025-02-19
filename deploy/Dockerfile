# ⚡ Step 1: Use a lightweight Go base image
FROM golang:1.19-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# ⚡ Step 2: Set working directory inside container
WORKDIR /app

# ⚡ Step 3: Copy Go modules first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# ⚡ Step 4: Copy the entire source code
COPY . .

# ⚡ Step 5: Build the application
RUN go build -o taskwise cmd/main.go

# ⚡ Step 6: Create a small image with only the built binary
FROM alpine:latest

WORKDIR /root/

# ⚡ Step 7: Copy binary from builder
COPY --from=builder /app/taskwise .

# ⚡ Step 8: Copy .env file (if using in production, use secrets instead)
COPY .env .env

# ⚡ Step 9: Expose port and run the application
EXPOSE 8080
CMD ["./taskwise"]
