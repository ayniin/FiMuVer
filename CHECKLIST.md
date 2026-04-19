# ✅ FiMuVer - Deine Checkliste zum Abhaken

## 🎯 Installation & Setup

- [ ] **Projekt im Explorer öffnen**
  ```bash
  cd /home/felix/Projects/FiMuVer
  ls -la
  ```

- [ ] **Verifikation durchführen**
  ```bash
  bash verify-installation.sh
  ```
  Sollte alle ✓ grüne Haken zeigen.

- [ ] **Dokumentation lesen**
  Starte mit: `GET-STARTED.md` oder `README.md`

---

## 🚀 Erste Schritte

- [ ] **Docker starten**
  ```bash
  docker-compose up -d
  ```
  Warte ~20 Sekunden bis alle Services laufen.

- [ ] **Backend Test**
  ```bash
  curl http://localhost:8080/health
  ```
  Sollte `{"status":"OK"}` zurückgeben.

- [ ] **Frontend öffnen**
  http://localhost:5173
  
  Du solltest das FiMuVer Dashboard sehen.

- [ ] **Logs überprüfen**
  ```bash
  docker-compose logs -f backend
  ```
  Sollte keine Fehler zeigen.

---

## 🎬 Funktionality Test

- [ ] **Ein Medium hinzufügen**
  - Klicke "+ Neues Medium hinzufügen"
  - Fülle das Formular aus
  - Klicke "Hinzufügen"

- [ ] **Medium anschauen**
  - Das neue Medium sollte in der Liste erscheinen

- [ ] **Nach Medientyp filtern**
  - Klicke auf "Bluray", "DVD", etc.
  - Liste sollte sich filtern

- [ ] **Suche testen**
  - Gib einen Suchtext ein
  - Medien sollten gefiltert werden

- [ ] **Medium bearbeiten**
  - Klicke "Bearbeiten" auf einem Medium
  - Ändere etwas und speichere

- [ ] **Medium löschen**
  - Klicke "Löschen" auf einem Medium
  - Bestätige das Löschen

---

## 🔌 API Test

- [ ] **Health Check**
  ```bash
  curl http://localhost:8080/health
  ```

- [ ] **Alle Medien abrufen**
  ```bash
  curl http://localhost:8080/api/v1/media | jq
  ```

- [ ] **Ein Medium erstellen (curl)**
  ```bash
  curl -X POST http://localhost:8080/api/v1/media \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Inception",
      "media_type": "dvd",
      "director": "Christopher Nolan",
      "year": 2010,
      "genre": "Thriller"
    }'
  ```

- [ ] **Nach Medientyp filtern**
  ```bash
  curl http://localhost:8080/api/v1/media?type=bluray
  ```

- [ ] **Suche testen**
  ```bash
  curl "http://localhost:8080/api/v1/search?q=Inception"
  ```

---

## 🗄️ Datenbank Check

- [ ] **pgAdmin öffnen**
  http://localhost:5050
  
  - Email: admin@fimuver.local
  - Password: admin

- [ ] **Mit PostgreSQL verbinden**
  - Neuen Server hinzufügen
  - Host: postgres
  - Username: fimuver_user
  - Password: fimuver_password

- [ ] **media Tabelle anschauen**
  - Expand: Servers → fimuver_db → Schemas → public → Tables
  - Rechtsklick auf "media" → View Data

- [ ] **Direct SQL ausführen**
  ```bash
  docker-compose exec postgres psql \
    -U fimuver_user -d fimuver_db \
    -c "SELECT COUNT(*) as media_count FROM media;"
  ```

---

## 📝 Code-Exploration

- [ ] **Backend Code durchlesen**
  - [ ] `backend/cmd/api/main.go` - Server Setup
  - [ ] `backend/internal/handlers/media.go` - API Handler
  - [ ] `backend/internal/db/database.go` - CRUD Operationen
  - [ ] `backend/internal/config/config.go` - Konfiguration

- [ ] **Frontend Code durchlesen**
  - [ ] `frontend/src/pages/Dashboard.jsx` - Hauptseite
  - [ ] `frontend/src/components/MediaCard.jsx` - Komponente
  - [ ] `frontend/src/services/api.js` - API Client
  - [ ] `frontend/src/hooks/useMedia.js` - Custom Hook

---

## 🛠️ Tooling Test

- [ ] **Make Shortcuts ausprobieren**
  ```bash
  make help              # Alle Targets anschauen
  make docker-logs       # Logs live
  ```

- [ ] **Docker Compose Commands**
  ```bash
  docker-compose ps     # Status aller Services
  docker-compose down   # Alles stoppen
  docker-compose up -d  # Wieder starten
  ```

- [ ] **Start Script testen**
  ```bash
  ./start.sh            # Interaktives Menü
  ```

---

## 📚 Dokumentation Review

- [ ] **README.md**
  - Durchgelesen und verstanden? ✓

- [ ] **QUICK-START.md**
  - Sind Probleme + Lösungen nützlich? ✓

- [ ] **STRUCTURE.md**
  - Verstandst du die Architektur? ✓

- [ ] **PROJECT-SUMMARY.md**
  - Kennst du die Übersicht? ✓

---

## 🎓 Learning Checkpoints

- [ ] **Go verstehen**
  - [ ] Wie funktioniert die REST API in main.go?
  - [ ] Wie funktioniert CRUD in database.go?
  - [ ] Wie wird YAML konfiguriert?

- [ ] **React verstehen**
  - [ ] Wie funktioniert der useMedia Hook?
  - [ ] Wie wird API aufgerufen?
  - [ ] Wie funktioniert Component State?

- [ ] **Database verstehen**
  - [ ] Wie sieht das Schema aus?
  - [ ] Wie funktioniert GORM ORM?
  - [ ] Wie werden Queries ausgeführt?

- [ ] **Docker verstehen**
  - [ ] Wie funktioniert docker-compose.yml?
  - [ ] Wie werden Services orchestriert?
  - [ ] Wie funktionieren Multi-Stage Builds?

---

## 🚀 Deployment Vorbereitung

- [ ] **Environment Variables**
  ```bash
  cp .env.example .env
  # Konfiguriere nach Bedarf
  ```

- [ ] **Production Build Frontend**
  ```bash
  cd frontend
  npm run build
  # Erzeugt dist/ für Production
  ```

- [ ] **Production Build Backend**
  ```bash
  cd backend
  go build -o api ./cmd/api
  # Erzeugt binary
  ```

- [ ] **Docker Image bauen**
  ```bash
  docker-compose build
  ```

---

## 📊 Erweiterungs-Ideen

- [ ] **Authentifizierung** - JWT implementieren
- [ ] **Validierung** - Input Validierung verfeinern
- [ ] **Pagination** - Für große Listen
- [ ] **Statistiken** - Dashboard mit Charts
- [ ] **Cover Upload** - Bilder hochladen
- [ ] **Export** - CSV/PDF Export
- [ ] **Tests** - Unit & Integration Tests
- [ ] **Caching** - Performance verbessern

---

## 🆘 Häufige Issues

- [ ] **Port 8080 in Verwendung?**
  ```bash
  lsof -i :8080
  kill -9 <PID>
  ```

- [ ] **Datenbank-Fehler?**
  ```bash
  docker-compose down -v
  docker-compose up -d
  ```

- [ ] **Frontend lädt nicht?**
  ```bash
  cd frontend
  rm -rf node_modules
  npm install
  npm run dev
  ```

- [ ] **Go Dependencies Problem?**
  ```bash
  cd backend
  go mod tidy
  go mod download
  ```

---

## ✨ Final Checklist

- [ ] ✅ Installation verifiziert
- [ ] ✅ Docker starten funktioniert
- [ ] ✅ Frontend & Backend erreichbar
- [ ] ✅ API funktioniert
- [ ] ✅ Datenbank funktioniert
- [ ] ✅ Code verstanden
- [ ] ✅ Dokumentation gelesen
- [ ] ✅ Tools getestet
- [ ] ✅ Funktionalität getestet
- [ ] ✅ Bereit zur Erweiterung!

---

## 🎉 FINISHED!

Du hast alle Checkpoints abgehakt! 🎊

**Herzlichen Glückwunsch!** Du hast jetzt:

✅ Eine produktionsreife Full-Stack Anwendung  
✅ Alle Komponenten korrekt installiert  
✅ Die Funktionalität verifiziert  
✅ Die Architektur verstanden  
✅ Ein solides Fundament zum Weiterbauen  

**Nächste Phase:** Erweitere die Anwendung nach deinen Bedürfnissen!

---

## 📞 Quick Help

```bash
# Angebot: Alle wichtigen Kommandos
cd /home/felix/Projects/FiMuVer

# Starten
docker-compose up -d

# Logs
docker-compose logs -f backend

# Stoppen
docker-compose down

# DB Shell
docker-compose exec postgres psql -U fimuver_user -d fimuver_db

# Mit Make
make help
make docker-up
make docker-logs
```

---

**Viel Erfolg mit deinem FiMuVer Projekt!** 🚀🎬📀

*Erstellt: April 19, 2026*  
*Status: MVP Ready ✅*

