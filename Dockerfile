# ---------- BUILD STAGE ----------
# Use official Go image to compile the application
FROM golang:1.25 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy project source code into container
COPY . .

# Compile Go binary for Linux with CGO disabled (static binary)
# -o app : output binary named "app"
# cmd/server/main.go : entry point of the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o app cmd/server/main.go


# ---------- PRODUCTION STAGE ----------
# Use lightweight Alpine image for final container
FROM alpine:latest AS production

# Set working directory inside production container
WORKDIR /app

# Copy only compiled binary from builder stage
COPY --from=builder /app/app .

COPY --from=builder /app/migrations ./migrations

COPY --from=builder /app/migrations /migrations

# Start the application when container runs
CMD ["./app"]
