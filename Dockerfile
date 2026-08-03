# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install ca-certificates and git
RUN apk add --no-cache ca-certificates git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build API binary and migrations CLI binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/migrate_bin ./migrations

# Production runtime stage
FROM alpine:3.20

LABEL org.opencontainers.image.source=https://github.com/AppeiYA/hotel_system2
LABEL org.opencontainers.image.description="Go Hotel Management System API"

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy binaries from build stage
COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/bin/migrate_bin /app/migrate_bin

# Copy migrations and seeds SQL files
COPY --from=builder /app/migrations/*.sql /app/migrations/
COPY --from=builder /app/seeds /app/seeds

# Copy and setup entrypoint script
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

EXPOSE 3333

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/api"]
