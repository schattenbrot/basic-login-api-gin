package models

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	TokenVersion int64 `json:"tokenVersion"`
	jwt.StandardClaims
}
