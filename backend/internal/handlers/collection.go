package handlers

import (
	"strconv"

	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type CollectionHandler struct {
	db *db.Database
}

type CollectionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func NewCollectionHandler(db *db.Database) *CollectionHandler {
	return &CollectionHandler{db: db}
}
func (h *CollectionHandler) GetAllCollectionsForUser(c *gin.Context) {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	// Use CollectionService to fetch collections
	svc := services.NewCollectionService(h.db)
	collections, err := svc.GetCollectionsByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "error fetching collections"})
		return
	}

	c.JSON(200, gin.H{"data": collections})
}

func (h *CollectionHandler) CreateCollection(c *gin.Context) {
	var collectionDTO CollectionDTO
	if err := c.ShouldBindJSON(&collectionDTO); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	// get user id from context
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	// create model and persist via service
	svc := services.NewCollectionService(h.db)
	col := models.Collection{
		UserID:      userID,
		Name:        collectionDTO.Name,
		Description: collectionDTO.Description,
	}
	created, err := svc.AddCollection(col)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not create collection"})
		return
	}

	c.JSON(201, gin.H{"data": created})
}

func (h *CollectionHandler) UpdateCollection(c *gin.Context) {
	var collectionDTO CollectionDTO
	if err := c.ShouldBindJSON(&collectionDTO); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	svc := services.NewCollectionService(h.db)
	// parse collection id from path
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "collection id required"})
		return
	}
	// parse uint using strconv for clarity
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid collection id"})
		return
	}
	id := uint(id64)

	// load existing collection
	existing, err := svc.GetCollectionByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "collection not found"})
		return
	}

	// ownership check
	if existing.UserID != userID {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	updates := models.Collection{
		Name:        collectionDTO.Name,
		Description: collectionDTO.Description,
	}

	updated, err := svc.UpdateCollection(uint(id), updates)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not update collection"})
		return
	}

	c.JSON(200, gin.H{"data": updated})
	return
}
