FROM golang:1.21-alpine AS builder

COPY . /app

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["goose", "-dir", "migrations", "postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable", "up"]