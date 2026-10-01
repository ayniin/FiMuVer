package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
}

type TVDBConfig struct {
	TvdbApiKey string `yaml:"tvdb_api_key"`
	BaseURL    string `yaml:"base_url"`
}

type TMBDConfig struct {
	TmdbApiKey   string `yaml:"tmdb_api_key"`
	BaseURL      string `yaml:"base_url"`
	ImageBaseURL string `yaml:"image_base_url"`
}

type DiscogsConfig struct {
	Token     string `yaml:"discogs_token"`
	BaseURL   string `yaml:"base_url"`
	UserAgent string `yaml:"user_agent"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type JWTConfig struct {
	Secret string        `yaml:"secret"`
	TTL    time.Duration `yaml:"ttl"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	TVDB     TVDBConfig     `yaml:"tvdb"`
	TMDB     TMBDConfig     `yaml:"tmdb"`
	Discogs  DiscogsConfig  `yaml:"discogs"`
	JWT      JWTConfig      `yaml:"jwt"`
}

// LoadConfig lädt die Konfiguration aus einer YAML-Datei
func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Lesen der Konfigurationsdatei: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("fehler beim Parsen der YAML-Datei: %w", err)
	}

	// Umgebungsvariablen auslesen und Konfiguration überschreiben (hat Vorrang)
	if host := os.Getenv("DATABASE_HOST"); host != "" {
		cfg.Database.Host = host
	}
	if port := os.Getenv("DATABASE_PORT"); port != "" {
		fmt.Sscanf(port, "%d", &cfg.Database.Port)
	}
	if user := os.Getenv("DATABASE_USER"); user != "" {
		cfg.Database.User = user
	}
	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		cfg.Database.Password = password
	}
	if dbname := os.Getenv("DATABASE_NAME"); dbname != "" {
		cfg.Database.Database = dbname
	}

	// Standard-Werte setzen, falls weder in YAML noch in ENV gesetzt
	if cfg.Server.Host == "" {
		cfg.Server.Host = "localhost"
	}
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	if cfg.Database.Port == 0 {
		cfg.Database.Port = 5432
	}
	if cfg.Database.SSLMode == "" {
		cfg.Database.SSLMode = "disable"
	}

	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		cfg.JWT.Secret = jwtSecret
	}
	if jwtTTL := os.Getenv("JWT_TTL"); jwtTTL != "" {
		cfg.JWT.TTL, _ = time.ParseDuration(jwtTTL)
	}

	if tvdbApiKey := os.Getenv("TVDB_API_KEY"); tvdbApiKey != "" {
		cfg.TVDB.TvdbApiKey = tvdbApiKey
	}
	if tvdbBaseURL := os.Getenv("TVDB_BASE_URL"); tvdbBaseURL != "" {
		cfg.TVDB.BaseURL = tvdbBaseURL
	}

	if cfg.TVDB.BaseURL == "" {
		cfg.TVDB.BaseURL = "https://api4.thetvdb.com/v4"
	}

	if tmdbApiKey := os.Getenv("TMDB_API_KEY"); tmdbApiKey != "" {
		cfg.TMDB.TmdbApiKey = tmdbApiKey
	}
	if tmdbBaseURL := os.Getenv("TMDB_BASE_URL"); tmdbBaseURL != "" {
		cfg.TMDB.BaseURL = tmdbBaseURL
	}

	if cfg.TMDB.BaseURL == "" {
		cfg.TMDB.BaseURL = "https://api.themoviedb.org/3"
	}

	if cfg.TMDB.ImageBaseURL == "" {
		cfg.TMDB.ImageBaseURL = "https://image.tmdb.org/t/p/w500"
	}

	if discogsToken := os.Getenv("DISCOGS_API_TOKEN"); discogsToken != "" {
		cfg.Discogs.Token = discogsToken
	}
	if discogsBaseURL := os.Getenv("DISCOGS_BASE_URL"); discogsBaseURL != "" {
		cfg.Discogs.BaseURL = discogsBaseURL
	}

	if cfg.Discogs.BaseURL == "" {
		cfg.Discogs.BaseURL = "https://api.discogs.com"
	}
	if cfg.Discogs.UserAgent == "" {
		cfg.Discogs.UserAgent = "FiMuVer/1.0"
	}

	return &cfg, nil
}

// GetDatabaseDSN gibt den PostgreSQL Connection String zurück
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}
