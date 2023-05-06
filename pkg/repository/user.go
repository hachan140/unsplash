package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}
type UserRepository interface {
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
