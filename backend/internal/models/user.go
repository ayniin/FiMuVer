package models

import "time"

// User repräsentiert einen Benutzer des Systems
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	// Relations
	Collections []Collection `gorm:"foreignKey:UserID" json:"collections,omitempty"`
}
