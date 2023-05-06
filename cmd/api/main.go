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
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

/*
TODO:
taọ provider cho service, controller, config


tạo 1 model user gồm id, username, password, fullname, phone_number, created_at, updated_at
tạo repository tương ứng gôm Insert, FindByID, FindByUsername, FindMany
tạo service và controller tương ứng gồm các function: CreateUser ( username, phone_number unique ), ListUser search like username

*/

func main() {
	configProvider := config.NewConfigProvider()
	unsplashAdapter, err := unsplash.NewAdapter(configProvider.UnsplashConfig().APIKey)
	if err != nil {
		panic(err)
	}
	db, err := connection.NewMySQLConnection(configProvider.MySQL())
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(model.Photo{}); err != nil {
		panic(err)
	}
	repoProvider := repository.NewProvider(db)
	serviceProvider := service.NewProvider(repoProvider, unsplashAdapter)
	controllerProvider := controller.NewProvider(serviceProvider)

	route := gin.Default()
	route.POST("/api/user", controllerProvider.UserController().CreateUser)
	route.GET("/api/photos", controllerProvider.PhotoController().ListPhotos)
	route.POST("/api/photos/fetch-unsplash", controllerProvider.PhotoController().FetchUnsplashPhotos)
	route.Run(":8080")

}
