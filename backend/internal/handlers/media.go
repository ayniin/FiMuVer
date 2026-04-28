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
// Holt alle Medien mit Pagination
// Query-Parameter: ?limit=20&offset=0
func (h *MediaHandler) GetAllMedia(c *gin.Context) {
	limit := 20
	offset := 0

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if o := c.Query("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	var medias []models.Media
	if err := h.db.DB.
		Preload("Director").
		Preload("Genre").
		Preload("MediaType").
		Preload("Condition").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&medias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Medien",
		})
		return
	}

	if medias == nil {
		medias = []models.Media{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   medias,
		"count":  len(medias),
		"limit":  limit,
		"offset": offset,
	})
}

// GetMediaByID GET /api/v1/media/:id
// Holt ein einzelnes Medium anhand der ID
func (h *MediaHandler) GetMediaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	var media models.Media
	if err := h.db.DB.
		Preload("Director").
		Preload("Genre").
		Preload("MediaType").
		Preload("Condition").
		First(&media, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Medium nicht gefunden",
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
// @Param       media body CreateMediaRequest true "Medium Daten"
// @Success     201 {object} map[string]interface{} "Medium erstellt"
// @Failure     400 {object} map[string]string "Ungültiger Request"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /media [post]
func (h *MediaHandler) CreateMedia(c *gin.Context) {
	var req CreateMediaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body: " + err.Error(),
		})
		return
	}

	// Validierung
	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Titel ist erforderlich",
		})
		return
	}

	if req.MediaTypeID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "MediaTypeID ist erforderlich",
		})
		return
	}

	// Erstelle neues Media-Objekt
	media := models.Media{
		Title:       req.Title,
		Description: req.Description,
		DirectorID:  req.DirectorID,
		Artist:      req.Artist,
		Year:        req.Year,
		GenreID:     req.GenreID,
		MediaTypeID: req.MediaTypeID,
		ConditionID: req.ConditionID,
		Location:    req.Location,
		Notes:       req.Notes,
	}

	if err := h.db.DB.Create(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Mediums",
		})
		return
	}

	// Lade Beziehungen
	h.db.DB.
		Preload("Director").
		Preload("Genre").
		Preload("MediaType").
		Preload("Condition").
		First(&media, media.ID)

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
// @Param       media body CreateMediaRequest true "Aktualisierte Daten"
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
	var media models.Media
	if err := h.db.DB.First(&media, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Medium nicht gefunden",
		})
		return
	}

	var req CreateMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body: " + err.Error(),
		})
		return
	}

	// Aktualisiere nur gesetzte Felder
	if req.Title != "" {
		media.Title = req.Title
	}
	if req.Description != "" {
		media.Description = req.Description
	}
	if req.DirectorID != nil {
		media.DirectorID = req.DirectorID
	}
	if req.Artist != "" {
		media.Artist = req.Artist
	}
	if req.Year != 0 {
		media.Year = req.Year
	}
	if req.GenreID != nil {
		media.GenreID = req.GenreID
	}
	if req.MediaTypeID != 0 {
		media.MediaTypeID = req.MediaTypeID
	}
	if req.ConditionID != nil {
		media.ConditionID = req.ConditionID
	}
	if req.Location != "" {
		media.Location = req.Location
	}
	if req.Notes != "" {
		media.Notes = req.Notes
	}

	if err := h.db.DB.Save(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Aktualisieren des Mediums",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Medium erfolgreich aktualisiert",
		"data":    media,
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
	var media models.Media
	if err := h.db.DB.First(&media, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Medium nicht gefunden",
		})
		return
	}

	if err := h.db.DB.Delete(&media).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Löschen des Mediums",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Medium erfolgreich gelöscht",
	})
}

// SearchMedia godoc
// @Summary     Nach Medien suchen
// @Description Sucht nach Medien anhand von Titel, Künstler oder Director
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       q query string true "Suchbegriff"
// @Param       limit query int false "Limit (default: 20)"
// @Param       offset query int false "Offset (default: 0)"
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

	limit := 20
	offset := 0

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if o := c.Query("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	var medias []models.Media
	if err := h.db.DB.
		Preload("Director").
		Preload("Genre").
		Preload("MediaType").
		Preload("Condition").
		Where("title ILIKE ? OR artist ILIKE ?", "%"+query+"%", "%"+query+"%").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&medias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler bei der Suche",
		})
		return
	}

	if medias == nil {
		medias = []models.Media{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   medias,
		"count":  len(medias),
		"query":  query,
		"limit":  limit,
		"offset": offset,
	})
}

// CreateMediaRequest DTO für Media-Erstellung
type CreateMediaRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DirectorID  *uint  `json:"director_id"`
	Artist      string `json:"artist"`
	Year        int    `json:"year"`
	GenreID     *uint  `json:"genre_id"`
	MediaTypeID uint   `json:"media_type_id" binding:"required"`
	ConditionID *uint  `json:"condition_id"`
	Location    string `json:"location"`
	Notes       string `json:"notes"`
}

