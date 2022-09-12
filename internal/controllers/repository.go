package controllers

import (
	"github.com/schattenbrot/basic-login-api-gin/internal/config"
	"github.com/schattenbrot/basic-login-api-gin/internal/database"
)

// Repository represents the handler repository to share the app configuration.
type Repository struct {
	App *config.AppConfig
	DB  database.DatabaseRepo
}

// Repo is the repository to share the app configuration.
var Repo *Repository

// NewRepo returns a new instance of a repository for the mongo driver.
func NewRepo(a *config.AppConfig, db database.DatabaseRepo) *Repository {
	return &Repository{
		App: a,
		DB:  db,
	}
}

// NewHandlers sets the handler repository.
func NewHandlers(r *Repository) {
	Repo = r
}
