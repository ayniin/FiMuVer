package main

import (
	"fmt"
	"log"

	"fimuver/internal/config"
	"fimuver/internal/db"
	"fimuver/internal/handlers"
	"fimuver/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "fimuver/docs"
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

	api := router.Group("/api/v1")
	{
		media := api.Group("/media")
		{
			media.GET("", mediaHandler.GetAllMedia)
			media.GET("/:id", mediaHandler.GetMediaByID)
			media.POST("", mediaHandler.CreateMedia)
			media.PUT("/:id", mediaHandler.UpdateMedia)
			media.DELETE("/:id", mediaHandler.DeleteMedia)
		}

		api.GET("/search", mediaHandler.SearchMedia)
	}

	// Server starten
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server startet auf http://%s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
