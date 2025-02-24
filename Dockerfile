FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

# ВАЖНО: указываем, что главный go-файл находится в cmd/main.go
RUN swag init -g cmd/main.go

RUN go build -o messenger ./cmd/main.go

# Финальный образ:
FROM debian:stable-slim

WORKDIR /app
COPY --from=builder /app/messenger /app/messenger
COPY --from=builder /app/web /app/web
COPY --from=builder /app/docs /app/docs

EXPOSE 8080

ENTRYPOINT ["/app/messenger"]
