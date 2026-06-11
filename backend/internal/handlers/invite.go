package handlers

import (
	"fimuver/internal/db"
	"fimuver/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type InviteHandler struct {
	db *db.Database
}

func NewInviteHandler(database *db.Database) *InviteHandler {
	return &InviteHandler{db: database}
}

type GenerateInviteRequest struct {
	MaxUses   int    `json:"max_uses"`
	ExpiresAt string `json:"expires_at"` // ISO 8601, optional
}

func (h *InviteHandler) GenerateInviteCode(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req GenerateInviteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		req = GenerateInviteRequest{}
	}

	var expiresAt *time.Time
	if req.ExpiresAt != "" {
		t, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ungültiges Datumsformat, erwartet ISO 8601"})
			return
		}
		expiresAt = &t
	}

	svc := services.NewInviteService(h.db)
	invite, err := svc.CreateInviteCode(userID, req.MaxUses, expiresAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Invite Code erstellt",
		"data":    invite,
	})
}

func (h *InviteHandler) ListInviteCodes(c *gin.Context) {
	userID := c.GetUint("user_id")

	svc := services.NewInviteService(h.db)
	codes, err := svc.GetInviteCodesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": codes})
}

func (h *InviteHandler) UseInviteCode(c *gin.Context) {
	userID := c.GetUint("user_id")
	code := c.Param("code")

	svc := services.NewInviteService(h.db)
	invite, err := svc.ValidateInviteCode(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := svc.UseInviteCode(invite, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invite Code erfolgreich verwendet"})
}

func (h *InviteHandler) DeleteInviteCode(c *gin.Context) {
	userID := c.GetUint("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ungültige ID"})
		return
	}

	svc := services.NewInviteService(h.db)
	if err := svc.DeleteInviteCode(uint(id), userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invite Code gelöscht"})
}
