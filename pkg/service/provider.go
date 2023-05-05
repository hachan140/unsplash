package service

import (
	"gin_unsplash/pkg/adapter/unsplash"
	"gin_unsplash/pkg/repository"
)

type provider struct {
	photoService PhotoService
}
type Provider interface {
	PhotoService() PhotoService
}

func NewProvider(repoProvider repository.Provider, unsplashAdapter unsplash.Adapter) Provider {
	return &provider{
		photoService: NewPhotoService(repoProvider, unsplashAdapter),
	}
}
func (p *provider) PhotoService() PhotoService {
	return p.photoService
}
