package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Adapter interface {
	ListPhotos(req ListPhotoRequest) ([]Photo, error)
	GetRandomPhoto() (*Photo, error)
	GetPhotoByID(id string) (*Photo, error)
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

func (adt adapter) ListPhotos(req ListPhotoRequest) ([]Photo, error) {

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
	var images = make([]Photo, 0)
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (adt adapter) GetRandomPhoto() (*Photo, error) {
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
	var img Photo
	if err := json.Unmarshal(body, &img); err != nil {
		return nil, err
	}
	return &img, nil
}
func (adt adapter) GetPhotoByID(id string) (*Photo, error) {
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
	var img Photo
	if err := json.Unmarshal(body, &img); err != nil {
		return nil, err
	}
	return &img, nil
}
