package database

import (
	"context"

	"github.com/schattenbrot/basic-login-api-gin/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setIndizes(app *config.AppConfig, db *mongo.Database) {
	coll := db.Collection("users")

	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		app.Logger.Fatalln(err)
	}
}
