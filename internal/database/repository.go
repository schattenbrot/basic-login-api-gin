package database

import (
	"context"
	"time"

	"github.com/schattenbrot/basic-login-api-gin/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseRepo is the interface for all repository functions
type DatabaseRepo interface{}

type mongoDBRepo struct {
	App *config.AppConfig
	DB  *mongo.Database
}

// NewMongoDBRepo is the function for returning a mongoDBRepo.
func NewDBRepo(app *config.AppConfig, conn *mongo.Database) DatabaseRepo {
	return &mongoDBRepo{
		App: app,
		DB:  conn,
	}
}

// openDB creates a new database connection and returns the Database
func OpenDB(app config.AppConfig) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(app.Config.DB.DSN))
	if err != nil {
		app.Logger.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		app.Logger.Fatal(err)
	}
	db := client.Database("basic-login-api-gin")

	return db
}
