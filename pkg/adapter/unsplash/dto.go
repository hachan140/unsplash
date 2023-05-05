package unsplash

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
}
type Photo struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Urls           Urls   `json:"urls"`
	Likes          int    `json:"likes"`
}
type ListPhotoRequest struct {
	Page    int
	PerPage int
	OrderBy PhotoOrder
}
type PhotoOrder string

const (
	PhotoOrderLatest  PhotoOrder = "latest"
	PhotoOrderOldest  PhotoOrder = "oldest"
	PhotoOrderPopular PhotoOrder = "popular"
)
