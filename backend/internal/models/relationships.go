package models

// MovieActor repräsentiert die Beziehung zwischen Filmen und Schauspielern
// Many-to-Many Tabelle: Ein Film hat viele Schauspieler, ein Schauspieler spielt in vielen Filmen
//
// Beispiel:
//
//	FilmID: 1 (The Matrix)
//	PersonID: 5 (Keanu Reeves)
//	CharacterName: "Neo"
type MovieActor struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	MovieID       uint   `gorm:"not null;index" json:"movie_id"`
	PersonID      uint   `gorm:"not null;index" json:"person_id"`
	CharacterName string `json:"character_name"` // Rolle

	// Beziehungen
	Movie  Movie  `gorm:"foreignKey:MovieID" json:"movie,omitempty"`
	Person Person `gorm:"foreignKey:PersonID" json:"person,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (MovieActor) TableName() string {
	return "movie_actors"
}

// FilmGenre repräsentiert die Beziehung zwischen Filmen und Genres
// Many-to-Many Tabelle: Ein Film hat mehrere Genres, ein Genre hat mehrere Filme
//
// Beispiel:
//
//	FilmID: 1 (The Matrix)
//	GenreID: 3 (Science Fiction)
type MovieGenre struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	MovieID uint `gorm:"not null;index" json:"movie_id"`
	GenreID uint `gorm:"not null;index" json:"genre_id"`

	// Beziehungen
	Movie Movie `gorm:"foreignKey:MovieID" json:"Movie,omitempty"`
	Genre Genre `gorm:"foreignKey:GenreID" json:"genre,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (MovieGenre) TableName() string {
	return "movie_genres"
}
