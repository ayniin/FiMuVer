package handlers

import (
	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type EditionHandler struct {
	db *db.Database
}

type EditionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewEditionHandler(db *db.Database) *EditionHandler {
	return &EditionHandler{db: db}
}

func (h *EditionHandler) GetAllEditions(c *gin.Context) {
	service := services.NewEditionService(h.db)
	editions, err := service.GetAllEditions()
	if err != nil {
		serverError(c, msgFetchEditions)
		return
	}

	ok(c, newEditionResponse(editions))
}

func newEditionResponse(c []models.Edition) []EditionResponse {
	editions := make([]EditionResponse, len(c))
	for _, col := range c {
		editions = append(editions, EditionResponse{
			ID:   col.ID,
			Name: col.Name,
		})
	}
	return editions
}
