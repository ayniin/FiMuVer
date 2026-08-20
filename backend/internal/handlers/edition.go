package handlers

import (
	"fimuver/internal/db"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type EditionHandler struct {
	db *db.Database
}

func NewEditionHandler(db *db.Database) *EditionHandler {
	return &EditionHandler{db: db}
}

func (h *EditionHandler) GetAllEditions(c *gin.Context) {
	_, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "user not authenticated"})
		return
	}

	service := services.NewEditionService(h.db)
	editions, err := service.GetAllEditions()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, editions)
}
