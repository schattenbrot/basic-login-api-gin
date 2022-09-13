package main

import (
	"strconv"

	"github.com/schattenbrot/basic-login-api-gin/internal/config"
	"github.com/schattenbrot/basic-login-api-gin/internal/controllers"
	"github.com/schattenbrot/basic-login-api-gin/internal/database"
	"github.com/schattenbrot/basic-login-api-gin/internal/middlewares"
	"github.com/schattenbrot/basic-login-api-gin/internal/routes"
)

func main() {
	app := config.Init()

	db := database.OpenDB(app)
	dbRepo := database.NewDBRepo(&app, db)

	controllerRepo := controllers.NewRepo(&app, dbRepo)
	controllers.NewHandlers(controllerRepo)

	middlewaresRepo := middlewares.NewRepo(&app, dbRepo)
	middlewares.NewMiddlewares(middlewaresRepo)

	r := routes.Routes(app)
	r.Run(":" + strconv.Itoa(app.Config.Port))
}
