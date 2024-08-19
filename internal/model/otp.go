package model

import (
	"time"

	"github.com/ffauzann/CAI-account/internal/constant"
)

type VerifyOTPRequest struct {
	Action      constant.VerifyOTPAction `json:"action" validate:"required"`
	PhoneNumber string                   `json:"phone_number" validate:"required"`
	AuthCode    string                   `json:"auth_code" validate:"required"`
	OTP         string                   `json:"otp" validate:"required"`
}

type VerifyOTPResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type WhatsappClientSendTextRequest struct {
	PhoneNumber string
	Content     string
}

type WhatsappClientSendTextRequestBody struct {
	InstanceID string `json:"instances_id"`
	Content    string `json:"content"`
	To         string `json:"to"`
}

type WhatsappSendTextResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RedisSetOTPRegisterData struct {
	PhoneNumber    string             `json:"phone_number"`
	AuthCode       string             `json:"auth_code"`
	OTP            string             `json:"otp"`
	RetryCount     uint8              `json:"retry_count"`
	LastResend     time.Time          `json:"last_resend"`
	RequestPayload *RegisterV2Request `json:"request_payload"`
}

type RedisSetOTPLoginData struct {
	PhoneNumber    string          `json:"phone_number"`
	AuthCode       string          `json:"auth_code"`
	OTP            string          `json:"otp"`
	RetryCount     uint8           `json:"retry_count"`
	LastResend     time.Time       `json:"last_resend"`
	RequestPayload *LoginV2Request `json:"request_payload"`
	User           *User           `json:"user"`
}
