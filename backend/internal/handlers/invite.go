package handlers

import (
	"strconv"
	"time"

	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"

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

type InviteResponse struct {
	ID          uint       `json:"id"`
	Code        string     `json:"code"`
	MaxUses     int        `json:"max_uses"`
	CurrentUses int        `json:"current_uses"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UsedAt      *time.Time `json:"used_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
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
			badRequest(c, msgInvalidDateFormat)
			return
		}
		expiresAt = &t
	}

	svc := services.NewInviteService(h.db)
	invite, err := svc.CreateInviteCode(userID, req.MaxUses, expiresAt)
	if err != nil {
		serverError(c, msgCreateInvite)
		return
	}

	created(c, msgInviteCreated, newInviteResponse(*invite))
}

func (h *InviteHandler) ListInviteCodes(c *gin.Context) {
	userID := c.GetUint("user_id")

	svc := services.NewInviteService(h.db)
	codes, err := svc.GetInviteCodesByUser(userID)
	if err != nil {
		serverError(c, msgFetchInvites)
		return
	}

	ok(c, newInviteArrayResponse(codes))
}

func (h *InviteHandler) DeleteInviteCode(c *gin.Context) {
	userID := c.GetUint("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		badRequest(c, msgInvalidInviteID)
		return
	}

	svc := services.NewInviteService(h.db)
	if err := svc.DeleteInviteCode(uint(id), userID); err != nil {
		notFound(c, msgInviteNotFound)
		return
	}

	okMessage(c, msgInviteDeleted, nil)
}

func newInviteArrayResponse(codes []models.InviteCode) []InviteResponse {
	invites := make([]InviteResponse, 0, len(codes))
	for _, code := range codes {
		invites = append(invites, newInviteResponse(code))
	}
	return invites
}

func newInviteResponse(i models.InviteCode) InviteResponse {
	return InviteResponse{
		ID:          i.ID,
		Code:        i.Code,
		MaxUses:     i.MaxUses,
		CurrentUses: i.CurrentUses,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		UsedAt:      i.UsedAt,
		ExpiresAt:   i.ExpiresAt,
	}
}
