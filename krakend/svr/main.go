package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		return
	})

	// Ping test
	r.POST("/ping", func(c *gin.Context) {
		fmt.Println("Origin:", c.Request.Header.Get("Origin"))
		c.JSON(http.StatusOK, gin.H{"status": "pong"})
	})

	// validate
	r.POST("/validate", func(c *gin.Context) {
		var _user User
		c.ShouldBindJSON(&_user)
		fmt.Println(_user)
		c.JSON(http.StatusOK, _user)
		return
	})

	// validate
	r.POST("/base/create", func(c *gin.Context) {
		var _user User
		c.ShouldBindJSON(&_user)
		fmt.Println(_user)
		c.JSON(http.StatusOK, _user)
		return
	})

	return r
}

// User User
type User struct {
	KID      string `json:"kid"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// JWK JWK
type JWK struct {
	KID  string `json:"kid"`
	Keys string `json:"keys"`
}

func main() {
	r := setupRouter()
	r.Run(":14000")
}
