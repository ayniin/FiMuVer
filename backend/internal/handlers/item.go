package handlers

import (
	"strconv"
	"time"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	db *db.Database
}

type ItemRequest struct {
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
	TVDBID      string `json:"tvdb_id"`
	ImageURL    string `json:"image_url"`
}

type ItemResponse struct {
	ID           uint      `json:"id"`
	CollectionID uint      `json:"collection_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	MediaType    string    `json:"media_type"`
	Artist       string    `json:"artist"`
	Director     string    `json:"director"`
	Year         int       `json:"year"`
	Genre        string    `json:"genre"`
	Condition    string    `json:"condition"`
	Location     string    `json:"location"`
	Notes        string    `json:"notes"`
	TVDBID       string    `json:"tvdb_id"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewItemHandler(db *db.Database) *ItemHandler {
	return &ItemHandler{db: db}
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	collectionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		badRequest(c, msgInvalidCollectionID)
		return
	}

	var req ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, msgInvalidRequestBody)
		return
	}

	if req.Title == "" {
		badRequest(c, msgTitleRequired)
		return
	}

	item := models.Item{
		CollectionID: uint(collectionID),
		Title:        req.Title,
		Description:  req.Description,
		MediaType:    req.MediaType,
		Artist:       req.Artist,
		Director:     req.Director,
		Year:         req.Year,
		Genre:        req.Genre,
		Condition:    req.Condition,
		Location:     req.Location,
		Notes:        req.Notes,
		TVDBID:       req.TVDBID,
		ImageURL:     req.ImageURL,
	}

	if err := h.db.DB.Create(&item).Error; err != nil {
		serverError(c, msgCreateItem)
		return
	}

	created(c, msgItemCreated, newItemResponse(item))
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 64)
	if err != nil {
		badRequest(c, msgInvalidItemID)
		return
	}

	result := h.db.DB.Delete(&models.Item{}, itemID)
	if result.Error != nil {
		serverError(c, msgDeleteItem)
		return
	}
	if result.RowsAffected == 0 {
		notFound(c, msgItemNotFound)
		return
	}

	okMessage(c, msgItemDeleted, nil)
}

func newItemResponse(i models.Item) ItemResponse {
	return ItemResponse{
		ID:           i.ID,
		CollectionID: i.CollectionID,
		Title:        i.Title,
		Description:  i.Description,
		MediaType:    i.MediaType,
		Artist:       i.Artist,
		Director:     i.Director,
		Year:         i.Year,
		Genre:        i.Genre,
		Condition:    i.Condition,
		Location:     i.Location,
		Notes:        i.Notes,
		TVDBID:       i.TVDBID,
		ImageURL:     i.ImageURL,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}
