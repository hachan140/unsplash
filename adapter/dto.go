package adapter

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
type adapter struct {
	apiKey string
}
