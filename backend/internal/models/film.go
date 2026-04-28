package models

// Film repräsentiert Film-Metadaten
// Enthält alle wichtigen Informationen über einen Film (für Beschreibung)
//
// Beispiel:
//   ID: 1
//   ImdbID: "tt0133093"
//   Title: "The Matrix"
//   Year: 1999
//   DirectorID: 10
//
// Beziehungen:
//   - Hat einen Director (N:1 zu Person)
//   - Hat mehrere Actors (N:M über FilmActor)
//   - Hat mehrere Genres (N:M über FilmGenre)
//   - Kann in mehreren CollectionItems vorkommen (1:N)
type Film struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ImdbID      string     `gorm:"uniqueIndex;not null" json:"imdb_id"`
	Title       string     `gorm:"not null;index" json:"title"`
	Description string     `json:"description"`
	PosterPath  string     `json:"poster_path"`
	Runtime     string     `json:"runtime"`
	ReleaseDate string     `json:"release_date"`
	Year        int        `json:"year"`
	VoteAverage string     `json:"vote_average"`
	DirectorID  *uint      `json:"director_id"` // Primärer Director

	// Beziehungen
	Director        *Person             `gorm:"foreignKey:DirectorID" json:"director,omitempty"`
	FilmActors      []FilmActor         `gorm:"foreignKey:FilmID" json:"actors,omitempty"`
	FilmGenres      []FilmGenre         `gorm:"foreignKey:FilmID" json:"genres,omitempty"`
	CollectionItems []CollectionItem    `gorm:"foreignKey:FilmID" json:"collection_items,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Film) TableName() string {
	return "films"
}

