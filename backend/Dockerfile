# Stage 1: Builder
FROM golang:1.25-alpine AS builder

# Setze das Arbeitsverzeichnis im Container
WORKDIR /app

# Kopiere go.mod und go.sum, um die Abhängigkeiten zu cachen
COPY go.mod go.sum ./

# Lade die Abhängigkeiten herunter
RUN go mod download

# Kopiere den gesamten Quellcode des Backends
COPY . .

# Baue die Go-Anwendung. CGO_ENABLED=0 erstellt ein statisches Binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api

# Stage 2: Finales Image
FROM alpine:latest

# Füge ca-certificates hinzu, um HTTPS-Verbindungen zu ermöglichen
RUN apk --no-cache add ca-certificates

# Setze das Arbeitsverzeichnis
WORKDIR /root/

# Kopiere das kompilierte Binary aus dem Builder-Stage
COPY --from=builder /app/api .

# Kopiere die Konfigurationsdatei
COPY config.yaml .

# Expose Port 8080, auf dem die Anwendung läuft
EXPOSE 8080

# Definiere den Befehl zum Starten der Anwendung
CMD ["./api"]

