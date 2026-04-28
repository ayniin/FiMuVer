package models

import "time"

// Collection repräsentiert eine Sammlung eines Benutzers
type Collection struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	UserID    uint             `gorm:"not null;index" json:"user_id"`
	Name      string           `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	// Relations
	User  User               `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items []CollectionItem   `gorm:"foreignKey:CollectionID" json:"items,omitempty"`
}
