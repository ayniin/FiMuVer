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

	// Auto-Migration für alle Modelle in korrekter Reihenfolge (Dependencies zuerst)
	if err := db.AutoMigrate(
		&models.User{},
		&models.Person{},
		&models.Genre{},
		&models.Edition{},
		&models.Label{},
		&models.MediaType{},
		&models.Condition{},
		&models.Movie{},
		&models.MovieGenre{},
		&models.MovieActor{},
		&models.Collection{},
		&models.CollectionItem{},
		&models.Settings{},
	); err != nil {
		return nil, fmt.Errorf("fehler bei der Datenbank-Migration: %w", err)
	}

	// Seed Default Settings
	if err := seedDefaultSettings(db); err != nil {
		return nil, fmt.Errorf("fehler beim Seeding der Settings: %w", err)
	}

	return &Database{DB: db}, nil
}

// seedDefaultSettings erstellt Default-Settings wenn keine vorhanden sind
func seedDefaultSettings(db *gorm.DB) error {
	// Definiere Default Settings
	defaultSettings := []models.Settings{
		{Name: "enable_registration", Value: false}, // Dark Mode enabled
	}

	// Für jedes Setting prüfe ob es bereits existiert
	for _, setting := range defaultSettings {
		var existing models.Settings
		result := db.Where("name = ?", setting.Name).First(&existing)

		// Wenn nicht existiert, erstelle es
		if result.RowsAffected == 0 {
			if err := db.Create(&setting).Error; err != nil {
				return fmt.Errorf("fehler beim Erstellen des Settings '%s': %w", setting.Name, err)
			}
		}
	}

	return nil
}

// Close schließt die Datenbankverbindung
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
