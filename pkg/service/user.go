package service

import (
	"context"
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/mapper"
	"gin_unsplash/pkg/model"
	"gin_unsplash/pkg/repository"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UserRepository
}
type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}

func NewUserService(repoProvider repository.Provider) UserService {
	return &userService{
		userRepo: repoProvider.UserRepository(),
	}
}

func (u *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	_, err := u.userRepo.FindUserByUsername(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == nil {
		return &dto.CreateUserResponse{
			Data:    nil,
			Message: "duplicate username",
		}, nil
	}
	user := &model.User{

		Username:    req.Username,
		Password:    req.Password,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}
	if err := u.userRepo.Insert(ctx, user); err != nil {
		return nil, err
	}
	res := &dto.CreateUserResponse{
		Data:    mapper.UserToDTO(user),
		Message: "success",
	}
	return res, nil
}
