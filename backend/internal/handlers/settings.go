package handlers

import (
	"net/http"
	"strconv"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"github.com/gin-gonic/gin"
)

// SettingsHandler verwaltet Anwendungs-Settings (in DB)
type SettingsHandler struct {
	db *db.Database
}

func NewSettingsHandler(database *db.Database) *SettingsHandler {
	return &SettingsHandler{db: database}
}

// GetAllSettings GET /api/v1/settings
// Holt alle Einstellungen
func (h *SettingsHandler) GetAllSettings(c *gin.Context) {
	var settings []models.Settings

	if err := h.db.DB.Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Abrufen der Settings",
		})
		return
	}

	if settings == nil {
		settings = []models.Settings{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": settings,
	})
}

// GetSettingByName GET /api/v1/settings/:name
// Holt eine einzelne Einstellung nach Name
func (h *SettingsHandler) GetSettingByName(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Setting-Name ist erforderlich",
		})
		return
	}

	var setting models.Settings
	if err := h.db.DB.Where("name = ?", name).First(&setting).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Setting nicht gefunden",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": setting,
	})
}

// UpdateSetting PUT /api/v1/settings/:name
// Aktualisiert eine Einstellung nach Name
// @Summary     Setting aktualisieren
// @Description Aktualisiert eine Anwendungs-Einstellung
// @Tags        Settings
// @Accept      json
// @Produce     json
// @Param       name path string true "Setting Name"
// @Param       setting body UpdateSettingRequest true "Setting Daten"
// @Success     200 {object} map[string]interface{} "Setting aktualisiert"
// @Failure     400 {object} map[string]string "Ungültige Eingabe"
// @Failure     404 {object} map[string]string "Setting nicht gefunden"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /settings/{name} [put]
func (h *SettingsHandler) UpdateSetting(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Setting-Name ist erforderlich",
		})
		return
	}

	var req UpdateSettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body",
		})
		return
	}

	// Prüfe ob Setting existiert
	var existing models.Settings
	if err := h.db.DB.Where("name = ?", name).First(&existing).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Setting nicht gefunden",
		})
		return
	}

	// Aktualisiere
	if err := h.db.DB.Model(&existing).Update("value", req.Value).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Aktualisieren des Settings",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Setting erfolgreich aktualisiert",
		"data": gin.H{
			"name":  existing.Name,
			"value": req.Value,
		},
	})
}

// CreateSetting POST /api/v1/settings
// Erstellt eine neue Einstellung
// @Summary     Setting erstellen
// @Description Erstellt eine neue Anwendungs-Einstellung
// @Tags        Settings
// @Accept      json
// @Produce     json
// @Param       setting body CreateSettingRequest true "Setting Daten"
// @Success     201 {object} map[string]interface{} "Setting erstellt"
// @Failure     400 {object} map[string]string "Ungültige Eingabe"
// @Failure     409 {object} map[string]string "Setting existiert bereits"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /settings [post]
func (h *SettingsHandler) CreateSetting(c *gin.Context) {
	var req CreateSettingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültiger Request-Body",
		})
		return
	}

	// Validierung
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name ist erforderlich",
		})
		return
	}

	// Prüfe ob Setting bereits existiert
	var existing models.Settings
	result := h.db.DB.Where("name = ?", req.Name).First(&existing)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Setting mit diesem Namen existiert bereits",
		})
		return
	}

	// Erstelle neues Setting
	setting := models.Settings{
		Name:  req.Name,
		Value: req.Value,
	}

	if err := h.db.DB.Create(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Erstellen des Settings",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Setting erfolgreich erstellt",
		"data":    setting,
	})
}

// DeleteSetting DELETE /api/v1/settings/:id
// Löscht ein Setting
// @Summary     Setting löschen
// @Description Löscht eine Anwendungs-Einstellung
// @Tags        Settings
// @Accept      json
// @Produce     json
// @Param       id path uint true "Setting ID"
// @Success     200 {object} map[string]string "Setting gelöscht"
// @Failure     400 {object} map[string]string "Ungültige ID"
// @Failure     404 {object} map[string]string "Setting nicht gefunden"
// @Failure     500 {object} map[string]string "Interner Fehler"
// @Router      /settings/{id} [delete]
func (h *SettingsHandler) DeleteSetting(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige Setting ID",
		})
		return
	}

	// Prüfe ob Setting existiert
	var setting models.Settings
	if err := h.db.DB.First(&setting, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Setting nicht gefunden",
		})
		return
	}

	// Lösche
	if err := h.db.DB.Delete(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fehler beim Löschen des Settings",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Setting erfolgreich gelöscht",
	})
}

// Request DTOs

// CreateSettingRequest für neue Settings
type CreateSettingRequest struct {
	Name  string `json:"name" binding:"required"`
	Value bool   `json:"value"`
}

// UpdateSettingRequest für Settings-Updates
type UpdateSettingRequest struct {
	Value bool `json:"value"`
}
