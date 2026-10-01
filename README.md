# FiMuVer - Medienverwaltungs-Software

FiMuVer ist eine Full-Stack-Anwendung zum Verwalten einer Mediensammlung (Blurays, DVDs, Vinyl, Tapes) mit Benutzerkonten, mehreren Sammlungen pro Nutzer, Invite-basierter Registrierung und TVDB-Anbindung für Film-/Serien-Metadaten.

## 🏗️ Projektstruktur

```
FiMuVer/
├── backend/                   # Go API Backend
│   ├── cmd/api/               # Einstiegspunkt (main.go, Router-Setup)
│   ├── internal/
│   │   ├── auth/               # JWT Erzeugung/Validierung
│   │   ├── config/              # YAML + ENV Konfiguration
│   │   ├── db/                  # DB-Verbindung, AutoMigration, Seeding
│   │   ├── models/              # GORM Datenmodelle
│   │   ├── handlers/            # HTTP Handler (Request/Response-DTOs)
│   │   ├── services/            # Business-Logik / DB-Zugriff
│   │   └── middleware/          # CORS, JWT-Auth
│   ├── go.mod
│   ├── config.yaml             # Server/DB/JWT/TVDB Konfiguration
│   └── Dockerfile
├── frontend/                  # React (Vite) Frontend
│   ├── src/
│   │   ├── components/          # Header, FilterBar, MediaCard/Form, InviteCodes, ...
│   │   ├── pages/                # Landing, Auth, Admin, CollectionPage, ItemDetailPage
│   │   ├── services/             # API Clients (apiClient, userapi, collection, editionapi, ...)
│   │   ├── hooks/                 # Custom React Hooks
│   │   ├── types/                 # Konstanten (Media-Typen etc.)
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── package.json
│   └── vite.config.js
├── docker-compose.yml          # Docker Orchestration
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

# .env im Projekt-Root anlegen (siehe .env.example) und config.yaml prüfen

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

Das Frontend läuft dann unter `http://localhost:5173`.

## 📋 Konfiguration

### Backend - config.yaml

Die Datei `backend/config.yaml` definiert Server-, Datenbank-, JWT- und TVDB-Einstellungen:

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

jwt:
  secret: "changeme"
  ttl: "24h"

tvdb:
  tvdb_api_key: ""
  base_url: "https://api4.thetvdb.com/v4"
```

Alle Werte können per Umgebungsvariable überschrieben werden: `DATABASE_HOST`, `DATABASE_PORT`, `DATABASE_USER`, `DATABASE_PASSWORD`, `DATABASE_NAME`, `JWT_SECRET`, `JWT_TTL`, `TVDB_API_KEY`, `TVDB_BASE_URL`. Diese werden zusätzlich aus einer `.env`-Datei im Projekt-Root geladen (`godotenv`).

**Bei Docker Compose:** `config.yaml` wird automatisch mit den Docker-Umgebungsvariablen gefüllt.

## 🗄️ Datenbank

PostgreSQL, verwaltet per GORM AutoMigration (Tabellen werden beim Start automatisch erstellt/aktualisiert). Details zum Schema siehe [`docs/DATA-MODEL.md`](docs/DATA-MODEL.md).

Kurzüberblick der wichtigsten Tabellen:
- **users** – Benutzerkonten (Email, Username, bcrypt-Passwort-Hash, `is_admin`)
- **collections** – Sammlungen eines Users (z. B. "Meine Blurays")
- **items** – Medieneinträge innerhalb einer Collection (Titel, Typ, Jahr, Genre, Zustand, Lagerort, optionale TVDB-ID/Cover)
- **invite_codes** – Einladungscodes für die Registrierung (mit Nutzungslimit/Ablaufdatum)
- **settings** – globale Boolean-Settings (z. B. `enable_registration`)
- **movies, persons, genres, editions, labels, media_types, conditions, movie_actors, movie_genres, collection_items** – normalisiertes Metadaten-Schema für Filme/Editionen; aktuell teilweise vorbereitet, aber noch nicht vollständig über die API verdrahtet (nur `editions` hat einen Read-Endpoint)

## 🔌 API Endpoints

Basis-URL: `http://localhost:8080/api/v1`. Endpunkte außer Registrierung/Login erfordern einen `Authorization: Bearer <token>` Header (JWT).

### Auth / User
```
POST /users            Registrierung (alias: POST /users/register)
POST /users/login      Login → { token, id, username, email, is_admin }
GET  /users/:id        Benutzer per ID abrufen (auth erforderlich)
```

### Collections
```
GET    /collections              Alle Collections des eingeloggten Users
POST   /collections               Neue Collection anlegen
GET    /collections/:id           Eine Collection abrufen (nur eigene)
PUT    /collections/:id           Collection aktualisieren (nur eigene)
DELETE /collections/:id           Collection löschen (nur eigene)
```

### Items (innerhalb einer Collection)
```
POST   /collections/:id/items             Item zur Collection hinzufügen (nur eigene Collection)
DELETE /collections/:id/items/:itemId     Item löschen (nur aus eigener Collection)
```

### Invite Codes
```
POST   /invite/generate   Neuen Invite-Code erzeugen
GET    /invite/list        Alle Invite-Codes auflisten
DELETE /invite/:id          Invite-Code löschen
```

### Settings
```
GET    /settings           Alle Settings
GET    /settings/:name      Ein Setting nach Name
PUT    /settings/:name      Setting aktualisieren
DELETE /settings/:id         Setting löschen
```

### TVDB (Metadaten-Suche)
```
GET /tvdb/search/series?q=...
GET /tvdb/search/movies?q=...
```

### Editionen
```
GET /editions               Alle Editionen (z. B. Steelbook, Limited, Standard)
```

### Sonstiges
```
GET /health                 Health Check
```

## 🎨 Frontend Features

- **Landing-Seite:** Übersicht der eigenen Collections
- **Auth:** Login/Registrierung, optional mit Invite-Code
- **Collections:** Anlegen, Umbenennen, Löschen; einzelne Items hinzufügen/löschen
- **Item-Detailseite:** Detailansicht eines Medieneintrags
- **Admin-Panel:** (nur für `is_admin`-User) Verwaltung von Invite-Codes und Settings
- **TVDB-Suche:** Anbindung zur Metadaten-Suche für Filme/Serien
- **Responsive Design:** Funktioniert auf Desktop und Mobile

## 🛠️ Entwicklung

### Backend Development

```bash
cd backend

# Server mit Hot Reload starten (mit Air)
go install github.com/cosmtrek/air@latest
air

# Tests ausführen
go test ./...
```

### Frontend Development

```bash
cd frontend

npm run dev        # Dev Server mit HMR
npm run build       # Production Build
npm run preview      # Preview Production Build
npm run lint          # ESLint
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
- **gorm** + **gorm/driver/postgres** - ORM & PostgreSQL Driver
- **golang-jwt** - JWT Erzeugung/Validierung
- **golang.org/x/crypto/bcrypt** - Passwort-Hashing
- **joho/godotenv** - `.env`-Support
- **gopkg.in/yaml.v2** - YAML Parsing

### Frontend (React)
- **react** / **react-dom** (v19)
- **react-icons**
- **vite** - Build Tool

## 🔐 Sicherheit — aktueller Stand & offene Punkte

Bereits umgesetzt:
- [x] JWT Authentication + bcrypt Passwort-Hashing
- [x] Ownership-Checks bei Collections/Items (nur eigene Daten abrufbar/änderbar)
- [x] CORS Middleware
- [x] Invite-Code-gesteuerte Registrierung

Noch offen / Roadmap:
- [ ] Rate Limiting (v. a. Login/Registrierung)
- [ ] Umfassendere Input-Validierung
- [ ] Tests für Handler/Services (aktuell nur CORS-Middleware getestet)
- [ ] Verschlüsselung sensibler Zusatzdaten

## 📝 Nächste Schritte

1. **Normalisiertes Metadaten-Schema fertig verdrahten:** Movie/Person/Genre/CollectionItem stehen als Models bereit, aber nur `editions` ist über die API erreichbar
2. **Item-Update-Endpoint:** Aktuell nur Erstellen/Löschen, kein `PUT` für Items
3. **Pagination:** Für Collections/Items bei großen Sammlungen
4. **Export:** PDF/CSV Export der Mediensammlung
5. **Statistiken:** Dashboard mit Statistiken (Medienanzahl, Genre-Verteilung, etc.)
6. **Cover-Upload:** Direkter Bild-Upload statt nur `image_url`-Feld
7. **Tests:** Unit- und Integrationstests für Backend und Frontend

## 📄 Lizenz

Dieses Projekt ist Open Source - frei zu verwenden und zu modifizieren.

## 👤 Author

Felix - FiMuVer Projekt (2026)

---

**Viel Spaß beim Verwalten deiner Mediensammlung!** 🎬🎵📀
