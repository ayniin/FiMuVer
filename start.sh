#!/bin/bash

# FiMuVer Startup Script

echo "🎬 FiMuVer - Medienverwaltungs-Software"
echo ""
echo "Wähle eine Option:"
echo "1) Docker (alle Services starten)"
echo "2) Lokal Backend (Go API)"
echo "3) Lokal Frontend (React)"
echo "4) Lokal Alles (Backend + Frontend + externe PostgreSQL)"
echo ""
read -p "Option wählen (1-4): " option

case $option in
  1)
    echo "🐳 Starte Docker Compose..."
    docker-compose up -d
    echo ""
    echo "✅ Services gestartet!"
    echo "Backend API: http://localhost:8080"
    echo "Frontend: http://localhost:5173"
    echo "pgAdmin: http://localhost:5050"
    echo ""
    echo "Logs anschauen: docker-compose logs -f"
    ;;
  2)
    echo "🚀 Starte Go Backend..."
    cd backend
    go mod tidy
    go run ./cmd/api/main.go
    ;;
  3)
    echo "⚛️  Starte React Frontend..."
    cd frontend
    npm install
    npm run dev
    ;;
  4)
    echo "🚀 Starte alles lokal (PostgreSQL muss laufen)..."
    echo ""
    echo "Vergewissere dich, dass PostgreSQL auf localhost:5432 erreichbar ist!"
    echo "mit Credentials:"
    echo "  User: fimuver_user"
    echo "  Password: fimuver_password"
    echo "  Database: fimuver_db"
    echo ""
    read -p "PostgreSQL ist bereit? (j/n): " ready

    if [ "$ready" = "j" ]; then
      # Starte Backend im Hintergrund
      echo "Starte Backend..."
      cd backend
      go mod tidy
      go run ./cmd/api/main.go &
      BACKEND_PID=$!

      # Warte ein bisschen
      sleep 2

      # Starte Frontend
      echo "Starte Frontend..."
      cd ../frontend
      npm install
      npm run dev

      # Cleanup
      trap "kill $BACKEND_PID" EXIT
    else
      echo "Breche ab."
      exit 1
    fi
    ;;
  *)
    echo "❌ Ungültige Option"
    exit 1
    ;;
esac

