package adapter

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
}
type Image struct {
	Description string `json:"description"`

	Urls  Urls `json:"urls"`
	Likes int  `json:"likes"`
}
type InfoLink struct {
	page     int
	per_page int
}
