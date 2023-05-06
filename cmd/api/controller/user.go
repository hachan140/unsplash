package controller

import "gin_unsplash/pkg/service"

type userController struct {
	userService service.UserService
}
type UserController interface {
}

func NewUserController(serviceProvider service.Provider) UserController {
	return &userController{
		userService: serviceProvider.UserService(),
	}
}
