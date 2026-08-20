package handlers

import (
	"strconv"
	"time"

	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type CollectionHandler struct {
	db *db.Database
}

type CollectionRequest struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description" required:"true"`
	Type        string `json:"type" required:"true"`
}

type CollectionResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCollectionHandler(db *db.Database) *CollectionHandler {
	return &CollectionHandler{db: db}
}

func (h *CollectionHandler) GetAllCollectionsForUser(c *gin.Context) {
	userID := c.GetUint("user_id")

	// Use CollectionService to fetch collections
	svc := services.NewCollectionService(h.db)
	collections, err := svc.GetCollectionsByUserID(userID)
	if err != nil {
		serverError(c, msgFetchCollections)
		return
	}

	ok(c, newCollectionArrayResponse(collections))
}

func (h *CollectionHandler) CreateCollection(c *gin.Context) {
	var collectionDTO CollectionRequest
	if err := c.ShouldBindJSON(&collectionDTO); err != nil {
		badRequest(c, msgInvalidRequestBody)
		return
	}

	userID := c.GetUint("user_id")

	// create model and persist via service
	svc := services.NewCollectionService(h.db)
	col := models.Collection{
		UserID:      userID,
		Name:        collectionDTO.Name,
		Description: collectionDTO.Description,
	}
	crea, err := svc.AddCollection(col)
	if err != nil {
		serverError(c, msgCreateCollection)
		return
	}

	created(c, msgCreateCollectionSuccess, newCollectionResponse(*crea))
}

func (h *CollectionHandler) UpdateCollection(c *gin.Context) {
	var collectionDTO CollectionRequest
	if err := c.ShouldBindJSON(&collectionDTO); err != nil {
		badRequest(c, msgInvalidRequestBody)
		return
	}

	userID := c.GetUint("user_id")

	svc := services.NewCollectionService(h.db)
	// parse collection id from path
	idStr := c.Param("id")
	if idStr == "" {
		badRequest(c, msgCollectionIDRequired)
		return
	}
	// parse uint using strconv for clarity
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		badRequest(c, msgInvalidCollectionID)
		return
	}
	id := uint(id64)

	// load existing collection
	existing, err := svc.GetCollectionByID(uint(id))
	if err != nil {
		notFound(c, msgCollectionNotFound)
		return
	}

	// ownership check
	if existing.UserID != userID {
		forbidden(c)
		return
	}

	updates := models.Collection{
		Name:        collectionDTO.Name,
		Description: collectionDTO.Description,
	}

	updated, err := svc.UpdateCollection(uint(id), updates)
	if err != nil {
		serverError(c, msgUpdateCollection)
		return
	}

	ok(c, newCollectionResponse(*updated))
}

func (h *CollectionHandler) GetCollectionByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		unauthorized(c)
		return
	}

	svc := services.NewCollectionService(h.db)
	col, err := svc.GetCollectionByID(uint(id64))
	if err != nil {
		notFound(c, msgCollectionNotFound)
		return
	}
	ok(c, newCollectionResponse(*col))
}

func (h *CollectionHandler) DeleteCollection(c *gin.Context) {
	userID := c.GetUint("user_id")

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		badRequest(c, msgCollectionIDRequired)
		return
	}

	svc := services.NewCollectionService(h.db)
	col, err := svc.DeleteCollection(uint(id64), userID)
	if err != nil {
		notFound(c, msgCollectionNotFound)
		return
	}
	okMessage(c, msgCollectionDeleted, newCollectionResponse(*col))
}

func newCollectionArrayResponse(c []models.Collection) []CollectionResponse {
	collections := make([]CollectionResponse, len(c))
	for _, col := range c {
		collections = append(collections, CollectionResponse{
			ID:          col.ID,
			UserID:      col.UserID,
			Name:        col.Name,
			Description: col.Description,
			CreatedAt:   col.CreatedAt,
			UpdatedAt:   col.UpdatedAt,
		})
	}
	return collections
}

func newCollectionResponse(c models.Collection) CollectionResponse {
	return CollectionResponse{
		ID:          c.ID,
		UserID:      c.UserID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
