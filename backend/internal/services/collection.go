package services

import (
	"fmt"

	"fimuver/internal/db"
	"fimuver/internal/models"
)

// CollectionService kapselt DB-Operationen für Collections
type CollectionService struct {
	db *db.Database
}

// NewCollectionService erstellt eine neue CollectionService Instanz
func NewCollectionService(database *db.Database) *CollectionService {
	return &CollectionService{db: database}
}

// GetCollectionsByUserID holt alle Collections für einen Benutzer inklusive Items
func (s *CollectionService) GetCollectionsByUserID(userID uint) ([]models.Collection, error) {
	var collections []models.Collection
	if err := s.db.DB.Where("user_id = ?", userID).Preload("Items").Find(&collections).Error; err != nil {
		return nil, fmt.Errorf("error while fetching the database: %w", err)
	}
	if collections == nil {
		collections = []models.Collection{}
	}
	return collections, nil
}

func (s *CollectionService) AddCollection(collection models.Collection) (*models.Collection, error) {
	if err := s.db.DB.Create(&collection).Error; err != nil {
		return nil, fmt.Errorf("error while creating the collection: %w", err)
	}
	return &collection, nil
}

// GetCollectionByID holt eine Collection anhand der ID
func (s *CollectionService) GetCollectionByID(id uint) (*models.Collection, error) {
	var col models.Collection
	if err := s.db.DB.Preload("Items").First(&col, id).Error; err != nil {
		return nil, fmt.Errorf("collection not found: %w", err)
	}
	return &col, nil
}

// UpdateCollection updated eine bestehende Collection (nur Felder, die gesetzt sind)
func (s *CollectionService) UpdateCollection(id uint, updates models.Collection) (*models.Collection, error) {
	col, err := s.GetCollectionByID(id)
	if err != nil {
		return nil, err
	}

	// nur erlaubte Felder überschreiben
	if updates.Name != "" {
		col.Name = updates.Name
	}
	if updates.Description != "" {
		col.Description = updates.Description
	}
	if updates.UserID != 0 {
		col.UserID = updates.UserID
	}

	if err := s.db.DB.Save(col).Error; err != nil {
		return nil, fmt.Errorf("failed to update collection: %w", err)
	}
	return col, nil
}
