# 📁 FiMuVer - Projektstruktur und Architektur

## Überblick

FiMuVer ist eine **Full-Stack-Webanwendung** für die Verwaltung von Mediensammlungen.

- **Backend:** Go mit Gin Framework und GORM ORM
- **Frontend:** React mit Vite
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
│   │       └── main.go                 # Einstiegspunkt, Router-Setup
│   │
│   ├── internal/                       # Private Pakete
│   │   ├── config/
│   │   │   └── config.go              # YAML-Konfiguration laden/parsen
│   │   │
│   │   ├── models/
│   │   │   └── media.go               # Media Datenmodell (GORM)
│   │   │
│   │   ├── db/
│   │   │   └── database.go            # Datenbank-Abstraktionsebene (CRUD)
│   │   │
│   │   ├── handlers/
│   │   │   └── media.go               # HTTP Handler für /api/v1/media Endpoints
│   │   │
│   │   └── middleware/
│   │       └── cors.go                # CORS Middleware für React
│   │
│   ├── migrations/                    # (Placeholder) Datenbank Migrationen
│   │
│   ├── config.yaml                    # Datenbank-Konfiguration
│   ├── go.mod                         # Go Modul Definition
│   ├── go.sum                         # Go Abhängigkeiten Lockfile
│   └── Dockerfile                     # Multi-Stage Docker Build
│
├── 📁 frontend/                        # React Vite Projekt
│   ├── src/
│   │   ├── components/                # Reusable UI Komponenten
│   │   │   ├── MediaCard.jsx         # Anzeige einzelner Medien
│   │   │   ├── MediaCard.css
│   │   │   ├── MediaForm.jsx         # Formular zum Erstellen/Bearbeiten
│   │   │   ├── MediaForm.css
│   │   │   ├── FilterBar.jsx         # Filter & Suche
│   │   │   └── FilterBar.css
│   │   │
│   │   ├── pages/                    # Page-Komponenten
│   │   │   ├── Dashboard.jsx         # Hauptseite/Übersicht
│   │   │   └── Dashboard.css
│   │   │
│   │   ├── hooks/                    # Custom React Hooks
│   │   │   └── useMedia.js           # API-Aufrufe für Media CRUD
│   │   │
│   │   ├── services/                 # API Client
│   │   │   └── api.js                # MediaAPI Klasse (Fetch-Wrapper)
│   │   │
│   │   ├── types/                    # Datentypen & Konstanten
│   │   │   └── index.js              # MEDIA_TYPES, CONDITIONS, Models
│   │   │
│   │   ├── App.jsx                   # Root Komponente
│   │   ├── App.css
│   │   ├── main.jsx                  # React Entry Point
│   │   └── index.css                 # Global Styles
│   │
│   ├── public/                       # Static Assets
│   ├── package.json                  # Node Dependencies
│   ├── vite.config.js                # Vite Konfiguration
│   └── README.md
│
├── docker-compose.yml                # Docker Orchestration
│                                     # Services: postgres, backend, pgadmin
│
├── Dockerfile                        # (Root - würde hier nicht sein)
│
├── .gitignore                        # Git ignore patterns
├── .env.example                      # Umgebungsvariablen Template
│
├── Makefile                          # Make Shortcuts (dev, docker, etc.)
├── start.sh                          # Interactive Start Script
│
├── README.md                         # Hauptdokumentation
├── QUICK-START.md                    # Schnellstart Guide
└── STRUCTURE.md                      # Diese Datei
```

---

## 🔄 Datenfluss

### Request Flow (API Call):

```
Browser (React)
    │
    ├─→ [Dashboard.jsx]
    │   ├─→ useMedia() Hook
    │   │   └─→ [api.js] MediaAPI.getAllMedia()
    │   │       └─→ fetch() → POST /api/v1/media
    │   │
    │   └─→ [MediaCard.jsx] (Display)
    │
REST API (Backend)
    │
    ├─→ [main.go] Router setup
    │   └─→ /api/v1/media → [media.go Handler]
    │       ├─→ Input Validierung
    │       ├─→ [database.go] GetAllMedia()
    │       │   └─→ GORM Query
    │       │
    │       └─→ JSON Response
    │
PostgreSQL
    │
    └─→ SELECT * FROM media
```

---

## 📊 Datenbank Schema

```sql
CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT,
    media_type VARCHAR(20) NOT NULL,  -- 'bluray', 'dvd', 'vinyl', 'tape'
    artist VARCHAR,
    director VARCHAR,
    year INT,
    genre VARCHAR,
    condition VARCHAR,                -- 'mint', 'good', 'fair', 'poor'
    location VARCHAR,
    notes JSONB,                       -- Flexible Metadaten
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

---

## 🔗 API Endpoints

```
GET    /health
GET    /api/v1/media                    # Alle Medien (optional: ?type=bluray)
GET    /api/v1/media/:id                # Ein Medium
POST   /api/v1/media                    # Neues Medium erstellen
PUT    /api/v1/media/:id                # Medium aktualisieren
DELETE /api/v1/media/:id                # Medium löschen
GET    /api/v1/search?q=Matrix          # Suche nach Titel/Künstler/Regisseur
```

---

## 📦 Dependencies

### Backend (Go)
- `github.com/gin-gonic/gin` - HTTP Web Framework
- `gorm.io/gorm` - ORM für Datenbankoperationen
- `gorm.io/driver/postgres` - PostgreSQL Driver für GORM
- `gopkg.in/yaml.v2` - YAML Parsing

### Frontend (React)
- `react` - UI Framework
- `vite` - Build Tool & Dev Server

---

## 🔐 Architektur-Prinzipien

### Backend
1. **Layered Architecture:** Config → Handlers → Database
2. **YAML Configuration:** Externe Konfiguration statt Hardcoding
3. **GORM ORM:** Abstraktionsebene über Raw SQL
4. **Middleware Pattern:** CORS, Auth (zukünftig)

### Frontend
1. **Component-Based:** Wiederverwendbare Komponenten
2. **Custom Hooks:** Logik-Wiederverwendung (useMedia)
3. **API Service Layer:** Zentraler API Client
4. **Type System:** JavaScript Klassen für Datentypen

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
- **Constants:** UPPERCASE_WITH_UNDERSCORES

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
```

### .env.example (Docker/Umgebung)
```
DB_HOST=postgres
DB_USER=fimuver_user
DB_PASSWORD=fimuver_password
DB_NAME=fimuver_db
```

---

## 🔄 Versionskontrolle

### Wichtige .gitignore Einträge
- `/backend/bin/` - Compiled binaries
- `/frontend/node_modules/` - Node packages
- `/.env` - Sensitive secrets
- `/.idea/`, `/.vscode/` - IDE Konfiguration

---

## 📚 Weitere Ressourcen

- **Dokumentation:** `README.md`
- **Schnellstart:** `QUICK-START.md`
- **Code-Kommentare:** In den einzelnen Dateien

---

## 🎯 Nächste Entwicklungsschritte

1. **Authentifizierung:** JWT-Token in Backend implementieren
2. **Frontend Auth:** Login-Seite, Token-Speicherung
3. **Testing:** Unit Tests (Backend), Component Tests (Frontend)
4. **Validierung:** Input-Validierung Backend/Frontend
5. **Fehlerbehandlung:** Bessere Error-Messages
6. **Pagination:** Für große Mediensammlungen
7. **Export:** CSV/PDF Export Funktionalität
8. **Statistiken:** Dashboard mit Grafiken

---

**Version:** 1.0.0 (Initial Scaffold)  
**Letzte Aktualisierung:** April 2026

