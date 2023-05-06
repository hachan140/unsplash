package repository

import (
	"context"
	"gin_unsplash/pkg/model"
	"gorm.io/gorm"
)

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
func (u *userRepository) Insert(ctx context.Context, data *model.User) error {
	if err := u.db.WithContext(ctx).Model(data).Create(data).Error; err != nil {
		return err
	}
	return nil
}
