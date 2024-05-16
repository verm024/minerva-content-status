package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func Initialize(db *gorm.DB) *Repository {
	repo := Repository{db}
	return &repo
}
