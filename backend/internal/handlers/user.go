package handlers

import (
	"net/http"
	"strconv"

	"fimuver/internal/auth"
	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *db.Database
}

func NewUserHandler(database *db.Database) *UserHandler {
	return &UserHandler{db: database}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Username   string `json:"username" binding:"required,min=3,max=50"`
	Password   string `json:"password" binding:"required,min=12,max=60"`
	InviteCode string `json:"invite_code"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
	IsAdmin  bool   `json:"is_admin"`
}

// RegisterUser godoc
// @Summary     Neuen Benutzer erstellen
// @Description Erstellt einen neuen Benutzer im System
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       user body CreateUserRequest true "Benutzer Daten"
// @Success     201 {object} map[string]interface{} "Benutzer erstellt"
// @Failure     400 {object} map[string]string "Ungültige Eingabe"
// @Failure     409 {object} map[string]string "Email oder Username bereits vorhanden"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /users [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, msgInvalidRequestBody)
		return
	}

	if req.Email == "" || req.Username == "" {
		badRequest(c, msgEmailUsernameReq)
		return
	}

	if len(req.Email) < 5 || len(req.Email) > 100 {
		badRequest(c, msgEmailLength)
		return
	}

	if len(req.Username) < 3 || len(req.Username) > 50 {
		badRequest(c, msgUsernameLength)
		return
	}

	if len(req.Password) < 12 || len(req.Password) > 60 {
		badRequest(c, msgPasswordLength)
		return
	}

	// Prüfe enable_registration Setting und validiere Invite Code
	var registrationSetting models.Settings
	inviteSvc := services.NewInviteService(h.db)
	settingResult := h.db.DB.Where("name = ?", "enable_registration").First(&registrationSetting)
	registrationEnabled := settingResult.Error == nil && registrationSetting.Value

	var inviteToUse *models.InviteCode

	if !registrationEnabled {
		if req.InviteCode == "" {
			fail(c, http.StatusForbidden, msgInviteCodeRequired)
			return
		}
		invite, err := inviteSvc.ValidateInviteCode(req.InviteCode)
		if err != nil {
			badRequest(c, msgInvalidInviteCode)
			return
		}
		inviteToUse = invite
	} else if req.InviteCode != "" {
		invite, err := inviteSvc.ValidateInviteCode(req.InviteCode)
		if err != nil {
			badRequest(c, msgInvalidInviteCode)
			return
		}
		inviteToUse = invite
	}

	var existingUser models.User
	result := h.db.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		conflict(c, msgEmailTaken)
		return
	}

	user := models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		IsAdmin:  false,
	}

	userService := services.NewUserService(h.db)
	newUser, err := userService.CreateUser(user)
	if err != nil {
		serverError(c, msgCreateUser)
		return
	}

	// Invite Code als verwendet markieren
	if inviteToUse != nil {
		_ = inviteSvc.UseInviteCode(inviteToUse, newUser.ID)
	}

	token, err := auth.GenerateToken(newUser.ID)
	if err != nil {
		serverError(c, msgTokenFailed)
		return
	}

	created(c, msgUserCreated, newUserResponse(newUser, token))
}

// GetUserByID GET /api/v1/users/:id
// Holt einen Benutzer anhand der ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		badRequest(c, msgInvalidUserID)
		return
	}

	var user models.User
	if err := h.db.DB.First(&user, uint(id)).Error; err != nil {
		notFound(c, msgUserNotFound)
		return
	}

	ok(c, newUserResponse(user, ""))
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var userRequest LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		badRequest(c, msgInvalidRequestBody)
		return
	}

	userService := services.NewUserService(h.db)
	user, token, err := userService.LoginUser(userRequest.Email, userRequest.Password)
	if err != nil {
		fail(c, http.StatusUnauthorized, msgInvalidLogin)
		return
	}

	okMessage(c, msgLoginSuccessful, newUserResponse(user, token))
}

func newUserResponse(u models.User, token string) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		IsAdmin:  u.IsAdmin,
		Token:    token,
	}
}
