package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gopperin/sme-mini/gateway/handler"
)

// SetupBaseRouter SetupBaseRouter
func SetupBaseRouter(g *gin.Engine) {
	r0 := g.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)
		r0.GET("release", handler.Release)
	}
}
