package handlers

import (
	"fimuver/internal/config"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type TMDBHandler struct {
	tmdbClient *services.TMDBClient
}

func NewTMDBHandler(cfg *config.TMBDConfig) *TMDBHandler {
	return &TMDBHandler{
		tmdbClient: services.NewTMDBClient(cfg),
	}
}

func (h *TMDBHandler) SearchMovies(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.tmdbClient.SearchMovies(query)
	if err != nil {
		badGateway(c, msgTMDBUnavailable)
		return
	}

	ok(c, results)
}

func (h *TMDBHandler) SearchSeries(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.tmdbClient.SearchSeries(query)
	if err != nil {
		badGateway(c, msgTMDBUnavailable)
		return
	}

	ok(c, results)
}
