# Go/Gin Dockerfile for Agents
# Inherits from base service Dockerfile with agent-specific optimizations

# ================================
# Build Stage
# ================================
FROM golang:1.21-alpine as builder

# Install build dependencies
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata

# Set working directory
WORKDIR /app

# Copy go mod files for dependency download
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with optimizations for agent workloads
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o /bin/agent-app \
    ./src/main/agents/${AGENT_TYPE}/

# ================================
# Production Stage
# ================================
FROM alpine:3.18 as production

# Install runtime dependencies
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    curl

# Security: Create non-root user
RUN addgroup -g 1000 -S agent && \
    adduser -u 1000 -S agent -G agent

# Set environment variables for production
ENV GIN_MODE=release \
    PORT=8080 \
    AGENT_TYPE="" \
    LOG_LEVEL=info

# Set working directory
WORKDIR /app

# Copy CA certificates for HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary from builder stage
COPY --from=builder /bin/agent-app /bin/agent-app

# Copy shared configuration and templates
COPY --chown=agent:agent src/main/shared/ ./src/main/shared/

# Copy agent-specific configurations
COPY --chown=agent:agent src/main/agents/${AGENT_TYPE}/config/ ./config/

# Switch to non-root user
USER agent

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Run application
CMD ["/bin/agent-app"]
