package controller

import "gin_unsplash/pkg/service"

type provider struct {
	photoController PhotoController
	userController  UserController
}
type Provider interface {
	PhotoController() PhotoController
	UserController() UserController
}

func NewProvider(serviceProvider service.Provider) Provider {
	return &provider{
		photoController: NewPhotoController(serviceProvider),
		userController:  NewUserController(serviceProvider),
	}
}
func (p *provider) PhotoController() PhotoController {
	return p.photoController
}
func (p *provider) UserController() UserController {
	return p.UserController()
}
