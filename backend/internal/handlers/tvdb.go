package handlers

import (
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
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.tvdbClient.SearchSeries(query)
	if err != nil {
		badGateway(c, msgTVDBUnavailable)
		return
	}

	ok(c, results)
}

func (h *TVDBHandler) SearchMovies(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.tvdbClient.SearchMovies(query)
	if err != nil {
		badGateway(c, msgTVDBUnavailable)
		return
	}

	ok(c, results)
}
