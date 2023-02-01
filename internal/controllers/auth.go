package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/schattenbrot/basic-login-api-gin/internal/models"
	"github.com/schattenbrot/basic-login-api-gin/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Repository) Login(c *gin.Context) {
	type LoginUser struct {
		Email string `json:"email"`
		Image string `json:"image"`
		Token string `json:"token"`
	}
	var loginUser LoginUser

	err := c.BindJSON(&loginUser)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	// get user
	var user *models.User
	user, err = m.DB.GetUserByEmail(loginUser.Email)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			m.App.Logger.Println(err)
			utils.ErrorJSON(c, err, http.StatusInternalServerError)
			return
		}

		// if user does not exist, create user
		user = &models.User{
			Email:             loginUser.Email,
			Image:             loginUser.Image,
			GoogleAccessToken: loginUser.Token,
			Roles:             []string{"user"},
			Created:           time.Now(),
			Updated:           time.Now(),
		}

		oid, err := m.DB.CreateUser(*user)
		if err != nil {
			m.App.Logger.Println(err)
			utils.ErrorJSON(c, err, http.StatusInternalServerError)
			return
		}

		user.ID, err = primitive.ObjectIDFromHex(*oid)
		if err != nil {
			m.App.Logger.Println(err)
			utils.ErrorJSON(c, err, http.StatusInternalServerError)
			return
		}
	}

	// create access token
	tokenString, expirationDate, err := m.createAccessToken(user.ID)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	// send cookie and access token
	type resp struct {
		ID             string   `json:"id"`
		AccessToken    string   `json:"accessToken"`
		ExpirationDate int64    `json:"exp"`
		Roles          []string `json:"roles"`
	}
	utils.WriteJSON(c, http.StatusOK, resp{
		ID:             user.ID.Hex(),
		AccessToken:    tokenString,
		ExpirationDate: expirationDate,
		Roles:          user.Roles,
	})
}

// ---------------------------------------------------------------------------------------------------------------------
// Auth utils
// ---------------------------------------------------------------------------------------------------------------------
func (m *Repository) createAccessToken(userId primitive.ObjectID) (string, int64, error) {
	expirationDate := time.Now().AddDate(0, 1, 0).UnixMilli() // add one month
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    userId.Hex(),
		ExpiresAt: expirationDate,
	})
	accessToken, err := token.SignedString(m.App.Config.JWT)
	if err != nil {
		return "", 0, err
	}
	return accessToken, expirationDate, err
}
