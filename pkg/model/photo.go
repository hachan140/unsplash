package model

type Photo struct {
	ID             string
	CreatedAt      string
	UpdatedAt      string
	Width          int
	Height         int
	Url            string
	Description    string
	AltDescription string
	Likes          int
}
