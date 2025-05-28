# Build stage: build the Go binary
FROM golang:1.23-bullseye AS builder

WORKDIR /app

# Copy go.mod and go.sum and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build the binary, output to /app/golden-trail
RUN CGO_ENABLED=0 GOOS=linux go build -o golden-trail ./cmd/main.go

# Final stage: minimal image
FROM alpine:latest

# Install CA certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/golden-trail .

# Expose API port
EXPOSE 5000

# Run the binary
CMD ["./golden-trail"]
