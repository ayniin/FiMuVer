package main

import (
	"fmt"
	"log"

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

	// Initialisiere Datenbank
	database, err := db.InitializeDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Fehler beim Initialisieren der Datenbank: %v", err)
	}
	defer database.Close()

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
	mediaHandler := handlers.NewMediaHandler(database)
	collectionHandler := handlers.NewCollectionHandler(database)
	collectionItemHandler := handlers.NewCollectionItemHandler(database)
	referenceHandler := handlers.NewReferenceHandler(database)
	userHandler := handlers.NewUserHandler(database)
	settingsHandler := handlers.NewSettingsHandler(database)

	api := router.Group("/api/v1")
	{
		// User Routes (Authentication/Registration)
		users := api.Group("/users")
		{
			users.POST("", userHandler.AddUser)
			users.GET("/:id", userHandler.GetUserByID)
		}

		// Settings Routes
		settings := api.Group("/settings")
		{
			settings.GET("", settingsHandler.GetAllSettings)
			settings.GET("/:name", settingsHandler.GetSettingByName)
			settings.PUT("/:name", settingsHandler.UpdateSetting)
			settings.DELETE("/:id", settingsHandler.DeleteSetting)
		}
		// Media Routes
		media := api.Group("/media")
		{
			media.GET("", mediaHandler.GetAllMedia)
			media.GET("/:id", mediaHandler.GetMediaByID)
			media.POST("", mediaHandler.CreateMedia)
			media.PUT("/:id", mediaHandler.UpdateMedia)
			media.DELETE("/:id", mediaHandler.DeleteMedia)
		}
		api.GET("/search", mediaHandler.SearchMedia)

		// Collection Routes
		collections := api.Group("/collections")
		{
			collections.GET("", collectionHandler.GetCollectionsByUser)
			collections.POST("", collectionHandler.CreateCollection)
			collections.GET("/:id", collectionHandler.GetCollectionByID)
			collections.PUT("/:id", collectionHandler.UpdateCollection)
			collections.DELETE("/:id", collectionHandler.DeleteCollection)

			// Collection Items (nested)
			collections.GET("/:collectionId/items", collectionItemHandler.GetCollectionItems)
			collections.POST("/:collectionId/items", collectionItemHandler.CreateCollectionItem)
			collections.PUT("/:collectionId/items/:itemId", collectionItemHandler.UpdateCollectionItem)
			collections.DELETE("/:collectionId/items/:itemId", collectionItemHandler.DeleteCollectionItem)
		}

		// Reference Tables Routes
		references := api.Group("/references")
		{
			// Genres
			references.GET("/genres", referenceHandler.GetAllGenres)
			references.POST("/genres", referenceHandler.CreateGenre)

			// Media Types
			references.GET("/media-types", referenceHandler.GetAllMediaTypes)
			references.POST("/media-types", referenceHandler.CreateMediaType)

			// Conditions
			references.GET("/conditions", referenceHandler.GetAllConditions)
			references.POST("/conditions", referenceHandler.CreateCondition)

			// Editions
			references.GET("/editions", referenceHandler.GetAllEditions)
			references.POST("/editions", referenceHandler.CreateEdition)

			// Labels
			references.GET("/labels", referenceHandler.GetAllLabels)
			references.POST("/labels", referenceHandler.CreateLabel)
		}
	}

	// Server starten
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server startet auf http://%s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
