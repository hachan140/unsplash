//package adapter
//
//import "fmt"
//
//type Adapter interface {
//	GetImage(info InfoLink) ([]Image, error)
//}
//
//func (img *Image) NewImage(description string, urls Urls, likes int) {
//	img.Description = description
//	img.Urls = urls
//	img.Likes = likes
//}
//
//func (img Image) GetImage(info InfoLink) ([]Image, error) {
//	client_id := "eHYWuHQhaaeo6b-EQ3HNjMQxuabrhYdUdUE57pxs_M8"
//	link := fmt.Sprintf("https://api.unsplash.com/photos?client_id=%v&page=%v&per_page=%v", client_id, info.page, info.per_page)
//
//}
