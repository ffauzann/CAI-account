package model

type RegisterRequestUserDetail struct {
	Name string `json:"name" validate:"required"`

	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`

	PlainPassword string `json:"password" name:"password" validate:"required,password"`
	UserPassword  string

	PlainPasscode string `json:"passcode" name:"passcode" validate:"required"`
	UserPasscode  string
}

type RegisterRequestGroupDetail struct {
	Code             string `json:"code"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	AllowSupervision bool   `json:"allow_supervision"`
}

type RegisterRequest struct {
	Name          string `json:"name" validate:"required"`
	Username      string `json:"username"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   string `json:"phone_number" validate:"required"`
	PlainPassword string `json:"password" name:"password" validate:"required,password"`
	RoleId        uint64 `json:"role_id"`
	UserPassword  string `json:"-"`
}

type RegisterResponse struct {
	StatusCode int
	Reasons    []string
}

type RegisterV2Request struct {
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Passcode    string `json:"passcode" validate:"required,number,len=6"`
}

type RegisterV2Response struct {
	Status   string `json:"status"`
	AuthCode string `json:"auth_code"`
}
