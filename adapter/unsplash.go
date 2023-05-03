package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Adapter interface {
	ListImages(req ListImageRequest) ([]Image, error)
	GetRandomImage() (*Image, error)
	GetImageById(id string) (*Image, error)
}
type adapter struct {
	apiKey string
}

func NewAdapter(apiKey string) (Adapter, error) {
	adt := &adapter{
		apiKey: apiKey,
	}
	return adt, nil
}

func (adt adapter) ListImages(req ListImageRequest) ([]Image, error) {

	url := fmt.Sprintf("https://api.unsplash.com/photos?client_id=%v&page=%v&per_page=%v&order_by=%v", adt.apiKey, req.Page, req.PerPage, req.OrderBy)
	res, err := http.Get(url)

	if err != nil {

		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {

		return nil, err
	}
	var images = make([]Image, 0)
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (adt adapter) GetRandomImage() (*Image, error) {
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%v", adt.apiKey)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	var img Image
	if err := json.Unmarshal(body, &img); err != nil {
		return nil, err
	}
	return &img, nil
}
func (adt adapter) GetImageById(id string) (*Image, error) {
	url := fmt.Sprintf("https://api.unsplash.com/photos/%v?client_id=%v", id, adt.apiKey)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	var img Image
	if err := json.Unmarshal(body, &img); err != nil {
		return nil, err
	}
	return &img, nil

}
