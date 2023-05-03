package main

import (
	"context"
	"gin_unsplash/pkg/adapter"
	"gin_unsplash/pkg/config"
	"gin_unsplash/pkg/connection"
	"gin_unsplash/pkg/model"
	"gin_unsplash/pkg/repository"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	var unsplashConfig config.UnsplashConfig
	if err := envconfig.Process("", &unsplashConfig); err != nil {
		panic(err)
	}
	var mysqlConfig config.MySQL
	if err := envconfig.Process("", &mysqlConfig); err != nil {
		panic(err)
	}

	unsplashAdapter, err := adapter.NewAdapter(unsplashConfig.APIKey)
	if err != nil {
		panic(err)
	}

	db, err := connection.NewMySQLConnection(mysqlConfig)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(model.Photo{}); err != nil {
		panic(err)
	}

	photoRepo := repository.NewPhotoRepository(db)

	photo, err := unsplashAdapter.GetRandomPhoto()
	if err != nil {
		panic(err)
	}

	photoModel := &model.Photo{
		ID:             photo.ID,
		CreatedAt:      photo.CreatedAt,
		UpdatedAt:      photo.UpdatedAt,
		Width:          photo.Width,
		Height:         photo.Height,
		Url:            photo.Urls.Raw,
		Description:    photo.Description,
		AltDescription: photo.AltDescription,
		Likes:          photo.Likes,
	}

	if err := photoRepo.Insert(ctx, photoModel); err != nil {
		panic(err)
	}

	// #TODO: get all photo in db and print its url
}
