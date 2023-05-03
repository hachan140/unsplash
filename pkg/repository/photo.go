package repository

import (
	"context"
	"gin_unsplash/pkg/model"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Insert(ctx context.Context, data *model.Photo) error
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
