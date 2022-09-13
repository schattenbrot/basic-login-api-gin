package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/schattenbrot/basic-login-api-gin/pkg/utils"
)

func (m *Repository) IsAuthenticated(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		m.App.Logger.Println("authorization header not found")
		utils.ErrorJSON(c, errors.New("authorization header not found"), http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, " ")[1] // bearer token

	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return m.App.Config.JWT, nil
	})
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err, http.StatusUnauthorized)
		return
	}

	c.Set("claims", claims)

	c.Next()
}
