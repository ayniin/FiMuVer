# Multi-Stage Build für Go Backend
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Kopiere Go Module Files
COPY backend/go.mod backend/go.sum ./

# Lade Dependencies
RUN go mod download

# Kopiere Source Code
COPY backend/ .

# Baue die Anwendung
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Final Stage
FROM alpine:latest

# Installiere ca-certificates für HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Kopiere die compilierte Anwendung aus dem Builder
COPY --from=builder /app/api .
COPY backend/config.yaml .

# Expose Port
EXPOSE 8080

# Starte die Anwendung
CMD ["./api"]

