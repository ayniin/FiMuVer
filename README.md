# FiMuVer - Medienverwaltungs-Software

FiMuVer ist eine Full-Stack-Anwendung zum Verwalten deiner Mediensammlung (Blurays, DVDs, Vinyl CDs und Tapes).

## 🏗️ Projektstruktur

```
FiMuVer/
├── backend/                   # Go API Backend
│   ├── cmd/api/              # Einstiegspunkt
│   │   └── main.go
│   ├── internal/
│   │   ├── config/           # YAML Konfiguration
│   │   ├── db/               # Datenbankverbindung (GORM)
│   │   ├── models/           # Datenmodelle
│   │   ├── handlers/         # API Endpoints
│   │   └── middleware/       # CORS, Auth, etc.
│   ├── migrations/           # Datenbank Migrationen
│   ├── go.mod               # Go Dependencies
│   ├── config.yaml          # Datenbank Konfiguration
│   └── Dockerfile           # Backend Container
├── frontend/                 # React Vite Frontend
│   ├── src/
│   │   ├── components/      # Reusable UI Komponenten
│   │   ├── pages/           # Seiten (Dashboard)
│   │   ├── services/        # API Clients
│   │   ├── hooks/           # Custom React Hooks
│   │   ├── types/           # TypeScript/JS Types
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── package.json
│   └── vite.config.js
├── docker-compose.yml       # Docker Orchestration
└── README.md
```

## 🚀 Schnellstart

### Voraussetzungen
- Docker und Docker Compose
- Oder lokal: Go 1.25+, Node.js 18+, PostgreSQL 16+

### Mit Docker Compose (empfohlen)

```bash
# Im Root-Verzeichnis
docker-compose up -d

# Backend läuft auf: http://localhost:8080
# Frontend: http://localhost:5173
# pgAdmin: http://localhost:5050
```

### Lokal ohne Docker

**Backend:**
```bash
cd backend

# Installiere Go Dependencies
go mod tidy

# Starte PostgreSQL (z.B. lokal oder in separatem Container)
# Stelle sicher, dass config.yaml korrekt konfiguriert ist

# Starte den Server
go run ./cmd/api/main.go
```

**Frontend:**
```bash
cd frontend

# Installiere Dependencies
npm install

# Starte Dev Server
npm run dev
```

Der Frontend lädt dann unter `http://localhost:5173`

## 📋 Konfiguration

### Backend - config.yaml

Die Datei `backend/config.yaml` definiert die Datenbank-Verbindung:

```yaml
server:
  host: "0.0.0.0"
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "fimuver_user"
  password: "fimuver_password"
  database: "fimuver_db"
  sslmode: "disable"
```

**Bei Docker Compose:** Die `config.yaml` wird automatisch mit den Docker-Umgebungsvariablen gefüllt.

**Für lokale Entwicklung:** Stelle sicher, dass PostgreSQL läuft und die Credentials korrekt sind.

## 🗄️ Datenbank

### Schema

**Media Tabelle:**
- `id` (Primary Key)
- `title` - Titel des Mediums
- `description` - Beschreibung
- `media_type` - bluray, dvd, vinyl oder tape
- `artist` - Künstler (für Musik)
- `director` - Regisseur (für Filme)
- `year` - Erscheinungsjahr
- `genre` - Genre
- `condition` - Zustand (mint, good, fair, poor)
- `location` - Lagerort
- `notes` - JSON Feld für zusätzliche Metadaten
- `created_at`, `updated_at` - Zeitstempel

Die Tabelle wird automatisch bei der ersten Verbindung erstellt (GORM AutoMigration).

## 🔌 API Endpoints

### Media Management

**Alle Medien abrufen (optional gefiltert):**
```
GET /api/v1/media?type=bluray
```

**Ein Medium abrufen:**
```
GET /api/v1/media/:id
```

**Medium erstellen:**
```
POST /api/v1/media
Content-Type: application/json

{
  "title": "The Matrix",
  "media_type": "bluray",
  "director": "Wachowski",
  "year": 1999,
  "genre": "Science Fiction",
  "condition": "good",
  "location": "Regal 1"
}
```

**Medium aktualisieren:**
```
PUT /api/v1/media/:id
```

**Medium löschen:**
```
DELETE /api/v1/media/:id
```

**Nach Medien suchen:**
```
GET /api/v1/search?q=Matrix
```

**Health Check:**
```
GET /health
```

## 🎨 Frontend Features

- **Dashboard:** Übersicht aller Medien
- **Filter:** Nach Medientyp filtern
- **Suche:** Nach Titel, Künstler oder Regisseur suchen
- **Formular:** Neue Medien hinzufügen oder bearbeiten
- **CRUD:** Vollständiges Löschen, Bearbeiten, Anzeigen
- **Responsive Design:** Funktioniert auf Desktop und Mobile

## 🛠️ Entwicklung

### Backend Development

```bash
cd backend

# Dependencies installieren
go get github.com/gin-gonic/gin
go get github.com/gorm.io/gorm

# Server mit Hot Reload starten (mit Air)
go install github.com/cosmtrek/air@latest
air

# Tests ausführen
go test ./...
```

### Frontend Development

```bash
cd frontend

# Dependencies installieren
npm install

# Dev Server mit Hot Module Replacement
npm run dev

# Production Build
npm run build

# Preview Production Build
npm run preview
```

## 🐳 Docker Kommandos

```bash
# Alle Services starten
docker-compose up -d

# Logs ansehen
docker-compose logs -f backend

# Services stoppen
docker-compose down

# Datenbank neu initialisieren
docker-compose down -v
docker-compose up -d

# In Container einloggen
docker-compose exec backend sh
docker-compose exec postgres psql -U fimuver_user -d fimuver_db
```

## 📦 Dependencies

### Backend (Go)
- **gin** - HTTP Framework
- **gorm** - ORM für Datenbankoperationen
- **gorm/driver/postgres** - PostgreSQL Driver
- **yaml** - YAML Parsing

### Frontend (React)
- **react** - UI Framework
- **vite** - Build Tool

## 🔐 Sicherheit (Roadmap)

Folgende Funktionen sollten später implementiert werden:
- [ ] JWT Authentication
- [ ] Input Validation
- [ ] Rate Limiting
- [ ] Encryption für sensitive Daten
- [ ] CORS Konfiguration
- [ ] Environment Variables Management

## 📝 Nächste Schritte

1. **Authentifizierung:** JWT oder Session-basierte Auth implementieren
2. **Persistierung:** Lokales Frontend Caching mit localStorage/IndexedDB
3. **Export:** PDF/CSV Export der Mediensammlung
4. **Statistiken:** Dashboard mit Statistiken (Medienanzahl, Genre-Verteilung, etc.)
5. **Uploads:** Bilder/Cover hochladen
6. **Tests:** Unit und Integration Tests
7. **Deployment:** Docker Hub, Kubernetes, etc.

## 📄 Lizenz

Dieses Projekt ist Open Source - frei zu verwenden und zu modifizieren.

## 👤 Author

Felix - FiMuVer Projekt (2026)

---

**Viel Spaß beim Verwalten deiner Mediensammlung!** 🎬🎵📀
