# Build stage
FROM golang:1.18.9-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8000
CMD ["/app/main"]
