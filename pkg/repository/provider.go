package repository

import "gorm.io/gorm"

type Provider interface {
	PhotoRepository() PhotoRepository
}

type provider struct {
	photoRepo PhotoRepository
}

func NewProvider(db *gorm.DB) Provider {
	return &provider{
		photoRepo: NewPhotoRepository(db),
	}
}

func (p *provider) PhotoRepository() PhotoRepository {
	return p.photoRepo
}
