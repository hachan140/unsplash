package repository

import (
	"context"
	"gin_unsplash/pkg/model"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Insert(ctx context.Context, data *model.Photo) error
	FindOneByID(ctx context.Context, id string) *model.Photo
	FindAllPhotos(ctx context.Context, page int, limit int) []model.Photo
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

func (p *photoRepository) FindOneByID(ctx context.Context, id string) *model.Photo {
	var photo model.Photo
	p.db.WithContext(ctx).First(&photo, "id = ?", id)
	return &photo
}

func (p *photoRepository) FindAllPhotos(ctx context.Context, page int, limit int) []model.Photo {
	offset := (page-1)*limit + 1
	var photos []model.Photo
	p.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&photos)
	return photos

}
