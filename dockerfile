FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o poc-sender cmd/main.go

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/poc-sender .

ENV PORT=8080

EXPOSE 8080

CMD ["./poc-sender"]