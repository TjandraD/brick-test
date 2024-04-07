package repo

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
