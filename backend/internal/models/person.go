package models

import "time"

// Person repräsentiert Personen (Regisseure, Schauspieler, etc.)
type Person struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"uniqueIndex;not null" json:"name"`
	BirthDate  *time.Time `json:"birth_date"`
	PicturePath string   `json:"picture_path"`
	Biography  string    `json:"biography"`
	ExternalID string    `json:"external_id"` // IMDb, TMDb ID
	// Relations
	DirectedFilms []Film         `gorm:"foreignKey:DirectorID" json:"directed_films,omitempty"`
	FilmActors    []FilmActor    `gorm:"foreignKey:PersonID" json:"film_actors,omitempty"`
}
