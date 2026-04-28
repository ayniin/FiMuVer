package services

import (
	"fmt"

	"fimuver/internal/db"
	"fimuver/internal/models"

	"gorm.io/gorm"
)

// ReferenceService kapselt die Business-Logic für Referenztabellen
type ReferenceService struct {
	db *db.Database
}

func NewReferenceService(database *db.Database) *ReferenceService {
	return &ReferenceService{db: database}
}

// Genres
func (s *ReferenceService) GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	if err := s.db.DB.Find(&genres).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Genres: %w", err)
	}
	return genres, nil
}

func (s *ReferenceService) GetOrCreateGenre(name string) (*models.Genre, error) {
	var genre models.Genre
	res := s.db.DB.Where("name = ?", name).First(&genre)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			genre.Name = name
			if err := s.db.DB.Create(&genre).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Genres: %w", err)
			}
			return &genre, nil
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Genres: %w", res.Error)
	}
	return &genre, nil
}

// MediaTypes
func (s *ReferenceService) GetAllMediaTypes() ([]models.MediaType, error) {
	var types []models.MediaType
	if err := s.db.DB.Find(&types).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Medientypen: %w", err)
	}
	return types, nil
}

func (s *ReferenceService) GetOrCreateMediaType(name string) (*models.MediaType, error) {
	var t models.MediaType
	res := s.db.DB.Where("name = ?", name).First(&t)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			t.Name = name
			if err := s.db.DB.Create(&t).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Medientyps: %w", err)
			}
			return &t, nil
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Medientyps: %w", res.Error)
	}
	return &t, nil
}

// Conditions
func (s *ReferenceService) GetAllConditions() ([]models.Condition, error) {
	var conditions []models.Condition
	if err := s.db.DB.Find(&conditions).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Zustände: %w", err)
	}
	return conditions, nil
}

func (s *ReferenceService) GetOrCreateCondition(name string) (*models.Condition, error) {
	var c models.Condition
	res := s.db.DB.Where("name = ?", name).First(&c)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			c.Name = name
			if err := s.db.DB.Create(&c).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Zustands: %w", err)
			}
			return &c, nil
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Zustands: %w", res.Error)
	}
	return &c, nil
}

// Editions
func (s *ReferenceService) GetAllEditions() ([]models.Edition, error) {
	var editions []models.Edition
	if err := s.db.DB.Find(&editions).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Editionen: %w", err)
	}
	return editions, nil
}

func (s *ReferenceService) GetOrCreateEdition(name string) (*models.Edition, error) {
	var e models.Edition
	res := s.db.DB.Where("name = ?", name).First(&e)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			e.Name = name
			if err := s.db.DB.Create(&e).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen der Edition: %w", err)
			}
			return &e, nil
		}
		return nil, fmt.Errorf("fehler beim Abrufen der Edition: %w", res.Error)
	}
	return &e, nil
}

// Labels
func (s *ReferenceService) GetAllLabels() ([]models.Label, error) {
	var labels []models.Label
	if err := s.db.DB.Find(&labels).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Labels: %w", err)
	}
	return labels, nil
}

func (s *ReferenceService) GetOrCreateLabel(name string) (*models.Label, error) {
	var l models.Label
	res := s.db.DB.Where("name = ?", name).First(&l)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			l.Name = name
			if err := s.db.DB.Create(&l).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Labels: %w", err)
			}
			return &l, nil
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Labels: %w", res.Error)
	}
	return &l, nil
}
