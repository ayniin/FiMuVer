# Quick Start Guide - FiMuVer

Willkommen! Hier ist eine schnelle Anleitung zum Starten der Anwendung.

## 🚀 Schnelleste Methode: Docker Compose

### Voraussetzung
- Docker und Docker Compose installiert

### Kommandos

```bash
# 1. Im Projekt-Root-Verzeichnis
cd /home/felix/Projects/FiMuVer

# 2. Starte alle Services
docker-compose up -d

# 3. Fertig! Öffne im Browser:
# Backend API: http://localhost:8080/health
# Frontend: http://localhost:5173
# pgAdmin: http://localhost:5050 (optional, für DB-Verwaltung)
```

**Logs anschauen:**
```bash
docker-compose logs -f backend
docker-compose logs -f postgres
```

**Services stoppen:**
```bash
docker-compose down
```

---

## 💻 Lokale Entwicklung (ohne Docker)

### Voraussetzungen
- Go 1.25+
- Node.js 18+
- PostgreSQL 16+ (lokal oder Docker-Container)

### Backend starten

```bash
# 1. Stelle sicher, dass PostgreSQL läuft
# Z.B. in einem separaten Terminal:
docker run --name fimuver-db -e POSTGRES_USER=fimuver_user \
  -e POSTGRES_PASSWORD=fimuver_password \
  -e POSTGRES_DB=fimuver_db \
  -p 5432:5432 postgres:16-alpine

# 2. Backend Dependencies installieren
cd backend
go mod tidy

# 3. Backend starten
go run ./cmd/api/main.go
# Backend läuft auf: http://localhost:8080
```

### Frontend starten (neues Terminal)

```bash
cd frontend
npm install
npm run dev
# Frontend läuft auf: http://localhost:5173
```

---

## 🔧 Mit Makefile (empfohlen)

Wenn `make` installiert ist, kannst du einfache Kommandos verwenden:

```bash
# Alle verfügbaren Kommandos anschauen
make help

# Backend starten
make dev-backend

# Frontend starten
make dev-frontend

# Docker starten
make docker-up

# Logs anschauen
make docker-logs
```

---

## 📝 Mit Start-Script

```bash
./start.sh
```

Das Script bietet interaktive Optionen zum Starten.

---

## 🎯 Test-Daten hinzufügen

Nach dem Start kannst du über das UI oder curl neue Medien hinzufügen:

```bash
# Beispiel: Ein Bluray hinzufügen
curl -X POST http://localhost:8080/api/v1/media \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Matrix",
    "media_type": "bluray",
    "director": "Wachowski",
    "year": 1999,
    "genre": "Science Fiction",
    "condition": "good",
    "location": "Regal 1"
  }'
```

---

## 🆘 Häufige Probleme

### Port bereits in Verwendung

**Problem:** "address already in use :8080"

**Lösung:**
```bash
# Finde den Prozess auf Port 8080
lsof -i :8080

# Oder mit Docker:
docker-compose down
```

### Datenbank-Verbindungsfehler

**Problem:** "cannot connect to database"

**Lösung:**
1. Stelle sicher, dass PostgreSQL läuft
2. Überprüfe die config.yaml im backend/ Verzeichnis
3. Überprüfe die Datenbank-Credentials

### Node Dependencies fehlen

```bash
cd frontend
npm install
npm cache clean --force
npm install
```

---

## 📚 Weitere Ressourcen

- Vollständiges README: `README.md`
- API-Dokumentation: Im README unter "API Endpoints"
- Docker-Kommandos: Im README unter "🐳 Docker Kommandos"

---

## 🎉 Du bist bereit!

Die Anwendung sollte jetzt laufen. Viel Spaß mit FiMuVer! 🎬🎵📀

Bei Fragen oder Problemen: README.md oder Code-Kommentare durchschauen.

