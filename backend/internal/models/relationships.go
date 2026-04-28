package models

// FilmActor repräsentiert die Beziehung zwischen Filmem und Schauspielern
type FilmActor struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	FilmID        uint    `gorm:"not null;index" json:"film_id"`
	PersonID      uint    `gorm:"not null;index" json:"person_id"`
	CharacterName string  `json:"character_name"` // Rolle
	// Relations
	Film   Film   `gorm:"foreignKey:FilmID" json:"film,omitempty"`
	Person Person `gorm:"foreignKey:PersonID" json:"person,omitempty"`
}

// FilmGenre repräsentiert die Beziehung zwischen Filmen und Genres
type FilmGenre struct {
	ID      uint  `gorm:"primaryKey" json:"id"`
	FilmID  uint  `gorm:"not null;index" json:"film_id"`
	GenreID uint  `gorm:"not null;index" json:"genre_id"`
	// Relations
	Film  Film  `gorm:"foreignKey:FilmID" json:"film,omitempty"`
	Genre Genre `gorm:"foreignKey:GenreID" json:"genre,omitempty"`
}
