package handlers

import (
	"net/http"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

// ReferenceHandler verwaltet Referenztabellen
type ReferenceHandler struct {
	db *db.Database
}

func NewReferenceHandler(database *db.Database) *ReferenceHandler {
	return &ReferenceHandler{db: database}
}

// GetAllGenres GET /api/v1/references/genres
func (h *ReferenceHandler) GetAllGenres(c *gin.Context) {
	genres, err := h.db.GetAllGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Genres",
		})
		return
	}

	if genres == nil {
		genres = []models.Genre{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": genres,
	})
}

// GetAllMediaTypes GET /api/v1/references/media-types
func (h *ReferenceHandler) GetAllMediaTypes(c *gin.Context) {
	mediaTypes, err := h.db.GetAllMediaTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Medientypen",
		})
		return
	}

	if mediaTypes == nil {
		mediaTypes = []models.MediaType{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mediaTypes,
	})
}

// GetAllConditions GET /api/v1/references/conditions
func (h *ReferenceHandler) GetAllConditions(c *gin.Context) {
	conditions, err := h.db.GetAllConditions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Zustände",
		})
		return
	}

	if conditions == nil {
		conditions = []models.Condition{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": conditions,
	})
}

// GetAllEditions GET /api/v1/references/editions
func (h *ReferenceHandler) GetAllEditions(c *gin.Context) {
	editions, err := h.db.GetAllEditions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Editionen",
		})
		return
	}

	if editions == nil {
		editions = []models.Edition{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": editions,
	})
}

// GetAllLabels GET /api/v1/references/labels
func (h *ReferenceHandler) GetAllLabels(c *gin.Context) {
	labels, err := h.db.GetAllLabels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Labels",
		})
		return
	}

	if labels == nil {
		labels = []models.Label{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": labels,
	})
}

// CreateGenre POST /api/v1/references/genres
func (h *ReferenceHandler) CreateGenre(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	genre, err := h.db.GetOrCreateGenre(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Genres",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": genre,
	})
}

// CreateMediaType POST /api/v1/references/media-types
func (h *ReferenceHandler) CreateMediaType(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	mediaType, err := h.db.GetOrCreateMediaType(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Medientyps",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": mediaType,
	})
}

// CreateCondition POST /api/v1/references/conditions
func (h *ReferenceHandler) CreateCondition(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	condition, err := h.db.GetOrCreateCondition(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Zustands",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": condition,
	})
}

// CreateEdition POST /api/v1/references/editions
func (h *ReferenceHandler) CreateEdition(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	edition, err := h.db.GetOrCreateEdition(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen der Edition",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": edition,
	})
}

// CreateLabel POST /api/v1/references/labels
func (h *ReferenceHandler) CreateLabel(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	label, err := h.db.GetOrCreateLabel(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Labels",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": label,
	})
}
