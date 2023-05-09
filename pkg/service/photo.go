package service

import (
	"context"
	"fmt"
	"gin_unsplash/pkg/adapter/unsplash"
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/httperror"
	"gin_unsplash/pkg/mapper"
	"gin_unsplash/pkg/model"
	"gin_unsplash/pkg/repository"
	"gorm.io/gorm"
)

type PhotoService interface {
	ListPhotos(ctx context.Context, req dto.ListPhotosRequest) (*dto.ListPhotosResponse, error)
	FetchUnsplashPhotos(ctx context.Context, req dto.FetchUnsplashPhotosRequest) (*dto.FetchUnsplashPhotosResponse, error)
	DeletePhotoByID(ctx context.Context, req dto.DeletePhotoByIDRequest) (*dto.DeletePhotoByIDResponse, error)
}

type photoService struct {
	photoRepo repository.PhotoRepository

	unsplashAdapter unsplash.Adapter
}

func NewPhotoService(
	repoProvider repository.Provider,
	unsplashAdapter unsplash.Adapter,
) PhotoService {
	return &photoService{
		photoRepo:       repoProvider.PhotoRepository(),
		unsplashAdapter: unsplashAdapter,
	}
}

func (p *photoService) ListPhotos(ctx context.Context, req dto.ListPhotosRequest) (*dto.ListPhotosResponse, error) {
	photos, err := p.photoRepo.FindAllPhotos(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return &dto.ListPhotosResponse{
		Data:    mapper.PhotosToDTOs(photos),
		Message: "success",
	}, nil
}

func (p *photoService) FetchUnsplashPhotos(ctx context.Context, req dto.FetchUnsplashPhotosRequest) (*dto.FetchUnsplashPhotosResponse, error) {
	unsplashPhotos, err := p.unsplashAdapter.ListPhotos(unsplash.ListPhotoRequest{
		Page:    1,
		PerPage: req.Quantity,
		OrderBy: unsplash.PhotoOrderLatest,
	})
	if err != nil {
		return nil, err
	}
	insertedPhoto := make([]*model.Photo, 0)
	for _, unsplashPhoto := range unsplashPhotos {
		// check duplicate
		_, err := p.photoRepo.FindOneByID(ctx, unsplashPhoto.ID)
		// if httperror happen and httperror != record not found -> real httperror
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		// if err != nil -> duplicated
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
		insertedPhoto = append(insertedPhoto, photo)
	}
	message := fmt.Sprintf("%v new photos", len(insertedPhoto))
	return &dto.FetchUnsplashPhotosResponse{
		Data:    mapper.PhotosToDTOs(insertedPhoto),
		Message: message,
	}, nil

}

func (p *photoService) DeletePhotoByID(ctx context.Context, req dto.DeletePhotoByIDRequest) (*dto.DeletePhotoByIDResponse, error) {
	_, err := p.photoRepo.FindOneByID(ctx, req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, httperror.BadRequest("record not found")
		}
		return nil, err

	}
	if err := p.photoRepo.DeletePhotoByID(ctx, req.Id); err != nil {
		return nil, err
	}
	return &dto.DeletePhotoByIDResponse{Message: "delete success"}, nil
}
