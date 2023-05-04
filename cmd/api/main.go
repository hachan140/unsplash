package main

import (
	"context"
	"fmt"
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

	//unsplashAdapter, err := adapter.NewAdapter(unsplashConfig.APIKey)
	//if err != nil {
	//	panic(err)
	//}

	db, err := connection.NewMySQLConnection(mysqlConfig)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(model.Photo{}); err != nil {
		panic(err)
	}

	photoRepo := repository.NewPhotoRepository(db)

	//for i := 1; i <= 10; i++ {
	//	photo_, err := unsplashAdapter.GetRandomPhoto()
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//	photoModel_ := &model.Photo{
	//		ID:             photo_.ID,
	//		CreatedAt:      photo_.CreatedAt,
	//		UpdatedAt:      photo_.UpdatedAt,
	//		Width:          photo_.Width,
	//		Height:         photo_.Height,
	//		Url:            photo_.Urls.Raw,
	//		Description:    photo_.Description,
	//		AltDescription: photo_.AltDescription,
	//		Likes:          photo_.Likes,
	//	}
	//
	//	if err := photoRepo.Insert(ctx, photoModel_); err != nil {
	//		panic(err)
	//	}
	//}

	fmt.Println("---Get one by ID---")
	photoByID, err := photoRepo.FindOneByID(ctx, "-HprBtc9dWY")
	if err != nil {
		panic(err)
	}
	fmt.Println(photoByID.Url)

	// #TODO: get all photo in db and print its url
	allPhotos, err := photoRepo.FindAllPhotos(ctx, 2, 5)
	if err != nil {
		panic(err)

	}
	fmt.Println("---All Photo Url---")
	for _, photo := range allPhotos {
		fmt.Println(photo.AltDescription)
	}
}
