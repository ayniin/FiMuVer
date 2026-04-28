package models

// FilmActor repräsentiert die Beziehung zwischen Filmen und Schauspielern
// Many-to-Many Tabelle: Ein Film hat viele Schauspieler, ein Schauspieler spielt in vielen Filmen
//
// Beispiel:
//   FilmID: 1 (The Matrix)
//   PersonID: 5 (Keanu Reeves)
//   CharacterName: "Neo"
type FilmActor struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	FilmID        uint    `gorm:"not null;index" json:"film_id"`
	PersonID      uint    `gorm:"not null;index" json:"person_id"`
	CharacterName string  `json:"character_name"` // Rolle

	// Beziehungen
	Film   Film   `gorm:"foreignKey:FilmID" json:"film,omitempty"`
	Person Person `gorm:"foreignKey:PersonID" json:"person,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (FilmActor) TableName() string {
	return "film_actors"
}

// FilmGenre repräsentiert die Beziehung zwischen Filmen und Genres
// Many-to-Many Tabelle: Ein Film hat mehrere Genres, ein Genre hat mehrere Filme
//
// Beispiel:
//   FilmID: 1 (The Matrix)
//   GenreID: 3 (Science Fiction)
type FilmGenre struct {
	ID      uint  `gorm:"primaryKey" json:"id"`
	FilmID  uint  `gorm:"not null;index" json:"film_id"`
	GenreID uint  `gorm:"not null;index" json:"genre_id"`

	// Beziehungen
	Film  Film  `gorm:"foreignKey:FilmID" json:"film,omitempty"`
	Genre Genre `gorm:"foreignKey:GenreID" json:"genre,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (FilmGenre) TableName() string {
	return "film_genres"
}

