package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/schattenbrot/basic-login-api-gin/internal/config"
	"github.com/schattenbrot/basic-login-api-gin/internal/controllers"
	"github.com/schattenbrot/basic-login-api-gin/internal/middlewares"
)

func Routes(app config.AppConfig) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     app.Config.Cors,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", controllers.Repo.StatusHandler)

	r.POST("/register", controllers.Repo.Register)
	r.POST("/login", controllers.Repo.Login)
	r.GET("/logout", controllers.Repo.Logout)
	r.POST("/refresh-token", middlewares.Repo.ParseAuthCookie, controllers.Repo.RefreshAccessToken)
	r.POST("/revoke-token", middlewares.Repo.ParseAuthCookie, controllers.Repo.RevokeRefreshAccessToken)
	r.GET("/users/own", middlewares.Repo.IsAuthenticated, controllers.Repo.GetOwnUser)

	return r
}
