package handlers

import (
	"net/http"
	"strconv"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *db.Database
}

func NewUserHandler(database *db.Database) *UserHandler {
	return &UserHandler{db: database}
}

// AddUser godoc
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
func (h *UserHandler) AddUser(c *gin.Context) {
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
		IsAdmin:  false, // Neue User sind standardmäßig keine Admins
	}

	// 5. Speichere User in Datenbank (direkt mit GORM)
	if err := h.db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Benutzers: " + err.Error(),
		})
		return
	}

	// 6. Erfolgreiche Response mit neu erstelltem User
	c.JSON(http.StatusCreated, gin.H{
		"message": "Benutzer erfolgreich erstellt",
		"data": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
			"is_admin": user.IsAdmin,
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

// CreateUserRequest DTO für User-Erstellung
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=12,max=60"`
}
