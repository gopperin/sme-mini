package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

var cacheManager *cache.ChainCache

func setupRouter() *gin.Engine {

	// Initialize Ristretto cache and Redis client
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{NumCounters: 1000, MaxCost: 100, BufferItems: 64})
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	// Initialize stores
	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	redisStore := store.NewRedis(redisClient, &store.Options{})

	// Initialize chained cache
	cacheManager = cache.NewChain(
		cache.New(ristrettoStore),
		cache.New(redisStore),
	)

	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		return
	})

	// Ping test
	r.POST("/ping", func(c *gin.Context) {
		fmt.Println("Origin:", c.Request.Header.Get("Origin"))
		c.JSON(http.StatusOK, gin.H{"status": "pong"})
	})

	// jwk/symmetric.json
	r.GET("/jwk/:kid", func(c *gin.Context) {

		_kid := c.Param("kid")

		_keys, err := cacheManager.Get("gudp.jwks." + _kid)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"status": "nil"})
			return
		}

		var _map map[string]interface{}
		err = json.Unmarshal([]byte(_keys.(string)), &_map)
		if err != nil {
			fmt.Println("JsonToMap err: ", err)
		}
		c.JSON(http.StatusOK, _map)

		return
	})

	// jwk/symmetric.json
	r.POST("/jwk", func(c *gin.Context) {

		var _jwk JWK
		c.ShouldBindJSON(&_jwk)

		err = cacheManager.Set("gudp.jwks."+_jwk.KID, _jwk.Keys, &store.Options{})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": err.Error()})
			return
		}

		c.JSON(200, _jwk)

		return
	})

	// signin
	r.POST("/signin", func(c *gin.Context) {

		var _user User
		c.ShouldBindJSON(&_user)

		_jwt := JWT{}

		switch _user.KID {
		case "gopperin":
			_jwt.SigningKey = []byte("WkFRIXhzdzI")
		case "greencomb":
			_jwt.SigningKey = []byte("a11111")
		}

		_claims := CustomClaims{}
		_claims.Issuer = "gopper.in"
		_claims.UID = "13810167616"
		_claims.Name = "eric"
		_claims.IssuedAt = time.Now().Unix()
		_claims.ExpiresAt = time.Now().Add(8 * time.Hour).Unix()
		_claims.Subject = "subject"

		_ret, _ := _jwt.CreateToken(_user.KID, _claims)
		c.JSON(http.StatusOK, gin.H{"jwt": _ret})
		return
	})

	// validate
	r.POST("/validate", func(c *gin.Context) {
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

	_port := ""
	flag.StringVar(&_port, "p", "4000", "input dir.")
	flag.Parse()

	r := setupRouter()
	r.Run(":" + _port)
}
