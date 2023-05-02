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

func NewAdapter(apiKey string) (*adapter, error) {
	adt := &adapter{
		apiKey: apiKey,
	}
	return adt, nil
}

func (adt adapter) ListImages(req ListImageRequest) ([]Image, error) {

	link1 := fmt.Sprintf("https://api.unsplash.com/photos?client_id=%v&page=%v&per_page=%v&order_by=%v", adt.apiKey, req.Page, req.Per_page, req.Order_by)
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
		return nil, err
	}
	return images, nil
}
