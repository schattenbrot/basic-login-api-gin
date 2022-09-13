package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/schattenbrot/basic-login-api-gin/internal/models"
	"github.com/schattenbrot/basic-login-api-gin/pkg/utils"
)

func (m *Repository) ParseAuthCookie(c *gin.Context) {
	token, err := c.Cookie(m.App.Config.Cookie.Name)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err, http.StatusUnauthorized)
		return
	}
	claims := models.CustomClaims{}
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return m.App.Config.Cookie.Secret, nil
	})
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err, http.StatusUnauthorized)
		return
	}

	c.Set("authCookie", claims)

	c.Next()
}
