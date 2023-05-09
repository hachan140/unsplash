package dto

import (
	"errors"
	"time"
)

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

func (r CreateUserRequest) Validate() error {
	if r.Username == "" {
		return errors.New("invalid username")
	}
	if r.Password == "" {
		return errors.New("invalid password")
	}
	if r.FullName == "" {
		return errors.New("invalid fullname")
	}
	if r.PhoneNumber == "" {
		return errors.New("invalid phonenumber")
	}
	return nil
}

type CreateUserResponse struct {
	Data    *User  `json:"data"`
	Message string `json:"message"`
}

type ListUsersByUsernameAndPhoneNumberRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Username    string `form:"username"`
	PhoneNumber string `form:"phone_number"`
}

func (l *ListUsersByUsernameAndPhoneNumberRequest) Validate() error {
	if l.Limit <= 0 {
		return errors.New("invalid limit")
	}
	if l.Page <= 0 {
		return errors.New("invalid page")
	}
	return nil
}

type ListUserByUsernameAndPhoneNumberResponse struct {
	Data    []*User `json:"data"`
	Message string  `json:"message"`
}

type DeleteUserByUsernameRequest struct {
	Username string `form:"username"`
}
type DeleteUserByUsernameResponse struct {
	Message string `json:"message"`
}
