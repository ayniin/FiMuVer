package services

import (
	"fimuver/internal/db"
	"fimuver/internal/models"
)

type EditionService struct {
	db *db.Database
}

func NewEditionService(database *db.Database) *EditionService {
	return &EditionService{db: database}
}

func (s *EditionService) GetAllEditions() ([]models.Edition, error) {
	var editions []models.Edition
	if err := s.db.DB.Find(&editions).Error; err != nil {
		return nil, err
	}

	return editions, nil
}
