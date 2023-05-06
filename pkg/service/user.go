package service

import "gin_unsplash/pkg/repository"

type userService struct {
	userRepo repository.UserRepository
}
type UserService interface {
}

func NewUserService(repoProvider repository.Provider) UserService {
	return &userService{
		userRepo: repoProvider.UserRepository(),
	}
}
