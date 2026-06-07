// Package models enthält alle Datenbankmodelle für FiMuVer
// Jedes Model ist in einer separaten Datei organisiert für bessere Wartbarkeit.
//
// Models pro Datei:
// - media.go: Media (Blurays, DVDs, Vinyls, Tapes) - DIESE DATEI
// - user.go: User (Benutzer-Verwaltung)
// - collection.go: Collection (Sammlungen von Benutzern)
// - collection_item.go: CollectionItem (Einzelne Kopien in einer Sammlung)
// - movie.go: Film (Film/Serien-Metadaten)
// - person.go: Person (Regisseure, Schauspieler, Künstler)
// - relationships.go: MovieActor, FilmGenre (Many-to-Many Beziehungen)
// - reference_types.go: Genre, Edition, Label, MediaType, Condition (Referenztabellen)
package models

import "time"

// Media repräsentiert ein physisches Medium (Bluray, DVD, Vinyl, Tape)
// Dies ist die Haupttabelle für alle physischen Medien in einer Collection
//
// Beispiel:
//
//	ID: 1
//	Title: "The Matrix"
//	MediaTypeID: 2 (4K Blu-ray)
//	DirectorID: 10 (Wachowski)
//	Year: 1999
//	ConditionID: 1 (Mint)
//	Location: "Regal 1, Fach 3"
type InviteCode struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"uniqueIndex;not null" json:"code"`

	MaxUses     int        `gorm:"not null;default:0" json:"max_uses"`
	CurrentUses int        `gorm:"not null;default:0" json:"current_uses"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	UsedAt      *time.Time `json:"used_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	CreatedByUserID uint  `gorm:"not null;index" json:"created_by_user_id"`
	UsedByUserID    *uint `gorm:"index" json:"used_by_user_id"`

	// Beziehungen
	CreatedByUser User  `gorm:"foreignKey:CreatedByUserID" json:"created_by_user,omitempty"`
	UsedByUser    *User `gorm:"foreignKey:UsedByUserID" json:"used_by_user,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (InviteCode) TableName() string {
	return "invite_codes"
}
