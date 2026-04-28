package handlers

import (
	"net/http"
	"strconv"
	"time"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

type CollectionHandler struct {
	db *db.Database
}

func NewCollectionHandler(database *db.Database) *CollectionHandler {
	return &CollectionHandler{db: database}
}

// GetCollectionsByUser GET /api/v1/collections
// Holt alle Sammlungen des aktuellen Users
// TODO: User-ID aus JWT Token extrahieren
func (h *CollectionHandler) GetCollectionsByUser(c *gin.Context) {
	userID := c.GetUint("user_id") // Aus JWT Token

	collections, err := h.db.GetCollectionsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Sammlungen",
		})
		return
	}

	if collections == nil {
		collections = []models.Collection{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": collections,
	})
}

// GetCollectionByID GET /api/v1/collections/:id
func (h *CollectionHandler) GetCollectionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	collection, err := h.db.GetCollectionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sammlung nicht gefunden",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": collection,
	})
}

// CreateCollection POST /api/v1/collections
func (h *CollectionHandler) CreateCollection(c *gin.Context) {
	var req CreateCollectionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body",
		})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	userID := c.GetUint("user_id") // Aus JWT Token

	collection := models.Collection{
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.db.CreateCollection(&collection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen der Sammlung",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    collection,
		"message": "Sammlung erfolgreich erstellt",
	})
}

// UpdateCollection PUT /api/v1/collections/:id
func (h *CollectionHandler) UpdateCollection(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	var req CreateCollectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body",
		})
		return
	}

	collection, err := h.db.GetCollectionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sammlung nicht gefunden",
		})
		return
	}

	collection.Name = req.Name
	collection.Description = req.Description

	if err := h.db.UpdateCollection(uint(id), collection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Aktualisieren der Sammlung",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sammlung erfolgreich aktualisiert",
		"data":    collection,
	})
}

// DeleteCollection DELETE /api/v1/collections/:id
func (h *CollectionHandler) DeleteCollection(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	if err := h.db.DeleteCollection(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Löschen der Sammlung",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sammlung erfolgreich gelöscht",
	})
}

// CollectionItemHandler für Items in Sammlungen
type CollectionItemHandler struct {
	db *db.Database
}

func NewCollectionItemHandler(database *db.Database) *CollectionItemHandler {
	return &CollectionItemHandler{db: database}
}

// GetCollectionItems GET /api/v1/collections/:collectionId/items
func (h *CollectionItemHandler) GetCollectionItems(c *gin.Context) {
	collectionID, err := strconv.ParseUint(c.Param("collectionId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige Collection ID",
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

	items, err := h.db.GetCollectionItemsByCollectionID(uint(collectionID), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Items",
		})
		return
	}

	if items == nil {
		items = []models.CollectionItem{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   items,
		"count":  len(items),
		"limit":  limit,
		"offset": offset,
	})
}

// CreateCollectionItem POST /api/v1/collections/:collectionId/items
func (h *CollectionItemHandler) CreateCollectionItem(c *gin.Context) {
	collectionID, err := strconv.ParseUint(c.Param("collectionId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige Collection ID",
		})
		return
	}

	var req CreateCollectionItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body",
		})
		return
	}

	if req.FilmID == 0 || req.MediaTypeID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FilmID und MediaTypeID sind erforderlich",
		})
		return
	}

	item := models.CollectionItem{
		CollectionID: uint(collectionID),
		FilmID:       req.FilmID,
		MediaTypeID:  req.MediaTypeID,
		EditionID:    req.EditionID,
		LabelID:      req.LabelID,
		ConditionID:  req.ConditionID,
		Location:     req.Location,
		PurchasePrice: req.PurchasePrice,
		Notes:        req.Notes,
	}

	// Konvertiere Unix timestamp zu time.Time wenn vorhanden
	if req.PurchaseDate != nil {
		t := time.Unix(*req.PurchaseDate, 0)
		item.PurchaseDate = &t
	}

	if err := h.db.CreateCollectionItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Items",
		})
		return
	}

	// Lade mit Preload
	h.db.DB.
		Preload("Film").
		Preload("Edition").
		Preload("Label").
		Preload("MediaType").
		Preload("Condition").
		First(&item, item.ID)

	c.JSON(http.StatusCreated, gin.H{
		"data":    item,
		"message": "Item erfolgreich erstellt",
	})
}

// UpdateCollectionItem PUT /api/v1/collections/:collectionId/items/:itemId
func (h *CollectionItemHandler) UpdateCollectionItem(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige Item ID",
		})
		return
	}

	var req CreateCollectionItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body",
		})
		return
	}

	item, err := h.db.GetCollectionItemByID(uint(itemID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item nicht gefunden",
		})
		return
	}

	// Aktualisiere Felder
	if req.MediaTypeID != 0 {
		item.MediaTypeID = req.MediaTypeID
	}
	if req.EditionID != nil {
		item.EditionID = req.EditionID
	}
	if req.LabelID != nil {
		item.LabelID = req.LabelID
	}
	if req.ConditionID != nil {
		item.ConditionID = req.ConditionID
	}
	if req.Location != "" {
		item.Location = req.Location
	}
	if req.PurchasePrice != nil {
		item.PurchasePrice = req.PurchasePrice
	}
	if req.PurchaseDate != nil {
		t := time.Unix(*req.PurchaseDate, 0)
		item.PurchaseDate = &t
	}
	if req.Notes != "" {
		item.Notes = req.Notes
	}

	if err := h.db.UpdateCollectionItem(uint(itemID), item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Aktualisieren des Items",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item erfolgreich aktualisiert",
		"data":    item,
	})
}

// DeleteCollectionItem DELETE /api/v1/collections/:collectionId/items/:itemId
func (h *CollectionItemHandler) DeleteCollectionItem(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige Item ID",
		})
		return
	}

	if err := h.db.DeleteCollectionItem(uint(itemID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Löschen des Items",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item erfolgreich gelöscht",
	})
}

// Request DTOs

type CreateCollectionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateCollectionItemRequest struct {
	FilmID        uint       `json:"film_id" binding:"required"`
	MediaTypeID   uint       `json:"media_type_id" binding:"required"`
	EditionID     *uint      `json:"edition_id"`
	LabelID       *uint      `json:"label_id"`
	ConditionID   *uint      `json:"condition_id"`
	Location      string     `json:"location"`
	PurchasePrice *float64   `json:"purchase_price"`
	PurchaseDate  *int64     `json:"purchase_date"` // Unix timestamp
	Notes         string     `json:"notes"`
}




