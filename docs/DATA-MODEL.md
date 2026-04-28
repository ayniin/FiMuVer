# 📊 FiMuVer Datenmodell

## Whiteboard Visualisierung

```mermaid
graph TB
    subgraph "👤 User Management"
        User["👤 User<br/>ID | Email | Name<br/>Password | Created_at"]
    end
    
    subgraph "📁 Collections"
        Collection["📁 Collection<br/>ID | UserID | Name<br/>Description | Created_at"]
        CollectionItem["📦 CollectionItem<br/>ID | CollectionID | MediaID<br/>Quantity | Condition | Location"]
    end
    
    subgraph "🎬 Media Types"
        Media["📺 Media<br/>ID | Title | Description<br/>MediaType | Year | Genre<br/>Condition | Location | Notes"]
        Film["🎥 Film<br/>ID | ImdbID | Title<br/>Runtime | Genres | ReleaseDate"]
        Series["📺 Series<br/>ID | TvdbID | Title<br/>RuntimeAverage | Genres"]
    end
    
    subgraph "👥 References"
        Person["👤 Person<br/>ID | ExternalID | Name"]
        Band["🎵 Band<br/>ID | Name"]
    end
    
    subgraph "🏷️ Reference Types"
        Genre["🏷️ Genre<br/>ID | Name"]
        MediaType["📋 MediaType<br/>ID | Type"]
        Condition["✨ Condition<br/>ID | Status"]
    end
    
    subgraph "🔗 Relationships"
        FilmActor["🎬 FilmActor<br/>FilmID | PersonID"]
        FilmGenre["🎥 FilmGenre<br/>FilmID | GenreID"]
    end
    
    User --> Collection
    Collection --> CollectionItem
    CollectionItem --> Media
    Media --> Film
    Media --> Series
    Film --> FilmActor
    Series --> Person
    FilmActor --> Person
    FilmGenre --> Genre
    Band --> Person
    
    style User fill:#e8f4f8
    style Collection fill:#f0e8f8
    style CollectionItem fill:#f0e8f8
    style Media fill:#fff3e0
    style Film fill:#fff3e0
    style Series fill:#fff3e0
    style Person fill:#e8f8e8
    style Band fill:#e8f8e8
    style Genre fill:#f8f0e8
    style MediaType fill:#f8f0e8
    style Condition fill:#f8f0e8
    style FilmActor fill:#f0f8e8
    style FilmGenre fill:#f0f8e8
```

## Model-Übersicht nach Datei

| Datei | Model | Beschreibung |
|-------|-------|---|
| `media.go` | Media | Hauptmedien-Tabelle (Bluray, DVD, Vinyl, Tape) |
| `film.go` | Film, Series | Film- und Serien-Metadaten |
| `user.go` | User | Benutzer-Verwaltung |
| `collection.go` | Collection | Sammlungen von Benutzern |
| `collection_item.go` | CollectionItem | Einzelne Kopien in einer Sammlung |
| `person.go` | Person, Band | Personen (Regisseure, Schauspieler, Künstler) |
| `relationships.go` | FilmActor, FilmGenre | Many-to-Many Beziehungen |
| `reference_types.go` | Genre, MediaType, Condition | Referenztabellen |

## Datenbank-Constraints

- 🔑 **Primary Keys:** Alle Tabellen haben `id` als PK
- 🔗 **Foreign Keys:** CollectionItem → Media/Collection, FilmActor → Film/Person, etc.
- 📍 **Indexed:** `title`, `name`, `email` für schnelle Suche
- ⛔ **NOT NULL:** title, name (wo sinnvoll)

---

*Generiert: April 2026*

