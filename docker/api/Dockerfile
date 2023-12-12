FROM golang:1.21-alpine AS builder

COPY . /app

WORKDIR /app

RUN go mod download
RUN go build -o pack-management ./cmd/main.go


# run img
FROM alpine:latest

COPY --from=builder /app/pack-management .

CMD ["./pack-management"]
