# 1. Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /product

# Gerekli bağımlılıkları yükle
COPY go.mod go.sum ./
RUN go mod download

# Kaynak kodları kopyala ve build et
COPY . .
RUN go build -o main ./cmd/server

# 2. Run stage
FROM alpine:latest

# Sertifikalar için ca-certificates paketi gerekli
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /product/main .

CMD ["./main"]
