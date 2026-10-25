package handlers

import (
	"fimuver/internal/config"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type TMDBHandler struct {
	tmdbClient *services.TMDBClient
}

func NewTMDBHandler(cfg *config.TMDBConfig) *TMDBHandler {
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

func (h *TMDBHandler) GetImdbID(c *gin.Context) {
	imdbID, err := h.tmdbClient.ImdbID(c.Param("id"))
	if err != nil {
		badGateway(c, msgTMDBUnavailable)
		return
	}

	ok(c, gin.H{"imdb_id": imdbID})
}
