package model

import "time"

type User struct {
	ID          int       `gorm:"column:id"`
	Username    string    `gorm:"column:username"`
	Password    string    `gorm:"column:password"`
	FullName    string    `gorm:"column:full_name"`
	PhoneNumber string    `gorm:"column:phone_number"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
