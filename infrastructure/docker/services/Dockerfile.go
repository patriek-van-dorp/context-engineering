# Go/Gin Multi-stage Production Dockerfile
# Optimized for size, security, and performance

# ================================
# Build Stage
# ================================
FROM golang:1.21-alpine as builder

# Security: Install ca-certificates for HTTPS requests
RUN apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY src/ ./src/
COPY *.go ./

# Build binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -o app .

# ================================
# Production Stage
# ================================
FROM scratch as production

# Copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary from builder
COPY --from=builder /build/app /app

# Set environment variables
ENV PORT=8080 \
    GIN_MODE=release

# Expose port
EXPOSE 8080

# Health check (note: scratch doesn't have curl, so this needs to be handled by orchestrator)
# HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
#     CMD ["/app", "-health-check"] || exit 1

# Run application
ENTRYPOINT ["/app"]
