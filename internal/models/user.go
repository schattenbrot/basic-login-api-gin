package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email             string             `json:"email,omitempty" bson:"email,omitempty"`
	GoogleAccessToken string             `json:"googleAccessToken,omitempty" bson:"googleAccessToken,omitempty"`
	Image             string             `json:"image,omitempty" bson:"image,omitempty"`
	Roles             []string           `json:"roles,omitempty" bson:"roles,omitempty"`
	Created           time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated           time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
}
