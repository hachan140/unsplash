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

	////list 5 imgs, order by popular
	//link1 := "https://api.unsplash.com/photos?client_id=eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8&page=1&per_page=5&order_by=popular"
	//res, err := http.Get(link1)
	//defer res.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	panic(err)
	//}
	//var images = make([]Image, 0)
	//err = json.Unmarshal([]byte(body), &images)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("-----List 5 images order by popular-----")
	//for i, img := range images {
	//	fmt.Println("image", i+1)
	//	fmt.Println(img.AltDescription)
	//}
	//
	////get a photo by id
	//id := "Lc63Mk1BN2s"
	//link2 := fmt.Sprintf("https://api.unsplash.com/photos/%v?client_id=eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8", id)
	//
	//res2, err := http.Get(link2)
	//defer res2.Body.Close()
	//
	//if err != nil {
	//	panic(err)
	//}
	//body2, err := io.ReadAll(res2.Body)
	//if err != nil {
	//	panic(err)
	//}
	//var img0 Image
	//err = json.Unmarshal([]byte(body2), &img0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("-----Get photo by ID-------")
	//fmt.Println(img0.AltDescription)
	//
	////get a random photo
	//link3 := "https://api.unsplash.com/photos/random?client_id=eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8"
	//res3, err := http.Get(link3)
	//defer res3.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//body3, err := io.ReadAll(res3.Body)
	//if err != nil {
	//	panic(err)
	//}
	//var imgRan Image
	//err = json.Unmarshal([]byte(body3), &imgRan)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("-----A random pic-----")
	//fmt.Println(imgRan.AltDescription)

}
