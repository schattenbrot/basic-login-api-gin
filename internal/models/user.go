package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type GoogleCredentials struct {
// 	ID    string `json:"id" bson:"id"`
// 	Email string `json:"email" bson:"email"`
// }

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	// GoogleCredentials GoogleCredentials  `json:"googleCredentials,omitempty" bson:"googleCredentials,omitempty"`
	Email   string    `json:"email,omitempty" bson:"email,omitempty"`
	Image   string    `json:"image,omitempty" bson:"image,omitempty"`
	Roles   []string  `json:"roles,omitempty" bson:"roles,omitempty"`
	Created time.Time `json:"created,omitempty" bson:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty" bson:"updated,omitempty"`
}
