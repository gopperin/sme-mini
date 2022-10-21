package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gopperin/sme-mini/gateway/handler"
)

// SetupUserRouter SetupUserRouter
func SetupUserRouter(g *gin.Engine) {
	r := g.Group("/api/v1/user")
	{
		// 绑定用户
		r.POST("/", handler.CreateUser)

		// 获取用户
		r.DELETE("/", handler.DeleteUser)

		// 修改用户
		r.PUT("/", handler.PutUser)

		// 获取用户
		r.PATCH("/", handler.PatchUser)

		// 获取用户
		r.GET("/:uid", handler.GetUserByUID)

		// 获取用户列表
		r.GET("/", handler.ListUser)
	}
}
