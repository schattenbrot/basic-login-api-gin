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
		Secret []byte
		Name   string
		Secure bool
	}
	DB struct {
		DSN  string
		Name string
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

	flag.IntVar(&app.Config.Port, "port", 8080, "the port")

	flag.StringVar(&app.Config.DB.DSN, "dsn", "mongodb://localhost:27017", "the db dsn")
	flag.StringVar(&app.Config.DB.Name, "dbName", "basic-api-db", "the name of the used database")

	var jwt string
	flag.StringVar(&jwt, "jwt", "wonderfulsecretphrase", "the jwt token secret")

	var cors string
	flag.StringVar(&cors, "cors", "http://* https://*", "the by cors allowed origins")

	var cookieSecret string
	flag.StringVar(&cookieSecret, "cookieSecret", "verywellkeptsecretphrase", "the cookie token secret")
	flag.StringVar(&app.Config.Cookie.Name, "cookieName", "basic-login-api-gin", "the name of the cookie")
	flag.BoolVar(&app.Config.Cookie.Secure, "cookieSameSite", false, "same site policy of cookies")

	flag.Parse()
	app.Config.JWT = []byte(jwt)
	app.Config.Cors = strings.Split(cors, " ")
	app.Config.Cookie.Secret = []byte(cookieSecret)

	app.ServerStartTime = time.Now()
	app.Version = "1.0.0"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app.Logger = logger

	return app
}
