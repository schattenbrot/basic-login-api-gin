package database

import (
	"context"
	"time"

	"github.com/schattenbrot/basic-login-api-gin/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepo interface {
	CreateUser(user models.User) (*string, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	// IncrementRefreshTokenVersion(id string) error
}

func (m *dbRepo) CreateUser(user models.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	oid := res.InsertedID.(primitive.ObjectID).Hex()

	return &oid, nil
}

func (m *dbRepo) GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	var user models.User

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, models.User{ID: oid}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *dbRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	var user models.User

	err := collection.FindOne(ctx, models.User{Email: email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
