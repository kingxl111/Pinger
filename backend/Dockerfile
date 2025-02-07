# Используем образ Golang для сборки приложения
FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o backend .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]
