package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware aktiviert CORS mit Credentials-Support für React Frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Erlaube nur spezifische Origins (nicht "*)
		allowedOrigins := getAllowedOrigins()
		if isOriginAllowed(origin, allowedOrigins) {
			// WICHTIG: Mit credentials="include" darf nicht "*" sein!
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Gib erlaubte Origins zurück (basierend auf Environment)
func getAllowedOrigins() []string {
	originEnv := os.Getenv("CORS_ALLOWED_ORIGINS")
	if originEnv != "" {
		return strings.Split(originEnv, ",")
	}
	
	// Defaults für Development
	return []string{
		"http://localhost:5173",    // React Dev Server
		"http://localhost:3000",    // Alternative React Port
		"http://127.0.0.1:5173",
		"http://127.0.0.1:3000",
	}
}

// Prüfe ob Origin in der Whitelist ist
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}
