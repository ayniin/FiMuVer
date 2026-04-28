package db

import (
	"fmt"

	"fimuver/internal/config"
	"fimuver/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database wrapper für GORM
type Database struct {
	DB *gorm.DB
}

// InitializeDatabase initialisiert die Datenbankverbindung und erstellt Tabellen
func InitializeDatabase(cfg *config.DatabaseConfig) (*Database, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("fehler beim Verbinden zur Datenbank: %w", err)
	}

	// Auto-Migration für alle Modelle in korrekter Reihenfolge (Dependencies zuerst)
	if err := db.AutoMigrate(
		&models.User{},
		&models.Person{},
		&models.Genre{},
		&models.Edition{},
		&models.Label{},
		&models.MediaType{},
		&models.Condition{},
		&models.Film{},
		&models.FilmGenre{},
		&models.FilmActor{},
		&models.Collection{},
		&models.CollectionItem{},
	); err != nil {
		return nil, fmt.Errorf("fehler bei der Datenbank-Migration: %w", err)
	}

	return &Database{DB: db}, nil
}

// User Repository Methoden

// GetUserByID holt einen Benutzer anhand der ID
func (d *Database) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := d.DB.Preload("Collections").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("benutzer nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Benutzers: %w", err)
	}
	return &user, nil
}

// GetUserByEmail holt einen Benutzer anhand der E-Mail
func (d *Database) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := d.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("benutzer nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Benutzers: %w", err)
	}
	return &user, nil
}

// CreateUser erstellt einen neuen Benutzer
func (d *Database) CreateUser(user *models.User) error {
	if err := d.DB.Create(user).Error; err != nil {
		return fmt.Errorf("fehler beim Erstellen des Benutzers: %w", err)
	}
	return nil
}

// Collection Repository Methoden

// GetCollectionsByUserID holt alle Sammlungen eines Benutzers
func (d *Database) GetCollectionsByUserID(userID uint) ([]models.Collection, error) {
	var collections []models.Collection
	if err := d.DB.Where("user_id = ?", userID).Preload("Items").Find(&collections).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Sammlungen: %w", err)
	}
	return collections, nil
}

// GetCollectionByID holt eine Sammlung anhand der ID mit allen Items
func (d *Database) GetCollectionByID(id uint) (*models.Collection, error) {
	var collection models.Collection
	if err := d.DB.Preload("Items").Preload("User").First(&collection, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("sammlung nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen der Sammlung: %w", err)
	}
	return &collection, nil
}

// CreateCollection erstellt eine neue Sammlung
func (d *Database) CreateCollection(collection *models.Collection) error {
	if err := d.DB.Create(collection).Error; err != nil {
		return fmt.Errorf("fehler beim Erstellen der Sammlung: %w", err)
	}
	return nil
}

// UpdateCollection aktualisiert eine Sammlung
func (d *Database) UpdateCollection(id uint, collection *models.Collection) error {
	if err := d.DB.Model(&models.Collection{}).Where("id = ?", id).Updates(collection).Error; err != nil {
		return fmt.Errorf("fehler beim Aktualisieren der Sammlung: %w", err)
	}
	return nil
}

// DeleteCollection löscht eine Sammlung und alle zugehörigen Items
func (d *Database) DeleteCollection(id uint) error {
	if err := d.DB.Transaction(func(tx *gorm.DB) error {
		// Lösche zuerst alle CollectionItems
		if err := tx.Where("collection_id = ?", id).Delete(&models.CollectionItem{}).Error; err != nil {
			return err
		}
		// Dann die Sammlung
		return tx.Delete(&models.Collection{}, id).Error
	}); err != nil {
		return fmt.Errorf("fehler beim Löschen der Sammlung: %w", err)
	}
	return nil
}

// Film Repository Methoden

// GetFilmByIMDbID holt einen Film anhand der IMDb ID
func (d *Database) GetFilmByIMDbID(imdbID string) (*models.Film, error) {
	var film models.Film
	if err := d.DB.Preload("FilmActors.Person").Preload("FilmGenres.Genre").Preload("Director").
		Where("imdb_id = ?", imdbID).First(&film).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("film nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Films: %w", err)
	}
	return &film, nil
}

// GetFilmByID holt einen Film anhand der ID
func (d *Database) GetFilmByID(id uint) (*models.Film, error) {
	var film models.Film
	if err := d.DB.Preload("FilmActors.Person").Preload("FilmGenres.Genre").Preload("Director").
		First(&film, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("film nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Films: %w", err)
	}
	return &film, nil
}

// CreateFilm erstellt einen neuen Film
func (d *Database) CreateFilm(film *models.Film) error {
	if err := d.DB.Create(film).Error; err != nil {
		return fmt.Errorf("fehler beim Erstellen des Films: %w", err)
	}
	return nil
}

// SearchFilms sucht nach Filmen nach Titel
func (d *Database) SearchFilms(query string, limit, offset int) ([]models.Film, error) {
	var films []models.Film
	if err := d.DB.Preload("Director").Preload("FilmGenres.Genre").
		Where("title ILIKE ?", "%"+query+"%").
		Order("year DESC").
		Limit(limit).Offset(offset).
		Find(&films).Error; err != nil {
		return nil, fmt.Errorf("fehler bei der Filmsuche: %w", err)
	}
	return films, nil
}

// CollectionItem Repository Methoden

// GetCollectionItemsByCollectionID holt alle Items einer Sammlung
func (d *Database) GetCollectionItemsByCollectionID(collectionID uint, limit, offset int) ([]models.CollectionItem, error) {
	var items []models.CollectionItem
	if err := d.DB.Preload("Film").Preload("Edition").Preload("Label").Preload("MediaType").Preload("Condition").
		Where("collection_id = ?", collectionID).
		Limit(limit).Offset(offset).
		Find(&items).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Items: %w", err)
	}
	return items, nil
}

// GetCollectionItemByID holt ein CollectionItem anhand der ID
func (d *Database) GetCollectionItemByID(id uint) (*models.CollectionItem, error) {
	var item models.CollectionItem
	if err := d.DB.Preload("Film").Preload("Edition").Preload("Label").Preload("MediaType").Preload("Condition").
		First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("item nicht gefunden")
		}
		return nil, fmt.Errorf("fehler beim Abrufen des Items: %w", err)
	}
	return &item, nil
}

// CreateCollectionItem erstellt ein neues Item in einer Sammlung
func (d *Database) CreateCollectionItem(item *models.CollectionItem) error {
	if err := d.DB.Create(item).Error; err != nil {
		return fmt.Errorf("fehler beim Erstellen des Items: %w", err)
	}
	return nil
}

// UpdateCollectionItem aktualisiert ein CollectionItem
func (d *Database) UpdateCollectionItem(id uint, item *models.CollectionItem) error {
	if err := d.DB.Model(&models.CollectionItem{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return fmt.Errorf("fehler beim Aktualisieren des Items: %w", err)
	}
	return nil
}

// DeleteCollectionItem löscht ein CollectionItem
func (d *Database) DeleteCollectionItem(id uint) error {
	if err := d.DB.Delete(&models.CollectionItem{}, id).Error; err != nil {
		return fmt.Errorf("fehler beim Löschen des Items: %w", err)
	}
	return nil
}

// Genre Repository Methoden

// GetAllGenres holt alle Genres
func (d *Database) GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	if err := d.DB.Find(&genres).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Genres: %w", err)
	}
	return genres, nil
}

// GetOrCreateGenre holt oder erstellt ein Genre
func (d *Database) GetOrCreateGenre(name string) (*models.Genre, error) {
	var genre models.Genre
	if err := d.DB.Where("name = ?", name).First(&genre).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			genre.Name = name
			if err := d.DB.Create(&genre).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Genres: %w", err)
			}
		} else {
			return nil, fmt.Errorf("fehler beim Abrufen des Genres: %w", err)
		}
	}
	return &genre, nil
}

// MediaType Repository Methoden

// GetAllMediaTypes holt alle Medientypen
func (d *Database) GetAllMediaTypes() ([]models.MediaType, error) {
	var types []models.MediaType
	if err := d.DB.Find(&types).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Medientypen: %w", err)
	}
	return types, nil
}

// GetOrCreateMediaType holt oder erstellt einen Medientyp
func (d *Database) GetOrCreateMediaType(name string) (*models.MediaType, error) {
	var mediaType models.MediaType
	if err := d.DB.Where("name = ?", name).First(&mediaType).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			mediaType.Name = name
			if err := d.DB.Create(&mediaType).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Medientyps: %w", err)
			}
		} else {
			return nil, fmt.Errorf("fehler beim Abrufen des Medientyps: %w", err)
		}
	}
	return &mediaType, nil
}

// Condition Repository Methoden

// GetAllConditions holt alle Zustände
func (d *Database) GetAllConditions() ([]models.Condition, error) {
	var conditions []models.Condition
	if err := d.DB.Find(&conditions).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Zustände: %w", err)
	}
	return conditions, nil
}

// GetOrCreateCondition holt oder erstellt einen Zustand
func (d *Database) GetOrCreateCondition(name string) (*models.Condition, error) {
	var condition models.Condition
	if err := d.DB.Where("name = ?", name).First(&condition).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			condition.Name = name
			if err := d.DB.Create(&condition).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Zustands: %w", err)
			}
		} else {
			return nil, fmt.Errorf("fehler beim Abrufen des Zustands: %w", err)
		}
	}
	return &condition, nil
}

// Edition Repository Methoden

// GetAllEditions holt alle Editionen
func (d *Database) GetAllEditions() ([]models.Edition, error) {
	var editions []models.Edition
	if err := d.DB.Find(&editions).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Editionen: %w", err)
	}
	return editions, nil
}

// GetOrCreateEdition holt oder erstellt eine Edition
func (d *Database) GetOrCreateEdition(name string) (*models.Edition, error) {
	var edition models.Edition
	if err := d.DB.Where("name = ?", name).First(&edition).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			edition.Name = name
			if err := d.DB.Create(&edition).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen der Edition: %w", err)
			}
		} else {
			return nil, fmt.Errorf("fehler beim Abrufen der Edition: %w", err)
		}
	}
	return &edition, nil
}

// Label Repository Methoden

// GetAllLabels holt alle Labels
func (d *Database) GetAllLabels() ([]models.Label, error) {
	var labels []models.Label
	if err := d.DB.Find(&labels).Error; err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen der Labels: %w", err)
	}
	return labels, nil
}

// GetOrCreateLabel holt oder erstellt ein Label
func (d *Database) GetOrCreateLabel(name string) (*models.Label, error) {
	var label models.Label
	if err := d.DB.Where("name = ?", name).First(&label).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			label.Name = name
			if err := d.DB.Create(&label).Error; err != nil {
				return nil, fmt.Errorf("fehler beim Erstellen des Labels: %w", err)
			}
		} else {
			return nil, fmt.Errorf("fehler beim Abrufen des Labels: %w", err)
		}
	}
	return &label, nil
}

// Close schließt die Datenbankverbindung
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
