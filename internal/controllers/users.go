package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/schattenbrot/basic-login-api-gin/pkg/utils"
)

func (m *Repository) GetOwnUser(c *gin.Context) {
	// get issuer from claims
	userID := c.MustGet("claims").(jwt.StandardClaims).Issuer

	user, err := m.DB.GetUserById(userID)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(c, http.StatusOK, user)
}
