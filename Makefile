.PHONY: help backend-install frontend-install docker-up docker-down docker-logs dev-backend dev-frontend test clean

help:
	@echo "FiMuVer - Medienverwaltungs-Software"
	@echo ""
	@echo "Verfügbare Kommandos:"
	@echo ""
	@echo "Entwicklung:"
	@echo "  make dev-backend     - Starte Go Backend mit Hot Reload"
	@echo "  make dev-frontend    - Starte React Frontend Dev Server"
	@echo "  make backend-install - Installiere Go Dependencies"
	@echo "  make frontend-install- Installiere Node Dependencies"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up       - Starte alle Docker Container"
	@echo "  make docker-down     - Stoppe alle Docker Container"
	@echo "  make docker-logs     - Zeige Docker Logs"
	@echo ""
	@echo "Verwaltung:"
	@echo "  make test            - Führe Tests aus"
	@echo "  make clean           - Cleanup"

# Backend Commands
backend-install:
	cd backend && go mod tidy && go mod download

dev-backend:
	cd backend && go run ./cmd/api/main.go

backend-test:
	cd backend && go test ./...

# Frontend Commands
frontend-install:
	cd frontend && npm install

dev-frontend:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build

# Docker Commands
docker-up:
	docker-compose up -d
	@echo "🚀 Services starten:"
	@echo "Backend API: http://localhost:8080"
	@echo "Frontend: http://localhost:5173"
	@echo "pgAdmin: http://localhost:5050"

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-db-reset:
	docker-compose down -v
	docker-compose up -d postgres

# Development Commands
install: backend-install frontend-install
	@echo "✅ Alle Dependencies installiert"

dev: docker-up
	@echo "🎯 Starte Entwicklungsumgebung..."
	@echo "Backend API: http://localhost:8080"
	@echo "Frontend: http://localhost:5173 (separat starten mit: make dev-frontend)"

test: backend-test
	@echo "✅ Alle Tests durchgeführt"

clean:
	cd backend && go clean
	cd frontend && rm -rf dist node_modules
	@echo "✅ Cleanup durchgeführt"

# Database Commands
db-shell:
	docker-compose exec postgres psql -U fimuver_user -d fimuver_db

# Utility
format:
	cd backend && go fmt ./...

lint:
	cd backend && go vet ./...

