package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/schattenbrot/basic-login-api-gin/internal/models"
	"github.com/schattenbrot/basic-login-api-gin/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) Register(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	user.Password = string(hashedPassword)
	user.Updated = time.Now()
	user.Created = time.Now()
	user.Roles = []string{"user"}

	oid, err := m.DB.CreateUser(user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err, http.StatusInternalServerError)
		return
	}

	type resp struct {
		ID string `json:"id"`
	}

	utils.WriteJSON(c, http.StatusCreated, resp{
		ID: *oid,
	})
}

func (m *Repository) Login(c *gin.Context) {
	var loginUser models.User

	err := c.BindJSON(&loginUser)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	// get user
	user, err := m.DB.GetUserByEmail(loginUser.Email)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err, http.StatusInternalServerError)
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	// create access token
	tokenString, expirationDate, err := m.createAccessToken(user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	// create cookie / refresh token
	refreshToken, err := m.createRefreshToken(user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}
	c.SetCookie(m.App.Config.Cookie.Name, refreshToken, 60*60*24*7, "/", "", m.App.Config.Cookie.Secure, m.App.Config.Cookie.Secure)

	// send cookie and access token
	type resp struct {
		ID             string   `json:"id"`
		Token          string   `json:"token"`
		ExpirationDate int64    `json:"exp"`
		Roles          []string `json:"roles"`
	}
	utils.WriteJSON(c, http.StatusOK, resp{
		ID:             user.ID.Hex(),
		Token:          tokenString,
		ExpirationDate: expirationDate,
		Roles:          user.Roles,
	})
}

func (m *Repository) Logout(c *gin.Context) {
	// set cookie expiration date in the past to remove it
	c.SetCookie(m.App.Config.Cookie.Name, "", -3600, "/", "", m.App.Config.Cookie.Secure, m.App.Config.Cookie.Secure)
}

func (m *Repository) RefreshAccessToken(c *gin.Context) {
	token := c.MustGet("authCookie").(models.CustomClaims)

	user, err := m.DB.GetUserById(token.Issuer)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	if user.TokenVersion != token.TokenVersion {
		err = errors.New("invalid token version")
		utils.ErrorJSON(c, err, http.StatusUnauthorized)
		m.App.Logger.Println(err)
		return
	}

	accessToken, expirationDate, err := m.createAccessToken(user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}

	refreshToken, err := m.createRefreshToken(user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}
	c.SetCookie(m.App.Config.Cookie.Name, refreshToken, 60*60*24*7, "/", "", m.App.Config.Cookie.Secure, m.App.Config.Cookie.Secure)

	type resp struct {
		Token          string `json:"token"`
		ExpirationDate int64  `json:"exp"`
	}
	utils.WriteJSON(c, http.StatusOK, resp{
		Token:          accessToken,
		ExpirationDate: expirationDate,
	})
}

func (m *Repository) RevokeRefreshAccessToken(c *gin.Context) {
	userID := c.MustGet("authCookie").(models.CustomClaims).StandardClaims.Issuer
	m.App.Logger.Println(userID)

	// revoke token by incrementing it's version
	err := m.DB.IncrementRefreshTokenVersion(userID)
	if err != nil {
		m.App.Logger.Println(err)
		utils.ErrorJSON(c, err)
		return
	}
}

// -------------------------------------------------------------------------------------------------
// Auth utils
// -------------------------------------------------------------------------------------------------
func (m *Repository) createAccessToken(user *models.User) (string, int64, error) {
	expirationDate := time.Now().Add(15 * time.Minute) // 15 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID.Hex(),
		ExpiresAt: expirationDate.Unix(),
	})
	accessToken, err := token.SignedString(m.App.Config.JWT)
	if err != nil {
		return "", 0, err
	}
	return accessToken, expirationDate.Unix(), err
}

func (m *Repository) createRefreshToken(user *models.User) (string, error) {
	claims := models.CustomClaims{
		TokenVersion: user.TokenVersion,
		StandardClaims: jwt.StandardClaims{
			Issuer: user.ID.Hex(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString(m.App.Config.Cookie.Secret)
	if err != nil {
		return "", err
	}
	return refreshToken, err
}
