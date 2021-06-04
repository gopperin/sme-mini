package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenExpired 一些常量
var (
	ErrTokenExpired     error  = errors.New("Token is expired")
	ErrTokenNotValidYet error  = errors.New("Token not active yet")
	ErrTokenMalformed   error  = errors.New("That's not even a token")
	ErrTokenInvalid     error  = errors.New("Couldn't handle this token")
	ErrSignKey          string = "newtrekWang"
)

// JWT JWT
type JWT struct {
	SigningKey []byte
}

// CustomClaims CustomClaims
type CustomClaims struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// CreateToken 创建token
func (j *JWT) CreateToken(kid string, claims CustomClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = kid
	token.Claims = claims
	res, err := token.SignedString(j.SigningKey)
	return res, err
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			log.Panicln("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// VerifyToken VerifyToken
func (j *JWT) VerifyToken(tokenString, hmacKey string) error {
	parts := strings.Split(tokenString, ".")
	method := jwt.GetSigningMethod("HS256")
	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], []byte(hmacKey))
	if err != nil {
		return err
	}
	return nil
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(8 * time.Hour).Unix()
		return j.CreateToken("sim2", *claims)
	}

	return "", ErrTokenInvalid
}
