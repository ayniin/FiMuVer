package models

// Film repräsentiert Film-Metadaten
type Film struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	ImdbID     string     `gorm:"uniqueIndex;not null" json:"imdb_id"`
	Title      string     `gorm:"not null;index" json:"title"`
	Description string    `json:"description"`
	PosterPath string     `json:"poster_path"`
	Runtime    string     `json:"runtime"`
	ReleaseDate string    `json:"release_date"`
	Year       int        `json:"year"`
	VoteAverage string    `json:"vote_average"`
	DirectorID *uint      `json:"director_id"` // Primärer Director
	// Relations
	Director      *Person             `gorm:"foreignKey:DirectorID" json:"director,omitempty"`
	FilmActors    []FilmActor         `gorm:"foreignKey:FilmID" json:"actors,omitempty"`
	FilmGenres    []FilmGenre         `gorm:"foreignKey:FilmID" json:"genres,omitempty"`
	CollectionItems []CollectionItem `gorm:"foreignKey:FilmID" json:"collection_items,omitempty"`
}
