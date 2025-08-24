# Build stage
FROM golang:1.24.6-alpine3.22 AS builder

WORKDIR /app

# Install dependencies
# Copy source
COPY . .

RUN go build -o main main.go
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz

# Final image
FROM alpine:3.22
WORKDIR /app

# Copy files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# Port and entry
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
