package dto

import "time"

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
type CreateUserResponse struct {
	Data    *User  `json:"data"`
	Message string `json:"message"`
}
