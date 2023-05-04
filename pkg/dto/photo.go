package dto

type Photo struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Url            string `json:"url"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Likes          int    `json:"likes"`
}

type ListPhotosRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type ListPhotosResponse struct {
	Data    []*Photo `json:"data"`
	Message string   `json:"message"`
}

type FetchUnsplashPhotosRequest struct {
	Quantity int `json:"quantity"`
}

type FetchUnsplashPhotosResponse struct {
	Data    []*Photo `json:"data"`
	Message string   `json:"message"`
}
