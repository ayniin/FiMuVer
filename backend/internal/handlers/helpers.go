package handlers

import (
    "fmt"

    "github.com/gin-gonic/gin"
)

// GetUserIDFromContext liest die user_id (gesetzt von der JWT-Middleware) und gibt sie als uint zurück.
func GetUserIDFromContext(c *gin.Context) (uint, error) {
    v, ok := c.Get("user_id")
    if !ok {
        return 0, fmt.Errorf("user_id not found in context")
    }
    switch id := v.(type) {
    case uint:
        return id, nil
    case int:
        return uint(id), nil
    case int64:
        return uint(id), nil
    case float64:
        return uint(id), nil
    default:
        return 0, fmt.Errorf("invalid user_id type: %T", v)
    }
}

