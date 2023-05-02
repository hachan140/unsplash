package main

import (
	"fmt"
	"gin_unsplash/adapter"
	_ "gin_unsplash/adapter"
)

func main() {
	var listImgReq = adapter.ListImageRequest{1, 5, "popular"}

	apiKey := "eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8"
	unsplashAdapter, err := adapter.NewAdapter(apiKey)
	if err != nil {
		panic(err)
	}
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
