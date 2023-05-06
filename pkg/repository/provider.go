package repository

import "gorm.io/gorm"

type Provider interface {
	PhotoRepository() PhotoRepository
	UserRepository() UserRepository
}

type provider struct {
	photoRepo      PhotoRepository
	userRepository UserRepository
}

func NewProvider(db *gorm.DB) Provider {
	return &provider{
		photoRepo:      NewPhotoRepository(db),
		userRepository: NewUserRepo(db),
	}
}

func (p *provider) PhotoRepository() PhotoRepository {
	return p.photoRepo
}
func (p *provider) UserRepository() UserRepository {
	return p.userRepository
}
