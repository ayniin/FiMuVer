package db

import (
	"fmt"

	"fimuver/internal/config"
	"fimuver/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database wrapper für GORM
type Database struct {
	DB *gorm.DB
}

// InitializeDatabase initialisiert die Datenbankverbindung und erstellt Tabellen
func InitializeDatabase(cfg *config.DatabaseConfig) (*Database, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("fehler beim Verbinden zur Datenbank: %w", err)
	}

	// Auto-Migration für Modelle
	if err := db.AutoMigrate(&models.Media{}); err != nil {
		return nil, fmt.Errorf("fehler bei der Datenbank-Migration: %w", err)
	}

	return &Database{DB: db}, nil
}

// Media Repository Methoden

// GetAllMedia holt alle Medien, optional gefiltert nach Typ
func (d *Database) GetAllMedia(mediaType string, limit uint, offset uint) ([]models.Media, error) {
	var media []models.Media
	query := d.DB

	if mediaType != "" {
		query = query.Where("media_type = ?", mediaType)
	}

	if limit > 0 {
		query = query.Limit(int(limit))
	}

	if offset > 0 {
		query = query.Offset(int(offset))
	}

	if err := query.Order("created_at DESC").Find(&media).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Medien: %w", err)
	}

	return media, nil
}

// GetMediaByID holt ein Medium anhand der ID
func (d *Database) GetMediaByID(id uint) (*models.Media, error) {
	var media models.Media
	if err := d.DB.First(&media, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("medium nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Mediums: %w", err)
	}

	return &media, nil
}

// CreateMedia erstellt ein neues Medium
func (d *Database) CreateMedia(media *models.Media) error {
	if err := d.DB.Create(media).Error; err != nil {
		return fmt.Errorf("fehler beim Erstellen des Mediums: %w", err)
	}

	return nil
}

// UpdateMedia aktualisiert ein existierendes Medium
func (d *Database) UpdateMedia(id uint, media *models.Media) error {
	if err := d.DB.Where("id = ?", id).Updates(media).Error; err != nil {
		return fmt.Errorf("fehler beim Aktualisieren des Mediums: %w", err)
	}

	return nil
}

// DeleteMedia löscht ein Medium
func (d *Database) DeleteMedia(id uint) error {
	if err := d.DB.Where("id = ?", id).Delete(&models.Media{}).Error; err != nil {
		return fmt.Errorf("fehler beim Löschen des Mediums: %w", err)
	}

	return nil
}

// SearchMedia sucht nach Medien nach Titel oder Künstler
func (d *Database) SearchMedia(query string) ([]models.Media, error) {
	var media []models.Media
	if err := d.DB.Where("title ILIKE ? OR artist ILIKE ? OR director ILIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Order("created_at DESC").
		Find(&media).Error; err != nil {
		return nil, fmt.Errorf("fehler bei der Suche: %w", err)
	}

	return media, nil
}

// Close schließt die Datenbankverbindung
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
