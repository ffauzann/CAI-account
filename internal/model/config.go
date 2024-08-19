package model

// Reusable config goes here.
type AppConfig struct {
	Encryption Encryption
	Jwt        JwtConfig
	Auth       AuthConfig
	Dependency DependencyConfig
}

type Encryption struct {
	Cost uint8
}

type AuthConfig struct {
	ExcludedMethods []string
}

type DependencyConfig struct {
	Whatsapp WhatsappConfig
}

type WhatsappConfig struct {
	SenderURL   string
	XAccessKey  string
	InstanceID  string
	MockOTP     string
	RegisterOTP WhatsappUsecaseConfig
	LoginOTP    WhatsappUsecaseConfig
}

type WhatsappUsecaseConfig struct {
	Content  string
	MaxRetry uint8
	Exp      string
	Cooldown string
}
