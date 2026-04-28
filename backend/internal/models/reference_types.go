package models

// Edition repräsentiert die Edition einer Kopie (Steelbook, Limited, etc.)
// Referenztabelle für verschiedene Editionen
//
// Beispiele:
//   - Steelbook
//   - Limited Edition
//   - Collector's Edition
//   - Standard
type Edition struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;uniqueIndex" json:"name"`

	// Beziehungen
	CollectionItems []CollectionItem `gorm:"foreignKey:EditionID" json:"collection_items,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Edition) TableName() string {
	return "editions"
}

// Label repräsentiert einen Publisher/Label (Warner Bros, Universal, etc.)
// Referenztabelle für Filme-Publisher
//
// Beispiele:
//   - Warner Bros.
//   - Universal Pictures
//   - Sony Pictures
//   - Paramount
type Label struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;uniqueIndex" json:"name"`

	// Beziehungen
	CollectionItems []CollectionItem `gorm:"foreignKey:LabelID" json:"collection_items,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Label) TableName() string {
	return "labels"
}

// MediaType repräsentiert den Medientyp (Blu-ray, 4K, DVD, Vinyl, Tape)
// Referenztabelle für Medienformate
//
// Beispiele:
//   - Blu-ray 1080p
//   - 4K Ultra HD
//   - DVD
//   - Vinyl LP
//   - Cassette Tape
type MediaType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;uniqueIndex" json:"name"`

	// Beziehungen
	CollectionItems []CollectionItem `gorm:"foreignKey:MediaTypeID" json:"collection_items,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (MediaType) TableName() string {
	return "media_types"
}

// Condition repräsentiert den Zustand einer Kopie (Mint, Good, Fair, Poor)
// Referenztabelle für Zustände
//
// Beispiele:
//   - Mint (Never opened)
//   - Near Mint
//   - Good
//   - Fair
//   - Poor
type Condition struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;uniqueIndex" json:"name"`

	// Beziehungen
	CollectionItems []CollectionItem `gorm:"foreignKey:ConditionID" json:"collection_items,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Condition) TableName() string {
	return "conditions"
}

// Genre repräsentiert ein Film-Genre
// Referenztabelle für Film-Genres
//
// Beispiele:
//   - Science Fiction
//   - Action
//   - Comedy
//   - Drama
//   - Horror
type Genre struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;uniqueIndex" json:"name"`

	// Beziehungen
	FilmGenres []FilmGenre `gorm:"foreignKey:GenreID" json:"film_genres,omitempty"`
}

// TableName gibt den Namen der Tabelle für GORM an
func (Genre) TableName() string {
	return "genres"
}

