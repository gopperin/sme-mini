package router

import (
	myapi "github.com/Eric-GreenComb/contrib/api"
	"github.com/gin-gonic/gin"

	"gateway/config"
	"gateway/handler"
)

// SetupKYCRouter SetupKYCRouter
func SetupKYCRouter(g *gin.Engine) {
	r := g.Group("/api/v1/user")
	{
	}
	r.Use(myapi.SignedAuth(config.Server.APIAppendKey, "", config.Server.APIMd5Key, nil))
	{

		// 绑定用户
		r.POST("/", handler.CreateUser)

		// 获取用户
		r.POST("/info", handler.GetUser)

	}
}
