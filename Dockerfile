# Build stage
FROM golang:1.24.6-alpine3.22 AS builder

WORKDIR /app

# Install dependencies
# Copy source
COPY . .

RUN go build -o main main.go

# Final image
FROM alpine:3.22
WORKDIR /app

# Copy files from builder
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

# Port and entry
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
