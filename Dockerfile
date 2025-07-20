# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary from your cmd/server package
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Copy the built binary from the 'builder' stage
COPY --from=builder /server /server

# Expose the port your app runs on
EXPOSE 8080

# The command to run when the container starts
ENTRYPOINT ["/server"]