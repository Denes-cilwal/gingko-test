package book

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) GetBooks() ([]string, error) {
	books := []string{}
	return books, r.db.Table("books").Select("name").Scan(&books).Error
}
