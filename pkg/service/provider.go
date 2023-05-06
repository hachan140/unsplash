package service

import (
	"gin_unsplash/pkg/adapter/unsplash"
	"gin_unsplash/pkg/repository"
)

type provider struct {
	photoService PhotoService
	userService  UserService
}
type Provider interface {
	PhotoService() PhotoService
	UserService() UserService
}

func NewProvider(repoProvider repository.Provider, unsplashAdapter unsplash.Adapter) Provider {
	return &provider{
		photoService: NewPhotoService(repoProvider, unsplashAdapter),
		userService:  NewUserService(repoProvider),
	}
}
func (p *provider) PhotoService() PhotoService {
	return p.photoService
}
func (p *provider) UserService() UserService {
	return p.userService
}
