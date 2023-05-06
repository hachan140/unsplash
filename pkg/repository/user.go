package repository

import (
	"context"
	"gin_unsplash/pkg/model"
	"gorm.io/gorm"
	"time"
)

type userRepository struct {
	db *gorm.DB
}
type UserRepository interface {
	Insert(ctx context.Context, data *model.User) error
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (u *userRepository) Insert(ctx context.Context, data *model.User) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	if err := u.db.WithContext(ctx).Model(data).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}
