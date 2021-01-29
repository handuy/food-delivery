package foodstorage

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func NewFoodStorage(db *gorm.DB) *store {
	return &store{db: db}
}