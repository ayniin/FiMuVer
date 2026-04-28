// Package models enthält alle Datenbankmodelle für FiMuVer
// Jedes Model ist in einer separaten Datei organisiert für bessere Wartbarkeit.
//
// Models pro Datei:
// - media.go: Media (Blurays, DVDs, Vinyls, Tapes) - DIESE DATEI
// - user.go: User (Benutzer-Verwaltung)
// - collection.go: Collection (Sammlungen von Benutzern)
// - collection_item.go: CollectionItem (Einzelne Kopien in einer Sammlung)
// - film.go: Film (Film/Serien-Metadaten)
// - person.go: Person (Regisseure, Schauspieler, Künstler)
// - relationships.go: FilmActor, FilmGenre (Many-to-Many Beziehungen)
// - reference_types.go: Genre, Edition, Label, MediaType, Condition (Referenztabellen)
package models

import "time"

// Media repräsentiert ein physisches Medium (Bluray, DVD, Vinyl, Tape)
// Dies ist die Haupttabelle für alle physischen Medien in einer Collection
//
// Beispiel:
//   ID: 1
//   Title: "The Matrix"
//   MediaTypeID: 2 (4K Blu-ray)
//   DirectorID: 10 (Wachowski)
//   Year: 1999
//   ConditionID: 1 (Mint)
//   Location: "Regal 1, Fach 3"
type Media struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"index;not null" json:"title"`
	Description string    `json:"description"`
	DirectorID  *uint     `json:"director_id"`   // Für Filme - Referenz zu Person
	Artist      string    `json:"artist"`       // Für Musik/Vinyl
	Year        int       `json:"year"`         // Erscheinungsjahr
	GenreID     *uint     `json:"genre_id"`     // Referenz zu Genre
	MediaTypeID uint      `gorm:"not null;index" json:"media_type_id"` // Referenz zu MediaType
	ConditionID *uint     `json:"condition_id"` // Referenz zu Condition
	Location    string    `json:"location"`     // Wo wird es aufbewahrt
	Notes       string    `json:"notes"`        // Flexible Notizen
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Beziehungen
	Director      *Person       `gorm:"foreignKey:DirectorID" json:"director,omitempty"`
	Genre         *Genre        `gorm:"foreignKey:GenreID" json:"genre,omitempty"`
	MediaType     MediaType     `gorm:"foreignKey:MediaTypeID" json:"media_type,omitempty"`
	Condition     *Condition    `gorm:"foreignKey:ConditionID" json:"condition,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Media) TableName() string {
	return "media"
}
