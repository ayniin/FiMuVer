package models

import (
	"time"

	"gorm.io/datatypes"
)

// MediaType definiert die verschiedenen Medientypen
type MediaType string

const (
	MediaTypeBluray MediaType = "bluray"
	MediaTypeDVD    MediaType = "dvd"
	MediaTypeVinyl  MediaType = "vinyl"
	MediaTypeTape   MediaType = "tape"
)

// Media repräsentiert ein Medium (Bluray, DVD, Vinyl, Tape)
type Media struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"index;not null" json:"title"`
	Description string         `json:"description"`
	MediaType   MediaType      `gorm:"type:varchar(20);not null" json:"media_type"`
	Artist      string         `json:"artist"`   // Für Musik/Vinyl
	Director    string         `json:"director"` // Für Filme
	Year        int            `json:"year"`     // Erscheinungsjahr
	Genre       string         `json:"genre"`
	Condition   string         `json:"condition"` // z.B. "mint", "good", "fair", "poor"
	Location    string         `json:"location"`  // Wo wird es aufbewahrt
	Notes       datatypes.JSON `json:"notes"`     // JSON für flexible Metadaten
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Movie struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ImdbID      string `json:"imdb_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterPath  string `json:"poster_path"`
	Runtime     string `json:"runtime"`
	Genres      string `json:"genres"`
	ReleaseDate string `json:"release_date"`
	VoteAverage string `json:"vote_average"`
}

type Series struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	TvdbID         string `json:"tvdb_id"`
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
