package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (m *Repository) GetOwnUser(c *gin.Context) {
	// get issuer from claims
	userID := c.MustGet("claims").(jwt.StandardClaims).Issuer
	m.App.Logger.Println(userID)

}
