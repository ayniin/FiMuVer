package handlers

import (
	"strconv"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	db *db.Database
}

type ItemDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MediaType   string `json:"media_type"`
	Artist      string `json:"artist"`
	Director    string `json:"director"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
	Condition   string `json:"condition"`
	Location    string `json:"location"`
	Notes       string `json:"notes"`
}

func NewItemHandler(db *db.Database) *ItemHandler {
	return &ItemHandler{db: db}
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	collectionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid collection id"})
		return
	}

	var dto ItemDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if dto.Title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	}

	item := models.Item{
		CollectionID: uint(collectionID),
		Title:        dto.Title,
		Description:  dto.Description,
		MediaType:    dto.MediaType,
		Artist:       dto.Artist,
		Director:     dto.Director,
		Year:         dto.Year,
		Genre:        dto.Genre,
		Condition:    dto.Condition,
		Location:     dto.Location,
		Notes:        dto.Notes,
	}

	if err := h.db.DB.Create(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "could not create item"})
		return
	}

	c.JSON(201, gin.H{"data": item})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item id"})
		return
	}

	result := h.db.DB.Delete(&models.Item{}, itemID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "could not delete item"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "item not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Item gelöscht"})
}
