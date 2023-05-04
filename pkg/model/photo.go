package model

type Photo struct {
	ID             string `gorm:"column:id"`
	CreatedAt      string `gorm:"column:created_at"`
	UpdatedAt      string `gorm:"column:updated_at"`
	Width          int    `gorm:"column:width"`
	Height         int    `gorm:"column:height"`
	Url            string `gorm:"column:url"`
	Description    string `gorm:"column:description"`
	AltDescription string `gorm:"column:alt_description"`
	Likes          int    `gorm:"column:likes"`
}

func (Photo) TableName() string {
	return "photos"
}
