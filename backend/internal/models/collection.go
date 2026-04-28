package models

import "time"

// Collection repräsentiert eine Sammlung eines Benutzers
// Ein User kann mehrere Collections haben (z.B. "Blurays", "DVDs", "Vinyl")
//
// Beispiel:
//   ID: 1
//   UserID: 1
//   Name: "Meine Blurays"
//   Description: "Alle meine Blurays"
//
// Beziehungen:
//   - Gehört zu einem User (N:1)
//   - Enthält mehrere CollectionItems (1:N)
type Collection struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	UserID      uint             `gorm:"not null;index" json:"user_id"`
	Name        string           `gorm:"not null" json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`

	// Beziehungen
	User  User               `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items []CollectionItem   `gorm:"foreignKey:CollectionID" json:"items,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Collection) TableName() string {
	return "collections"
}

