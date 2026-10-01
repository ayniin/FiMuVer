# 📁 FiMuVer - Projektstruktur und Architektur

## Überblick

FiMuVer ist eine **Full-Stack-Webanwendung** für die Verwaltung von Mediensammlungen mit Benutzerkonten und Invite-basierter Registrierung.

- **Backend:** Go mit Gin Framework, GORM ORM, JWT-Auth
- **Frontend:** React (Vite)
- **Datenbank:** PostgreSQL
- **Containerisierung:** Docker & Docker Compose

---

## 📂 Verzeichnisstruktur

```
FiMuVer/
│
├── 📁 backend/                         # Go REST API
│   ├── cmd/
│   │   └── api/
│   │       └── main.go                 # Einstiegspunkt, Router-Setup (public + JWT-secured Gruppen)
│   │
│   ├── internal/                       # Private Pakete
│   │   ├── auth/
│   │   │   └── jwt.go                  # Token-Erzeugung/-Validierung (GenerateToken, ParseToken)
│   │   │
│   │   ├── config/
│   │   │   └── config.go               # YAML + ENV Konfiguration (Server, DB, JWT, TVDB)
│   │   │
│   │   ├── models/                     # GORM Datenmodelle (siehe docs/DATA-MODEL.md)
│   │   │   ├── user.go, collection.go, item.go, media.go
│   │   │   ├── invite_code.go, settings.go
│   │   │   └── movie.go, person.go, reference_types.go, relationships.go, collection_item.go
│   │   │
│   │   ├── db/
│   │   │   └── database.go             # DB-Verbindung, AutoMigration, Default-Settings/Edition-Seeding
│   │   │
│   │   ├── handlers/                   # HTTP Handler: Request-Binding, DTOs, Responses
│   │   │   ├── user.go                 # Register/Login/GetByID
│   │   │   ├── collection.go           # Collection CRUD (mit Ownership-Checks)
│   │   │   ├── item.go                 # Item Add/Delete (mit Ownership-Checks über Collection)
│   │   │   ├── invite.go               # Invite-Code Generieren/Auflisten/Löschen
│   │   │   ├── settings.go             # Settings CRUD
│   │   │   ├── edition.go              # Editionen (read-only)
│   │   │   ├── tvdb.go                 # TVDB Proxy-Suche
│   │   │   ├── respond.go              # gemeinsame Response-Helper (ok/created/fail/...)
│   │   │   └── messages.go             # zentrale Fehler-/Erfolgsmeldungen
│   │   │
│   │   ├── services/                   # Business-Logik / DB-Zugriff, getrennt von Handlern
│   │   │   ├── user.go                 # Passwort-Hashing (bcrypt), Login, User-Lookups
│   │   │   ├── collection.go           # Collection CRUD inkl. Ownership-Filterung
│   │   │   ├── invite.go               # Invite-Code Validierung/Verbrauch
│   │   │   ├── edition.go, movie.go
│   │   │   ├── reference_service.go    # Genres/Labels/Conditions/MediaTypes
│   │   │   └── tvdb.go                 # TVDB API Client
│   │   │
│   │   └── middleware/
│   │       ├── cors.go                 # CORS Middleware für Frontend
│   │       ├── jwt.go                  # JWTAuthMiddleware, setzt user_id im Context
│   │       └── cors_test.go
│   │
│   ├── config.yaml                     # Server/DB/JWT/TVDB-Konfiguration
│   ├── go.mod / go.sum
│   └── Dockerfile                      # Multi-Stage Docker Build
│
├── 📁 frontend/                        # React Vite Projekt
│   ├── src/
│   │   ├── components/
│   │   │   ├── Header.jsx / .css
│   │   │   ├── MediaCard.jsx / .css        # Anzeige einzelner Medien
│   │   │   ├── MediaForm.jsx / .css        # Formular zum Erstellen/Bearbeiten
│   │   │   ├── FilterBar.jsx / .css        # Filter & Suche
│   │   │   ├── CreateCollectionModal.jsx / .css
│   │   │   └── InviteCodes.jsx / .css      # Invite-Code Verwaltung (Admin)
│   │   │
│   │   ├── pages/
│   │   │   ├── Landing.jsx / .css      # Übersicht der eigenen Collections
│   │   │   ├── Auth.jsx / .css         # Login/Registrierung
│   │   │   ├── Admin.jsx / .css        # Admin-Panel (Invite-Codes, Settings)
│   │   │   ├── CollectionPage.jsx      # Items einer Collection
│   │   │   └── ItemDetailPage.jsx      # Detailansicht eines Items
│   │   │
│   │   ├── hooks/
│   │   │   └── useMedia.js             # API-Aufrufe für Media/Item CRUD
│   │   │
│   │   ├── services/                   # API Clients
│   │   │   ├── apiClient.js            # Fetch-Wrapper, hängt Bearer-Token an
│   │   │   ├── userapi.js              # Register/Login/Logout, User-Metadaten in localStorage
│   │   │   ├── collection.js           # Collection CRUD
│   │   │   ├── editionapi.js, inviteapi.js, settingsapi.js, tvdb.js
│   │   │   └── api.js
│   │   │
│   │   ├── types/
│   │   │   └── index.js                # MEDIA_TYPES, CONDITIONS, Konstanten
│   │   │
│   │   ├── App.jsx                     # Root Komponente, einfaches State-basiertes Routing
│   │   ├── App.css
│   │   ├── main.jsx                    # React Entry Point
│   │   └── index.css                   # Global Styles
│   │
│   ├── public/                         # Static Assets
│   ├── package.json                    # Node Dependencies
│   ├── vite.config.js                  # Vite Konfiguration
│   └── README.md
│
├── docker-compose.yml                  # Docker Orchestration
│                                        # Services: postgres, backend, pgadmin
│
├── .gitignore                          # Git ignore patterns
├── .env.example                        # Umgebungsvariablen Template
│
├── Makefile                            # Make Shortcuts (dev, docker, etc.)
├── start.sh                            # Interactive Start Script
│
├── README.md                           # Hauptdokumentation
├── STRUCTURE.md                        # Diese Datei
└── docs/DATA-MODEL.md                  # Detailliertes Datenbank-Schema
```

---

## 🔄 Datenfluss

### Request Flow (authentifizierter API Call, z. B. Item hinzufügen):

```
Browser (React)
    │
    ├─→ [CollectionPage.jsx]
    │   └─→ [services/collection.js] addItem()
    │       └─→ [apiClient.js] fetch() mit Authorization: Bearer <token>
    │           └─→ POST /api/v1/collections/:id/items
    │
REST API (Backend)
    │
    ├─→ [main.go] secure-Gruppe (JWTAuthMiddleware prüft Token, setzt user_id)
    │   └─→ [handlers/item.go] AddItem()
    │       ├─→ [services/collection.go] GetCollectionByID() → Ownership-Check (user_id == collection.UserID)
    │       ├─→ Input-Validierung (Title required)
    │       └─→ GORM Create(&item)
    │
PostgreSQL
    │
    └─→ INSERT INTO items (...)
```

---

## 📊 Datenbank Schema

Siehe [`docs/DATA-MODEL.md`](docs/DATA-MODEL.md) für das vollständige Schema inkl. Diagramm. Kurzfassung der aktiv über die API genutzten Tabellen:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,        -- bcrypt-Hash
    is_admin BOOLEAN DEFAULT false,
    created_at TIMESTAMP
);

CREATE TABLE collections (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    name VARCHAR NOT NULL,
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    collection_id INT NOT NULL REFERENCES collections(id),
    title VARCHAR NOT NULL,
    description TEXT,
    media_type VARCHAR,   -- 'bluray', 'dvd', 'vinyl', 'tape', ...
    artist VARCHAR,
    director VARCHAR,
    year INT,
    genre VARCHAR,
    condition VARCHAR,
    location VARCHAR,
    notes TEXT,
    tvdb_id VARCHAR,
    image_url VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE invite_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR UNIQUE NOT NULL,
    max_uses INT DEFAULT 0,
    current_uses INT DEFAULT 0,
    created_by_user_id INT NOT NULL REFERENCES users(id),
    used_by_user_id INT REFERENCES users(id),
    created_at TIMESTAMP,
    used_at TIMESTAMP,
    expires_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE settings (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    value BOOLEAN
);
```

Zusätzlich existiert ein normalisiertes Metadaten-Schema (`movies`, `persons`, `genres`, `editions`, `labels`, `media_types`, `conditions`, `movie_actors`, `movie_genres`, `collection_items`) für zukünftige Film-/Serien-Metadaten — per AutoMigration angelegt, aber aktuell nur `editions` über einen API-Endpoint erreichbar.

---

## 🔗 API Endpoints

```
GET    /health                              Health Check
POST   /api/v1/users                        Registrierung
POST   /api/v1/users/register               Registrierung (Alias)
POST   /api/v1/users/login                  Login

# ab hier: Authorization: Bearer <token> erforderlich
GET    /api/v1/users/:id
GET    /api/v1/collections
POST   /api/v1/collections
GET    /api/v1/collections/:id
PUT    /api/v1/collections/:id
DELETE /api/v1/collections/:id
POST   /api/v1/collections/:id/items
DELETE /api/v1/collections/:id/items/:itemId
GET    /api/v1/settings
GET    /api/v1/settings/:name
PUT    /api/v1/settings/:name
DELETE /api/v1/settings/:id
POST   /api/v1/invite/generate
GET    /api/v1/invite/list
DELETE /api/v1/invite/:id
GET    /api/v1/tvdb/search/series
GET    /api/v1/tvdb/search/movies
GET    /api/v1/editions
```

---

## 📦 Dependencies

### Backend (Go)
- `github.com/gin-gonic/gin` - HTTP Web Framework
- `gorm.io/gorm` + `gorm.io/driver/postgres` - ORM & PostgreSQL Driver
- `golang-jwt/jwt` - JWT Erzeugung/Validierung
- `golang.org/x/crypto/bcrypt` - Passwort-Hashing
- `joho/godotenv` - `.env`-Support
- `gopkg.in/yaml.v2` - YAML Parsing

### Frontend (React)
- `react` / `react-dom` (v19)
- `react-icons`
- `vite` - Build Tool & Dev Server

---

## 🔐 Architektur-Prinzipien

### Backend
1. **Layered Architecture:** Handler (HTTP/DTOs) → Service (Business-Logik) → GORM/DB
2. **JWT-Middleware:** trennt öffentliche Routen (Register/Login) von geschützten (`secure`-Gruppe in `main.go`)
3. **Ownership-Checks:** Collections/Items werden immer gegen `user_id` aus dem JWT-Context geprüft, bevor gelesen/geändert wird
4. **Zentrale Responses:** `handlers/respond.go` und `handlers/messages.go` vereinheitlichen Erfolgs-/Fehlerantworten
5. **YAML + ENV Configuration:** ENV-Variablen überschreiben `config.yaml`

### Frontend
1. **Component-Based:** Wiederverwendbare Komponenten
2. **State-basiertes Routing:** `App.jsx` steuert Seitenwechsel per State (kein Router-Package)
3. **API Service Layer:** Ein Client pro Ressource (`userapi`, `collection`, `editionapi`, ...) auf Basis von `apiClient.js`
4. **Token-Handling:** JWT wird nach Login in `localStorage`/`sessionStorage` gehalten und von `apiClient.js` als Bearer-Header mitgeschickt

---

## 🚀 Deployment Architektur

```
┌─────────────────────────────────────┐
│       Docker Compose Stack          │
├─────────────────────────────────────┤
│                                     │
│  ┌──────────────┐  ┌──────────────┐│
│  │  PostgreSQL  │  │  Go Backend  ││
│  │  (Port 5432) │  │ (Port 8080)  ││
│  └──────────────┘  └──────────────┘│
│                                     │
│  ┌──────────────┐  ┌──────────────┐│
│  │   pgAdmin    │  │   React Dev  ││
│  │ (Port 5050)  │  │ (Port 5173)  ││
│  └──────────────┘  └──────────────┘│
│                                     │
└─────────────────────────────────────┘
```

---

## 📝 Naming Conventions

### Backend (Go)
- **Packages:** lowercase, keine Unterstriche
- **Functions:** CamelCase, exported (großer Anfangsbuchstabe)
- **Types:** CamelCase
- **Fehlermeldungen/Konstanten:** in `handlers/messages.go` gebündelt (`msgXxx`)

### Frontend (React)
- **Components:** PascalCase (.jsx)
- **Hooks:** camelCase, prefix "use" (.js)
- **Styles:** ComponentName.css
- **Utils/Services:** camelCase (.js)

---

## 🔧 Konfiguration

### config.yaml (Backend)
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

### .env (Projekt-Root)
```
DATABASE_HOST=postgres
DATABASE_USER=fimuver_user
DATABASE_PASSWORD=fimuver_password
DATABASE_NAME=fimuver_db
JWT_SECRET=...
JWT_TTL=24h
TVDB_API_KEY=...
```

---

## 🔄 Versionskontrolle

### Wichtige .gitignore Einträge
- `/backend/bin/` - Compiled binaries
- `/frontend/node_modules/` - Node packages
- `/.env` - Sensitive secrets
- `/.idea/`, `/.vscode/` - IDE Konfiguration
- `/.serena/` - Tooling-Metadaten

---

## 📚 Weitere Ressourcen

- **Dokumentation:** `README.md`
- **Datenmodell:** `docs/DATA-MODEL.md`
- **Code-Kommentare:** In den einzelnen Dateien

---

## 🎯 Nächste Entwicklungsschritte

1. **Normalisiertes Metadaten-Schema verdrahten:** Movie/Person/Genre/CollectionItem existieren als Models, sind aber (bis auf `editions`) noch nicht über Handler/Endpoints erreichbar
2. **Item-Update-Endpoint** ergänzen (`PUT /collections/:id/items/:itemId`)
3. **Testing:** Unit Tests (Backend), Component Tests (Frontend) — aktuell nur `cors_test.go`
4. **Rate Limiting** für Login/Registrierung
5. **Pagination** für große Mediensammlungen
6. **Export:** CSV/PDF Export Funktionalität
7. **Statistiken:** Dashboard mit Grafiken

---

**Letzte Aktualisierung:** September 2026
