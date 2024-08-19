package model

import "github.com/ffauzann/CAI-account/internal/constant"

type User struct {
	CommonModel

	Name        string `json:"name" db:"name"`
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	RoleId      uint64 `json:"role_id" db:"role_id"`

	Password string `json:"password" db:"password"`
	Passcode string `json:"passcode" db:"passcode"`

	Status          constant.UserStatus `json:"status" db:"status"`
	IsEmailVerified bool                `json:"is_email_verified" db:"is_email_verified"`
}

type IsUserExistRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type IsUserExistResponse struct {
	IsExist bool     `json:"is_exist"`
	Reasons []string `json:"reasons"`
}

type CloseUserAccountRequest struct {
	UserId uint64 `json:"-"` // From claims.
}

type CloseUserAccountResponse struct {
	UserId uint64              `json:"user_id"`
	Status constant.UserStatus `json:"status"`
}
