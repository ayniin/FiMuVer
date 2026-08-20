package handlers

import (
	"net/http"

	"fimuver/internal/config"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type TVDBHandler struct {
	tvdbClient *services.TVDBClient
}

func NewTVDBHandler(cfg *config.TVDBConfig) *TVDBHandler {

	return &TVDBHandler{
		tvdbClient: services.NewTVDBClient(cfg),
	}
}

func (h *TVDBHandler) SearchSeries(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query-Parameter 'q' fehlt"})
		return
	}

	results, err := h.tvdbClient.SearchSeries(query)
	if err != nil {

		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": results})

}

func (h *TVDBHandler) SearchMovies(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query-Parameter 'q' fehlt"})
		return
	}

	results, err := h.tvdbClient.SearchMovies(query)
	if err != nil {

		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
