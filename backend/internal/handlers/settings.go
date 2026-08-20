package handlers

import (
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

// CreateSettingRequest für neue Settings
type CreateSettingRequest struct {
	Name  string `json:"name" binding:"required"`
	Value bool   `json:"value"`
}

// UpdateSettingRequest für Settings-Updates
type UpdateSettingRequest struct {
	Value bool `json:"value"`
}

type SettingResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

// GetAllSettings GET /api/v1/settings
func (h *SettingsHandler) GetAllSettings(c *gin.Context) {
	var settings []models.Settings

	if err := h.db.DB.Find(&settings).Error; err != nil {
		serverError(c, msgFetchSettings)
		return
	}

	ok(c, newSettingArrayResponse(settings))
}

// GetSettingByName GET /api/v1/settings/:name
func (h *SettingsHandler) GetSettingByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		badRequest(c, msgSettingNameReq)
		return
	}

	var setting models.Settings
	if err := h.db.DB.Where("name = ?", name).First(&setting).Error; err != nil {
		notFound(c, msgSettingNotFound)
		return
	}

	ok(c, newSettingResponse(setting))
}

// UpdateSetting PUT /api/v1/settings/:name
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
		badRequest(c, msgSettingNameReq)
		return
	}

	var req UpdateSettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, msgInvalidRequestBody)
		return
	}

	var existing models.Settings
	if err := h.db.DB.Where("name = ?", name).First(&existing).Error; err != nil {
		notFound(c, msgSettingNotFound)
		return
	}

	if err := h.db.DB.Model(&existing).Update("value", req.Value).Error; err != nil {
		serverError(c, msgUpdateSetting)
		return
	}

	existing.Value = req.Value
	okMessage(c, msgSettingUpdated, newSettingResponse(existing))
}

// CreateSetting POST /api/v1/settings
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
		badRequest(c, msgInvalidRequestBody)
		return
	}

	if req.Name == "" {
		badRequest(c, msgSettingNameReq)
		return
	}

	var existing models.Settings
	result := h.db.DB.Where("name = ?", req.Name).First(&existing)
	if result.RowsAffected > 0 {
		conflict(c, msgSettingExists)
		return
	}

	setting := models.Settings{
		Name:  req.Name,
		Value: req.Value,
	}

	if err := h.db.DB.Create(&setting).Error; err != nil {
		serverError(c, msgCreateSetting)
		return
	}

	created(c, msgSettingCreated, newSettingResponse(setting))
}

// DeleteSetting DELETE /api/v1/settings/:id
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
		badRequest(c, msgInvalidSettingID)
		return
	}

	var setting models.Settings
	if err := h.db.DB.First(&setting, uint(id)).Error; err != nil {
		notFound(c, msgSettingNotFound)
		return
	}

	if err := h.db.DB.Delete(&setting).Error; err != nil {
		serverError(c, msgDeleteSetting)
		return
	}

	okMessage(c, msgSettingDeleted, nil)
}

func newSettingArrayResponse(s []models.Settings) []SettingResponse {
	settings := make([]SettingResponse, 0, len(s))
	for _, set := range s {
		settings = append(settings, newSettingResponse(set))
	}
	return settings
}

func newSettingResponse(s models.Settings) SettingResponse {
	return SettingResponse{
		ID:    s.ID,
		Name:  s.Name,
		Value: s.Value,
	}
}
