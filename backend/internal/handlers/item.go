package handlers

import (
	"strconv"
	"time"

	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"

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
	Edition     string `json:"edition"`
	TVDBID      string `json:"tvdb_id"`
	ImdbID      string `json:"imdb_id"`
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
	Edition      string    `json:"edition"`
	TVDBID       string    `json:"tvdb_id"`
	ImdbID       string    `json:"imdb_id"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewItemHandler(db *db.Database) *ItemHandler {
	return &ItemHandler{db: db}
}

// ownedCollectionID parst :id und prüft, ob die Collection dem eingeloggten User gehört.
// Schreibt bei Fehler selbst die Response und gibt ok=false zurück.
func (h *ItemHandler) ownedCollectionID(c *gin.Context) (uint, bool) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		badRequest(c, msgInvalidCollectionID)
		return 0, false
	}

	col, err := services.NewCollectionService(h.db).GetCollectionByID(uint(id64))
	if err != nil {
		notFound(c, msgCollectionNotFound)
		return 0, false
	}
	if col.UserID != c.GetUint("user_id") {
		forbidden(c)
		return 0, false
	}
	return col.ID, true
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	collectionID, ok := h.ownedCollectionID(c)
	if !ok {
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

	item := models.Item{CollectionID: collectionID}
	applyItemRequest(&item, req)

	if err := h.db.DB.Create(&item).Error; err != nil {
		serverError(c, msgCreateItem)
		return
	}

	created(c, msgItemCreated, newItemResponse(item))
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	collectionID, ok := h.ownedCollectionID(c)
	if !ok {
		return
	}

	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 64)
	if err != nil {
		badRequest(c, msgInvalidItemID)
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

	var item models.Item
	if err := h.db.DB.Where("id = ? AND collection_id = ?", itemID, collectionID).First(&item).Error; err != nil {
		notFound(c, msgItemNotFound)
		return
	}

	applyItemRequest(&item, req)

	if err := h.db.DB.Save(&item).Error; err != nil {
		serverError(c, msgUpdateItem)
		return
	}

	okMessage(c, msgItemUpdated, newItemResponse(item))
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	collectionID, ok := h.ownedCollectionID(c)
	if !ok {
		return
	}

	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 64)
	if err != nil {
		badRequest(c, msgInvalidItemID)
		return
	}

	result := h.db.DB.Where("id = ? AND collection_id = ?", itemID, collectionID).Delete(&models.Item{})
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

func applyItemRequest(i *models.Item, req ItemRequest) {
	i.Title = req.Title
	i.Description = req.Description
	i.MediaType = req.MediaType
	i.Artist = req.Artist
	i.Director = req.Director
	i.Year = req.Year
	i.Genre = req.Genre
	i.Condition = req.Condition
	i.Location = req.Location
	i.Notes = req.Notes
	i.Edition = req.Edition
	i.TVDBID = req.TVDBID
	i.ImdbID = req.ImdbID
	i.ImageURL = req.ImageURL
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
		Edition:      i.Edition,
		TVDBID:       i.TVDBID,
		ImdbID:       i.ImdbID,
		ImageURL:     i.ImageURL,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}
