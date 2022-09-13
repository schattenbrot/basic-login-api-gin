package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	Roles        []string           `json:"roles,omitempty" bson:"roles,omitempty"`
	TokenVersion int64              `json:"tokenVersion,omitempty" bson:"tokenVersion,omitempty"`
	Created      time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated      time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
}
