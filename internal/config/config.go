package config

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Port   int
	Env    string
	Cors   []string
	Cookie struct {
		Name     string
		SameSite string
	}
	DB struct {
		DSN string
	}
	JWT []byte
}

type AppConfig struct {
	Version         string
	ServerStartTime time.Time
	Config          Config
	Logger          *log.Logger
}

func Init() AppConfig {
	var app AppConfig

	flag.StringVar(&app.Config.Env, "env", "dev", "the app environment")

	flag.IntVar(&app.Config.Port, "port", 4000, "the port")

	flag.StringVar(&app.Config.DB.DSN, "dsn", "mongodb://localhost:27017", "the db dsn")

	var jwt string
	flag.StringVar(&jwt, "jwt", "wonderfulsecretphrase", "the jwt token secret")
	app.Config.JWT = []byte(jwt)

	var cors string
	flag.StringVar(&cors, "cors", "http://* https://*", "the by cors allowed origins")
	app.Config.Cors = strings.Split(cors, " ")

	flag.StringVar(&app.Config.Cookie.Name, "cookieName", "basic-login-api-gin", "the name of the cookie")

	flag.StringVar(&app.Config.Cookie.SameSite, "cookieSameSite", "none", "same site policy of cookies")

	flag.Parse()

	app.ServerStartTime = time.Now()
	app.Version = "1.0.0"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app.Logger = logger

	return app
}
