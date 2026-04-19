# 📚 Swagger API Dokumentation

## 🎯 Zugang zur Swagger UI

Wenn der Backend läuft, öffne diese URL im Browser:

```
http://localhost:8080/swagger/index.html
```

Du siehst dort eine **interaktive API-Dokumentation** mit:
- ✅ Alle Endpoints aufgelistet
- ✅ Testmöglichkeit direkt aus der UI (Try it out)
- ✅ Request & Response Schemas
- ✅ Beispiel-Daten

---

## 🔄 Swagger nach Änderungen regenerieren

Wenn du neue Endpoints oder Änderungen an den Kommentaren machst:

```bash
cd backend
go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/api/main.go
```

Das regeneriert:
- `docs/docs.go` - Go Structs
- `docs/swagger.json` - JSON Specification
- `docs/swagger.yaml` - YAML Specification

---

## 📝 Swagger Kommentare für neue Endpoints

Wenn du einen neuen Endpoint hinzufügst, verwende dieses Format:

```go
// MyEndpoint godoc
// @Summary     Kurze Beschreibung
// @Description Längere Beschreibung
// @Tags        TagName
// @Accept      json
// @Produce     json
// @Param       id path uint true "Parameter Beschreibung"
// @Param       name query string false "Query Parameter"
// @Success     200 {object} map[string]interface{} "Erfolgreiche Response"
// @Failure     400 {object} map[string]string "Fehlerhafte Request"
// @Router      /path/{id} [get]
func (h *MyHandler) MyEndpoint(c *gin.Context) {
	// Implementation...
}
```

### Swagger Annotation Referenz:

| Annotation | Bedeutung |
|-----------|-----------|
| `@Summary` | Kurze Beschreibung (eine Zeile) |
| `@Description` | Längere Beschreibung |
| `@Tags` | Kategorie (z.B. "Media", "User") |
| `@Accept` | Input Format (json, xml, form) |
| `@Produce` | Output Format (json, xml) |
| `@Param` | Parameter (path, query, body) |
| `@Success` | Success Response Code & Type |
| `@Failure` | Error Response Code & Type |
| `@Router` | API Route & HTTP Method |

---

## 🚀 Aktuelle API Endpoints im Swagger

```
GET    /api/v1/media              - Alle Medien
GET    /api/v1/media/:id          - Ein Medium
POST   /api/v1/media              - Neues Medium
PUT    /api/v1/media/:id          - Medium aktualisieren
DELETE /api/v1/media/:id          - Medium löschen
GET    /api/v1/search             - Medien suchen
GET    /health                    - Health Check
```

---

## 💡 Tipps

### Test-Daten mit Swagger hinzufügen:

1. Öffne http://localhost:8080/swagger/index.html
2. Expandiere "POST /api/v1/media"
3. Klicke "Try it out"
4. Fülle das JSON-Formular aus
5. Klicke "Execute"

### API Spec exportieren:

Die Spec ist verfügbar unter:
- `docs/swagger.json` - JSON Format
- `docs/swagger.yaml` - YAML Format

Du kannst diese z.B. zu anderen Tools (Postman, etc.) importieren.

---

## 🔧 Konfiguration

Die Swagger-Konfiguration ist in `backend/docs.go`:

```go
// @title       FiMuVer API
// @version     1.0
// @host        localhost:8080
// @basePath    /api/v1
```

Ändere diese Werte, wenn dein API auf anderen Host/Port läuft!

---

## 📖 Weitere Ressourcen

- [Swaggo Documentation](https://github.com/swaggo/swag)
- [OpenAPI Specification](https://spec.openapis.org/oas/v3.0.3)
- [Gin-Swagger](https://github.com/swaggo/gin-swagger)

---

**Happy API Testing!** 🚀

