package handlers

import (
	"fimuver/internal/auth"
	"fimuver/internal/db"
	"fimuver/internal/models"
	"fimuver/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *db.Database
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=12,max=60"`
}

func NewUserHandler(database *db.Database) *UserHandler {
	return &UserHandler{db: database}
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
	// 1. Bind JSON Request zu CreateUserRequest DTO
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body: " + err.Error(),
		})
		return
	}

	// 2. Validierung der Input-Felder
	if req.Email == "" || req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email und Username sind erforderlich",
		})
		return
	}

	// Validiere Email Format (einfache Prüfung)
	if len(req.Email) < 5 || len(req.Email) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email muss zwischen 5 und 100 Zeichen lang sein",
		})
		return
	}

	// Validiere Username Länge
	if len(req.Username) < 3 || len(req.Username) > 50 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username muss zwischen 3 und 50 Zeichen lang sein",
		})
		return
	}

	// Validiere Password Länge
	if len(req.Password) < 12 || len(req.Password) > 60 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be between 12 and 60 characters long",
		})
	}

	// 3. Prüfe ob Email bereits existiert
	var existingUser models.User
	result := h.db.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email ist bereits registriert",
		})
		return
	}

	// 4. Erstelle neues User Objekt
	user := models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		IsAdmin:  false, // Neue User sind standardmäßig keine Admins
	}

	var userService = services.NewUserService(h.db)

	created, err := userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Erzeuge JWT Token für neu registrierten Nutzer
	token, err := auth.GenerateToken(created.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "konnte token nicht erstellen"})
		return
	}

	// 6. Erfolgreiche Response mit neu erstelltem User + Token
	c.JSON(http.StatusCreated, gin.H{
		"message": "Benutzer erfolgreich erstellt",
		"data": gin.H{
			"id":       created.ID,
			"email":    created.Email,
			"username": created.Username,
			"token":    token,
			"is_admin": created.IsAdmin,
		},
	})
}

// GetUserByID GET /api/v1/users/:id
// Holt einen Benutzer anhand der ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige User ID",
		})
		return
	}

	// Direkt GORM nutzen
	var user models.User
	if err := h.db.DB.Preload("Collections").First(&user, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Benutzer nicht gefunden",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var userRequest LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	var userService = services.NewUserService(h.db)
	user, token, err := userService.LoginUser(userRequest.Email, userRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid login",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
			"token":    token,
			"is_admin": user.IsAdmin,
		},
	})
}

// CreateUserRequest DTO für User-Erstellung
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=12,max=60"`
}
