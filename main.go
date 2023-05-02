package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
}
type Image struct {
	AltDescription string `json:"alt_description"`
	Urls           Urls   `json:"urls"`
	Likes          int    `json:"likes"`
}

func main() {
	//list 5 imgs, order by popular
	link1 := "https://api.unsplash.com/photos?client_id=eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8&page=1&per_page=5&order_by=popular"
	res, err := http.Get(link1)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var images = make([]Image, 0)
	err = json.Unmarshal([]byte(body), &images)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----List 5 images order by popular-----")
	for i, img := range images {
		fmt.Println("image", i+1)
		fmt.Println(img.AltDescription)
	}

	//get a photo by id
	link2 := "https://api.unsplash.com/photos/:id?client_id=eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8"
	params := url.Values{}
	params.Add("id", "LBI7cgq3pbM")
	u, _ := url.Parse(link2)
	res2, err := http.Get(u.String())
	defer res2.Body.Close()

	if err != nil {
		panic(err)
	}
	body2, err := io.ReadAll(res2.Body)
	if err != nil {
		panic(err)
	}
	var img0 Image
	err = json.Unmarshal([]byte(body2), &img0)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----Get photo by ID-------")
	fmt.Println(img0)

	//get a random photo
	link3 := "https://api.unsplash.com/photos/random?client_id=eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8"
	res3, err := http.Get(link3)
	defer res3.Body.Close()
	if err != nil {
		panic(err)
	}
	body3, err := io.ReadAll(res3.Body)
	if err != nil {
		panic(err)
	}
	var imgRan Image
	err = json.Unmarshal([]byte(body3), &imgRan)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----A random pic-----")
	fmt.Println(imgRan.AltDescription)

}
