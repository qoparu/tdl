# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

# Stage 2: Create the final lightweight image
FROM alpine:latest
WORKDIR /app

# Copy the built binary and wait-for-it.sh from builder stage
COPY --from=builder /server /server
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Expose the port your app runs on
EXPOSE 8080

# Use wait-for-it to wait for the db before starting the server
ENTRYPOINT ["/wait-for-it.sh", "db:5432", "--", "/server"]
