package service

import (
	"context"
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/httperror"
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
	ListUsersByUsernameAndPhoneNumber(ctx context.Context, req dto.ListUsersByUsernameAndPhoneNumberRequest) (*dto.ListUserByUsernameAndPhoneNumberResponse, error)
	DeleteUserByUsername(ctx context.Context, req dto.DeleteUserByUsernameRequest) (*dto.DeleteUserByUsernameResponse, error)
}

func NewUserService(repoProvider repository.Provider) UserService {
	return &userService{
		userRepo: repoProvider.UserRepository(),
	}
}
func (u *userService) validateDuplicateUsername(ctx context.Context, username string) error {
	_, err := u.userRepo.FindUserByUsername(ctx, username)
	if err == nil {
		return httperror.BadRequest("duplicate username")
	}
	if err != gorm.ErrRecordNotFound {
		return nil
	}
	return nil
}

func (u *userService) validatePhoneNumber(ctx context.Context, phone_number string) error {
	_, err := u.userRepo.FindUserByPhoneNumber(ctx, phone_number)
	if err == nil {
		return httperror.BadRequest("duplicate phone number")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (u *userService) DeleteUserByUsername(ctx context.Context, req dto.DeleteUserByUsernameRequest) (*dto.DeleteUserByUsernameResponse, error) {
	_, err := u.userRepo.FindUserByUsername(ctx, req.Username)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, httperror.BadRequest("record not found")
		}
		return nil, err
	}
	if err := u.userRepo.DeleteUserByUsername(ctx, req.Username); err != nil {
		return nil, err
	}
	return &dto.DeleteUserByUsernameResponse{Message: "delete success"}, nil

}

func (u *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	if err := u.validateDuplicateUsername(ctx, req.Username); err != nil {
		return nil, err
	}
	if err := u.validatePhoneNumber(ctx, req.PhoneNumber); err != nil {
		return nil, err
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

func (u *userService) ListUsersByUsernameAndPhoneNumber(ctx context.Context, req dto.ListUsersByUsernameAndPhoneNumberRequest) (*dto.ListUserByUsernameAndPhoneNumberResponse, error) {
	users, err := u.userRepo.ListUsersByUsernameAndPhoneNumber(ctx, req.Page, req.Limit, req.Username, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	res := &dto.ListUserByUsernameAndPhoneNumberResponse{
		Data:    mapper.UsersToDTOs(users),
		Message: "success",
	}
	return res, nil
}
