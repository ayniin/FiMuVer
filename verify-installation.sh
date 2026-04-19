#!/bin/bash

# FiMuVer - Installation Verification Script

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║          🎬 FiMuVer - Installation Verification 🎬           ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

ERRORS=0
WARNINGS=0

# Farben
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

check_file() {
    if [ -f "$1" ]; then
        echo -e "${GREEN}✓${NC} $1"
    else
        echo -e "${RED}✗${NC} $1"
        ((ERRORS++))
    fi
}

check_dir() {
    if [ -d "$1" ]; then
        echo -e "${GREEN}✓${NC} $1/"
    else
        echo -e "${RED}✗${NC} $1/"
        ((ERRORS++))
    fi
}

echo "📁 Backend Dateien:"
check_file "backend/cmd/api/main.go"
check_file "backend/internal/config/config.go"
check_file "backend/internal/models/media.go"
check_file "backend/internal/db/database.go"
check_file "backend/internal/handlers/media.go"
check_file "backend/internal/middleware/cors.go"
check_file "backend/config.yaml"
check_file "backend/go.mod"
check_file "backend/Dockerfile"
echo ""

echo "📁 Frontend Dateien:"
check_file "frontend/src/pages/Dashboard.jsx"
check_file "frontend/src/components/MediaCard.jsx"
check_file "frontend/src/components/MediaForm.jsx"
check_file "frontend/src/components/FilterBar.jsx"
check_file "frontend/src/hooks/useMedia.js"
check_file "frontend/src/services/api.js"
check_file "frontend/src/types/index.js"
check_file "frontend/src/App.jsx"
check_file "frontend/package.json"
echo ""

echo "🐳 Docker & Konfiguration:"
check_file "docker-compose.yml"
check_file ".env.example"
check_file ".gitignore"
echo ""

echo "📚 Dokumentation:"
check_file "README.md"
check_file "QUICK-START.md"
check_file "STRUCTURE.md"
check_file "PROJECT-SUMMARY.md"
echo ""

echo "🛠️  Tools:"
check_file "Makefile"
check_file "start.sh"
echo ""

echo "📦 Go Dependencies:"
if [ -f "backend/go.mod" ]; then
    DEPS=$(grep -c "^require" backend/go.mod 2>/dev/null || echo "0")
    if [ "$DEPS" -gt 0 ]; then
        echo -e "${GREEN}✓${NC} go.mod mit Dependencies"
    else
        echo -e "${YELLOW}⚠${NC} go.mod könnte Dependencies benötigen"
        ((WARNINGS++))
    fi
else
    echo -e "${RED}✗${NC} backend/go.mod nicht gefunden"
    ((ERRORS++))
fi
echo ""

echo "📊 Code Statistiken:"
BACKEND_LINES=$(find backend -name "*.go" -type f | xargs wc -l 2>/dev/null | tail -1 | awk '{print $1}')
FRONTEND_LINES=$(find frontend/src -name "*.jsx" -o -name "*.js" -o -name "*.css" 2>/dev/null | xargs wc -l 2>/dev/null | tail -1 | awk '{print $1}')
echo -e "Backend Go Code:        ~${BACKEND_LINES} Zeilen"
echo -e "Frontend React Code:    ~${FRONTEND_LINES} Zeilen"
echo ""

echo "════════════════════════════════════════════════════════════════"
echo ""

if [ $ERRORS -eq 0 ]; then
    echo -e "${GREEN}✅ Alle Dateien vorhanden!${NC}"
else
    echo -e "${RED}❌ $ERRORS Fehler gefunden!${NC}"
fi

if [ $WARNINGS -gt 0 ]; then
    echo -e "${YELLOW}⚠️  $WARNINGS Warnungen${NC}"
fi

echo ""
echo "════════════════════════════════════════════════════════════════"
echo ""

if [ $ERRORS -eq 0 ]; then
    echo "🚀 NÄCHSTE SCHRITTE:"
    echo ""
    echo "1️⃣  Dokumentation lesen:"
    echo "   less README.md"
    echo ""
    echo "2️⃣  Mit Docker starten:"
    echo "   docker-compose up -d"
    echo ""
    echo "3️⃣  Oder manuell starten:"
    echo "   ./start.sh"
    echo ""
    echo "4️⃣  Frontend öffnen:"
    echo "   http://localhost:5173"
    echo ""
    echo "════════════════════════════════════════════════════════════════"
    echo ""
    echo "📝 Tipps:"
    echo "  • make help           - Alle Make Targets anschauen"
    echo "  • make docker-up      - Docker starten"
    echo "  • make docker-logs    - Logs anschauen"
    echo "  • ./start.sh          - Interaktives Start-Menu"
    echo ""
    exit 0
else
    echo "❌ Installation unvollständig!"
    echo "Einige Dateien fehlen. Bitte repository neu clonen oder Dateien hinzufügen."
    exit 1
fi

