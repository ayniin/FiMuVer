package handlers

import (
	"fimuver/internal/config"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type DiscogsHandler struct {
	discogsClient *services.DiscogsClient
}

func NewDiscogsHandler(cfg *config.DiscogsConfig) *DiscogsHandler {
	return &DiscogsHandler{
		discogsClient: services.NewDiscogsClient(cfg),
	}
}

func (h *DiscogsHandler) SearchReleases(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.discogsClient.SearchReleases(query)
	if err != nil {
		badGateway(c, msgDiscogsUnavailable)
		return
	}

	ok(c, results)
}

func (h *DiscogsHandler) SearchMasters(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.discogsClient.SearchMasters(query)
	if err != nil {
		badGateway(c, msgDiscogsUnavailable)
		return
	}

	ok(c, results)
}

func (h *DiscogsHandler) SearchArtists(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		badRequest(c, msgMissingQueryParam)
		return
	}

	results, err := h.discogsClient.SearchArtists(query)
	if err != nil {
		badGateway(c, msgDiscogsUnavailable)
		return
	}

	ok(c, results)
}
