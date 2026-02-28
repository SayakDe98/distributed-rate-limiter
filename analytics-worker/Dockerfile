# Stage 1: Build
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates gcc musl-dev

WORKDIR /app

# Copy go.mod/go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the worker binary
RUN go build -o worker ./cmd/worker

# Stage 2: Minimal final image
FROM alpine:3.18

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/worker .

# Set entrypoint
CMD ["./worker"]