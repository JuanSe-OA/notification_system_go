# Etapa de build
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Dependencias
COPY go.mod go.sum ./
RUN go mod download

# Código
COPY . .

# Compilamos el binario principal (usa cmd/main.go)
RUN go build -o notification-service ./cmd

# Etapa de runtime
FROM alpine:3.20

WORKDIR /app

# Usuario no root (opcional)
RUN adduser -D appuser
USER appuser

# Puerto por defecto (lo sobreescribimos con env SERVER_PORT)
ENV SERVER_PORT=8080

COPY --from=builder /app/notification-service .

EXPOSE 8080

CMD ["./notification-service"]
