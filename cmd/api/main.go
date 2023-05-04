package main

import (
	"gin_unsplash/cmd/api/controller"
	"gin_unsplash/pkg/adapter/unsplash"
	"gin_unsplash/pkg/config"
	"gin_unsplash/pkg/connection"
	"gin_unsplash/pkg/model"
	"gin_unsplash/pkg/repository"
	"gin_unsplash/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	var mysqlConfig config.MySQL
	if err := envconfig.Process("", &mysqlConfig); err != nil {
		panic(err)
	}

	var unsplashConfig config.UnsplashConfig
	if err := envconfig.Process("", &unsplashConfig); err != nil {
		panic(err)
	}

	unsplashAdapter, err := unsplash.NewAdapter(unsplashConfig.APIKey)
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

	repoProvider := repository.NewProvider(db)

	photoService := service.NewPhotoService(repoProvider, unsplashAdapter)
	photoController := controller.NewPhotoController(photoService)

	route := gin.Default()

	route.GET("/api/photos", photoController.ListPhotos)
	route.POST("/api/photos/fetch-unsplash", photoController.FetchUnsplashPhotos)

	route.Run(":8080")

}
