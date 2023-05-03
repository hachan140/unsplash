package main

import (
	"fmt"
	"gin_unsplash/adapter"
	_ "gin_unsplash/adapter"
	"gin_unsplash/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var unsplashConfig config.UnsplashConfig
	if err := envconfig.Process("", &unsplashConfig); err != nil {
		panic(err)
	}

	unsplashAdapter, err := adapter.NewAdapter(unsplashConfig.APIKey)
	if err != nil {
		panic(err)
	}

	listImgReq := adapter.ListImageRequest{
		Page:    1,
		PerPage: 5,
		OrderBy: "popular"}
	images, err := unsplashAdapter.ListImages(listImgReq)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----List 5 images order by popular-----")
	for i, img := range images {
		fmt.Println("image", i+1)
		fmt.Println(img.AltDescription)
	}

	imageById, err := unsplashAdapter.GetImageById("Lc63Mk1BN2s")
	if err != nil {
		panic(err)
	}
	fmt.Println("-----Get photo by ID-------")
	fmt.Println(imageById.AltDescription)

	imageRandom, err := unsplashAdapter.GetRandomImage()
	if err != nil {
		panic(err)
	}
	fmt.Println("-----A random pic-----")
	fmt.Println(imageRandom.AltDescription)

}
