package models

import "time"

// Item repräsentiert einen Mediaeintrag in einer Collection.
// Flache Struktur passend zum Frontend — keine FK-Referenztabellen.
type Item struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CollectionID uint      `gorm:"not null;index" json:"collection_id"`
	Title        string    `gorm:"not null" json:"title"`
	Description  string    `json:"description"`
	MediaType    string    `json:"media_type"`
	Artist       string    `json:"artist"`
	Director     string    `json:"director"`
	Year         int       `json:"year"`
	Genre        string    `json:"genre"`
	Condition    string    `json:"condition"`
	Location     string    `json:"location"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Item) TableName() string {
	return "items"
}
