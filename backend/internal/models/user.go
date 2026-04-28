package models

import "time"

// User repräsentiert einen Benutzer des Systems
// Model für Benutzer-Verwaltung und Authentifizierung
//
// Beispiel:
//
//	ID: 1
//	Email: "felix@example.com"
//	Username: "felix"
//	IsAdmin: false
//
// Beziehungen:
//   - Ein User hat mehrere Collections (1:N)
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	IsAdmin   bool      `gorm:"default:false" json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`

	// Beziehungen
	Collections []Collection `gorm:"foreignKey:UserID" json:"collections,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (User) TableName() string {
	return "users"
}
