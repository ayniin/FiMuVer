package models

import "time"

// Person repräsentiert Personen (Regisseure, Schauspieler, etc.)
// Zentrale Tabelle für alle an Filmen beteiligten Personen
//
// Beispiel:
//   ID: 1
//   Name: "Christopher Nolan"
//   ExternalID: "nm0001104" (IMDb ID)
//   Biography: "British-American film director..."
//
// Beziehungen:
//   - Regie von Filmen (1:N)
//   - Schauspieler in Filmen (N:M über FilmActor)
type Person struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"uniqueIndex;not null" json:"name"`
	BirthDate  *time.Time `json:"birth_date"`
	PicturePath string    `json:"picture_path"`
	Biography  string     `json:"biography"`
	ExternalID string     `json:"external_id"` // IMDb, TMDb ID

	// Beziehungen
	DirectedFilms []Film         `gorm:"foreignKey:DirectorID" json:"directed_films,omitempty"`
	FilmActors    []FilmActor    `gorm:"foreignKey:PersonID" json:"film_actors,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Person) TableName() string {
	return "persons"
}

