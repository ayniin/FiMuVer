package main

import (
	"fmt"
	"log"

	"fimuver/internal/auth"
	"fimuver/internal/config"
	"fimuver/internal/db"
	"fimuver/internal/handlers"
	"fimuver/internal/middleware"

	_ "fimuver/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Lade Konfiguration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Fehler beim Laden der Konfiguration: %v", err)
	}

	// Gib die geladene Konfiguration aus
	log.Printf("Geladene Konfiguration: %+v\n", cfg)
	log.Printf("Datenbank DSN: %s\n", cfg.Database.GetDSN())

	// Initialisiere Datenbank
	database, err := db.InitializeDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Fehler beim Initialisieren der Datenbank: %v", err)
	}
	defer database.Close()

	// Init Auth (JWT) mit Werten aus Config
	// cfg.JWT.TTL ist a time.Duration parsed from YAML or ENV
	var ttlSeconds int
	if cfg.JWT.TTL > 0 {
		ttlSeconds = int(cfg.JWT.TTL.Seconds())
	}
	auth.Init(cfg.JWT.Secret, ttlSeconds)

	// Erstelle Gin Router
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORSMiddleware())

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// API v1 Routes
	userHandler := handlers.NewUserHandler(database)
	settingsHandler := handlers.NewSettingsHandler(database)

	api := router.Group("/api/v1")
	{
		// User Routes (Authentication/Registration)
		users := api.Group("/users")
		{
			users.POST("", userHandler.RegisterUser)
			users.GET("/:id", userHandler.GetUserByID)
			users.POST("login", userHandler.LoginUser)
			users.POST("register", userHandler.RegisterUser)
		}
	}

	secure := api.Group("/")
	secure.Use(middleware.JWTAuthMiddleware())
	{
		settings := secure.Group("/settings")
		{
			settings.GET("", settingsHandler.GetAllSettings)
			settings.GET("/:name", settingsHandler.GetSettingByName)
			settings.PUT("/:name", settingsHandler.UpdateSetting)
			settings.DELETE("/:id", settingsHandler.DeleteSetting)
		}

		// weitere geschützte Routen z.B. Collections, Items, etc.
		// secure.GET("/collection", collectionHandler.List)
	}

	// Server starten
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server startet auf http://%s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
