# ----------------------------
# Stage 1: Build
# ----------------------------
FROM golang:1.22-alpine AS builder

# Install dependencies for module download and building
RUN apk add --no-cache git ca-certificates gcc musl-dev

WORKDIR /app

# Copy go.mod/go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the server binary
RUN go build -o server ./cmd/server

# ----------------------------
# Stage 2: Minimal final image
# ----------------------------
FROM alpine:3.18

WORKDIR /app

# Install curl for healthcheck
RUN apk add --no-cache curl

# Copy the binary from builder stage
COPY --from=builder /app/server .

# Expose port (optional, just for documentation)
EXPOSE 8080

# Start the server
CMD ["./server"]