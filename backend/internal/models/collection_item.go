package models

import "time"

// CollectionItem repräsentiert eine physische Kopie in einer Sammlung
// Verbindet eine Collection mit einem Film und speichert Details zur physischen Kopie
//
// Beispiel:
//
//	ID: 1
//	CollectionID: 1 (Meine Blurays)
//	FilmID: 1 (The Matrix)
//	MediaTypeID: 2 (4K Blu-ray)
//	Location: "Regal 1, Fach 3"
//	Condition: Mint
//
// Beziehungen:
//   - Gehört zu einer Collection (N:1)
//   - Referenziert einen Film (N:1)
//   - Hat Edition, Label, MediaType, Condition (N:1 jeweils)
type CollectionItem struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	CollectionID  uint       `gorm:"not null;index" json:"collection_id"`
	MovieID       uint       `gorm:"not null;index" json:"movie_id"`
	EditionID     *uint      `json:"edition_id"`
	LabelID       *uint      `json:"label_id"`
	MediaTypeID   uint       `gorm:"not null;index" json:"media_type_id"`
	ConditionID   *uint      `json:"condition_id"`
	Location      string     `json:"location"` // Lagerlocation
	PurchasePrice *float64   `json:"purchase_price"`
	PurchaseDate  *time.Time `json:"purchase_date"`
	Notes         string     `json:"notes"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`

	// Beziehungen
	Collection Collection `gorm:"foreignKey:CollectionID" json:"collection,omitempty"`
	Movie      Movie      `gorm:"foreignKey:MovieID" json:"movie,omitempty"`
	Edition    *Edition   `gorm:"foreignKey:EditionID" json:"edition,omitempty"`
	Label      *Label     `gorm:"foreignKey:LabelID" json:"label,omitempty"`
	MediaType  MediaType  `gorm:"foreignKey:MediaTypeID" json:"media_type,omitempty"`
	Condition  *Condition `gorm:"foreignKey:ConditionID" json:"condition,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (CollectionItem) TableName() string {
	return "collection_items"
}
