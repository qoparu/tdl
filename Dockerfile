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

COPY --from=builder /server /server
COPY wait-for-it.sh /wait-for-it.sh
COPY docker-entrypoint.sh /docker-entrypoint.sh

RUN chmod +x /wait-for-it.sh /docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/docker-entrypoint.sh"]
