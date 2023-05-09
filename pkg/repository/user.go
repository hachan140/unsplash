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
	FindUserByPhoneNumber(ctx context.Context, phone_number string) (*model.User, error)
	DeleteUserByUsername(ctx context.Context, username string) error
	ListUsersByUsernameAndPhoneNumber(ctx context.Context, page int, limit int, username string, phoneNumber string) ([]*model.User, error)
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
func (u *userRepository) FindUserByPhoneNumber(ctx context.Context, phone_number string) (*model.User, error) {
	var user *model.User
	//cho err co p bad request kh a????alooooo
	if err := u.db.WithContext(ctx).Where("phone_number = ?", phone_number).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) ListUsersByUsernameAndPhoneNumber(ctx context.Context, page int, limit int, username string, phoneNumber string) ([]*model.User, error) {
	var users []*model.User
	offset := (page - 1) * limit
	if err := u.db.WithContext(ctx).Offset(offset).Limit(limit).
		Where("username like ?", username+"%").
		Where("phone_number like ?", phoneNumber+"%").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepository) DeleteUserByUsername(ctx context.Context, username string) error {

	if err := u.db.WithContext(ctx).Model(model.User{}).Where("username = ?", username).Delete(model.User{}).Error; err != nil {
		return err
	}
	return nil
}
