// Package models enthält alle Datenbankmodelle für FiMuVer
// Jedes Model ist in einer separaten Datei organisiert für bessere Wartbarkeit.
//
// Models:
// - user.go: Benutzer-Verwaltung
// - collection.go: Sammlungen von Benutzern
// - film.go: Film-Metadaten
// - person.go: Regisseure, Schauspieler und andere Personen
// - relationships.go: Many-to-Many Beziehungen (FilmActor, FilmGenre)
// - reference_types.go: Referenztabellen (Genre, Edition, Label, MediaType, Condition)
// - collection_item.go: Einzelne Kopien in einer Sammlung
package models
	Title          string `json:"title"`
	Description    string `json:"description"`
	PosterPath     string `json:"poster_path"`
	RuntimeAverage string `json:"runtime_average"`
	Genres         string `json:"genres"`
	ReleaseDate    string `json:"release_date"`
	VoteAverage    string `json:"vote_average"`
}

type Person struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ExternalID string `json:"external_id"`
	Name       string `gorm:"index;not null" json:"name"`
}

type Band struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"index;not null" json:"name"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Media) TableName() string {
	return "media"
}
