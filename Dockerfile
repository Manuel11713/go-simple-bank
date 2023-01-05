# Build stage
FROM golang:1.18.9-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder  /app/migrate ./migrate
COPY app.env .
COPY db/migration ./migration
COPY start.sh .

EXPOSE 8000
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
