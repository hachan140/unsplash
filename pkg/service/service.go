package service

import (
	"context"
	"fmt"
	"gin_unsplash/pkg/adapter/unsplash"
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/mapper"
	"gin_unsplash/pkg/model"
	"gin_unsplash/pkg/repository"
	"gorm.io/gorm"
)

type photoService struct {
	photoRepo       repository.PhotoRepository
	unsplashAdapter unsplash.Adapter
}
type PhotoService interface {
	ListPhotos(ctx context.Context, req dto.ListPhotosRequest) (*dto.ListPhotosResponse, error)
	FetchUnsplashPhotos(ctx context.Context, req dto.FetchUnsplashPhotoRequest) (*dto.FetchUnsplashPhotoResponse, error)
}

func NewPhotoService(photoRepo repository.PhotoRepository, unsplashAdapter unsplash.Adapter) PhotoService {
	return &photoService{
		photoRepo:       photoRepo,
		unsplashAdapter: unsplashAdapter,
	}
}

func (p *photoService) ListPhotos(ctx context.Context, req dto.ListPhotosRequest) (*dto.ListPhotosResponse, error) {
	photos, err := p.photoRepo.FindAllPhotos(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return &dto.ListPhotosResponse{
		Data:    mapper.PhotosDTOs(photos),
		Message: "success",
	}, nil
}

func (p *photoService) FetchUnsplashPhotos(ctx context.Context, req dto.FetchUnsplashPhotoRequest) (*dto.FetchUnsplashPhotoResponse, error) {
	unsplashPhotos, err := p.unsplashAdapter.ListPhotos(unsplash.ListPhotoRequest{
		Page:    1,
		PerPage: req.Quantity,
		OrderBy: unsplash.PhotoOrderLatest,
	})
	if err != nil {
		return nil, err
	}
	insertedPhotos := make([]*model.Photo, 0)
	for _, unsplashPhoto := range unsplashPhotos {
		_, err := p.photoRepo.FindOneByID(ctx, unsplashPhoto.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if err == nil {
			continue
		}
		photo := &model.Photo{
			ID:             unsplashPhoto.ID,
			CreatedAt:      unsplashPhoto.CreatedAt,
			UpdatedAt:      unsplashPhoto.UpdatedAt,
			Width:          unsplashPhoto.Width,
			Height:         unsplashPhoto.Height,
			Url:            unsplashPhoto.Urls.Raw,
			Description:    unsplashPhoto.Description,
			AltDescription: unsplashPhoto.AltDescription,
			Likes:          unsplashPhoto.Likes,
		}
		if err := p.photoRepo.Insert(ctx, photo); err != nil {
			return nil, err
		}
		insertedPhotos = append(insertedPhotos, photo)
	}
	message := fmt.Sprintf("%v new photos", len(insertedPhotos))
	return &dto.FetchUnsplashPhotoResponse{
		Data:    mapper.PhotosDTOs(insertedPhotos),
		Message: message,
	}, nil
}
