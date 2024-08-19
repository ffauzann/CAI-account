package constant

type VerifyOTPAction uint8

const (
	VOTPAUnspecified VerifyOTPAction = iota
	VOTPARegister
	VOTPALogin
)

const (
	DefaultOTPLength      = 6
	DefaultAuthCodeLength = 16
)
