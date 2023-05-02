package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Adapter interface {
	ListImages(req ListImageRequest) ([]Image, error)
}

func NewAdapter(apiKey string) (Adapter, error) {
	adt := &adapter{
		apiKey: apiKey,
	}
	return adt, nil
}

type ListImageRequest struct {
	Page    int
	PerPage int
	OrderBy string
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
