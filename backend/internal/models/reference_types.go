package models

// Edition repräsentiert die Edition einer Kopie (Steelbook, Limited, etc.)
type Edition struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"` // Steelbook, Limited, Collector's, Standard
	// Relations
	CollectionItems []CollectionItem `gorm:"foreignKey:EditionID" json:"collection_items,omitempty"`
}

// Label repräsentiert einen Publisher/Label (Warner Bros, Universal, etc.)
type Label struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	// Relations
	CollectionItems []CollectionItem `gorm:"foreignKey:LabelID" json:"collection_items,omitempty"`
}

// MediaType repräsentiert den Medientyp (Blu-ray, 4K, DVD, Vinyl, Tape)
type MediaType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"` // Blu-ray, 4K, DVD, Vinyl, Tape
	// Relations
	CollectionItems []CollectionItem `gorm:"foreignKey:MediaTypeID" json:"collection_items,omitempty"`
}

// Condition repräsentiert den Zustand einer Kopie (Mint, Good, Fair, Poor)
type Condition struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"` // Mint, Good, Fair, Poor
	// Relations
	CollectionItems []CollectionItem `gorm:"foreignKey:ConditionID" json:"collection_items,omitempty"`
}

// Genre repräsentiert ein Film-Genre
type Genre struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	// Relations
	FilmGenres []FilmGenre `gorm:"foreignKey:GenreID" json:"film_genres,omitempty"`
}
