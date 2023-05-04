package repository

import (
	"context"
	"gin_unsplash/pkg/model"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Insert(ctx context.Context, data *model.Photo) error
	FindOneByID(ctx context.Context, id string) (*model.Photo, error)
	FindAllPhotos(ctx context.Context, page int, limit int) ([]*model.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (p *photoRepository) Insert(ctx context.Context, data *model.Photo) error {
	if err := p.db.WithContext(ctx).Model(data).Create(data).Error; err != nil {
		return err
	}
	return nil
}

// #TODO: implement FindOneByID and FindMany function for repository
// FindMany param: page, limit

func (p *photoRepository) FindOneByID(ctx context.Context, id string) (*model.Photo, error) {
	var photo model.Photo
	if err := p.db.WithContext(ctx).Where("id=?", id).First(&photo).Error; err != nil {
		return nil, err
	}
	return &photo, nil
}

func (p *photoRepository) FindAllPhotos(ctx context.Context, page int, limit int) ([]*model.Photo, error) {
	offset := (page - 1) * limit
	var photos []*model.Photo
	if err := p.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}
