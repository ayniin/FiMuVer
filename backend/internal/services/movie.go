package services

import (
	"fimuver/internal/db"
	"fimuver/internal/models"
)

type MovieService struct {
	db *db.Database
}

func NewMovieService(db *db.Database) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) GetAllMoviesForUser(UserID uint, limit int, offset int) ([]models.Movie, int, error) {
	var movies []models.Movie
	if err := s.db.DB.Where("user_id = ?", UserID).Limit(limit).Offset(offset).Find(&movies).Error; err != nil {
		return nil, 0, err
	}

	return movies, len(movies), nil
}
