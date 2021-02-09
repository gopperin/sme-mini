package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gateway/config"
)

// Index Index
func Index(c *gin.Context) {
	c.String(http.StatusOK, "index")
}

// Health Health
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

// Release Release
func Release(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"release": config.Server.Release, "version": config.Server.Version})
}
