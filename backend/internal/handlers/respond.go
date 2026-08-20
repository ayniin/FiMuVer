package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fail(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"error": msg})
}

func badRequest(c *gin.Context, msg string)  { fail(c, http.StatusBadRequest, msg) }
func unauthorized(c *gin.Context)            { fail(c, http.StatusUnauthorized, msgNotAuthenticated) }
func forbidden(c *gin.Context)               { fail(c, http.StatusForbidden, msgForbidden) }
func notFound(c *gin.Context, msg string)    { fail(c, http.StatusNotFound, msg) }
func conflict(c *gin.Context, msg string)    { fail(c, http.StatusConflict, msg) }
func serverError(c *gin.Context, msg string) { fail(c, http.StatusInternalServerError, msg) }
func badGateway(c *gin.Context, msg string)  { fail(c, http.StatusBadGateway, msg) }

func ok(c *gin.Context, data any) { c.JSON(http.StatusOK, gin.H{"data": data}) }
func okMessage(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{"message": msg, "data": data})
}
func created(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusCreated, gin.H{"message": msg, "data": data})
}
