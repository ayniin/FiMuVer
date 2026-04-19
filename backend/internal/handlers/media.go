package handlers

import (
	"net/http"
	"strconv"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	db *db.Database
}

func NewMediaHandler(database *db.Database) *MediaHandler {
	return &MediaHandler{db: database}
}

// GetAllMedia GET /api/v1/media
// Optional query parameter: ?type=bluray|dvd|vinyl|tape
func (h *MediaHandler) GetAllMedia(c *gin.Context) {
	mediaType := c.Query("type")

	media, err := h.db.GetAllMedia(mediaType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if media == nil {
		media = []models.Media{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": media,
	})
}

// GetMediaByID GET /api/v1/media/:id
func (h *MediaHandler) GetMediaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	media, err := h.db.GetMediaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": media,
	})
}

// CreateMedia godoc
// @Summary     Neues Medium erstellen
// @Description Erstellt ein neues Medium in der Datenbank
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       media body models.Media true "Medium Daten"
// @Success     201 {object} map[string]interface{} "Medium erstellt"
// @Failure     400 {object} map[string]string "Ungültiger Request"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /media [post]
func (h *MediaHandler) CreateMedia(c *gin.Context) {
	var media models.Media

	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body: " + err.Error(),
		})
		return
	}

	// Validierung
	if media.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Titel ist erforderlich",
		})
		return
	}

	if err := h.db.CreateMedia(&media); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    media,
		"message": "Medium erfolgreich erstellt",
	})
}

// UpdateMedia godoc
// @Summary     Medium aktualisieren
// @Description Aktualisiert ein existierendes Medium
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       id path uint true "Medium ID"
// @Param       media body models.Media true "Aktualisierte Daten"
// @Success     200 {object} map[string]string "Medium aktualisiert"
// @Failure     400 {object} map[string]string "Ungültige ID"
// @Failure     404 {object} map[string]string "Medium nicht gefunden"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /media/{id} [put]
func (h *MediaHandler) UpdateMedia(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	// Prüfe, ob das Medium existiert
	if _, err := h.db.GetMediaByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updateData models.Media
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body: " + err.Error(),
		})
		return
	}

	if err := h.db.UpdateMedia(uint(id), &updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Medium erfolgreich aktualisiert",
	})
}

// DeleteMedia godoc
// @Summary     Medium löschen
// @Description Löscht ein Medium aus der Datenbank
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       id path uint true "Medium ID"
// @Success     200 {object} map[string]string "Medium gelöscht"
// @Failure     400 {object} map[string]string "Ungültige ID"
// @Failure     404 {object} map[string]string "Medium nicht gefunden"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /media/{id} [delete]
func (h *MediaHandler) DeleteMedia(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	// Prüfe, ob das Medium existiert
	if _, err := h.db.GetMediaByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.db.DeleteMedia(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Medium erfolgreich gelöscht",
	})
}

// SearchMedia godoc
// @Summary     Nach Medien suchen
// @Description Sucht nach Medien anhand von Titel, Künstler oder Regisseur
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       q query string true "Suchbegriff"
// @Success     200 {object} map[string]interface{} "Suchergebnisse"
// @Failure     400 {object} map[string]string "Suchparameter erforderlich"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /search [get]
func (h *MediaHandler) SearchMedia(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Suchparameter 'q' ist erforderlich",
		})
		return
	}

	media, err := h.db.SearchMedia(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if media == nil {
		media = []models.Media{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": media,
	})
}
