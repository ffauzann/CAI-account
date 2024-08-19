package model

import (
	"encoding/json"

	"github.com/ffauzann/CAI-account/internal/constant"

	"github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	AsymmetricKeys JwtAsymmetricKeysConfig
	AccessToken    JwtAccessTokenConfig
	RefreshToken   JwtRefreshTokenConfig
}

type JwtAsymmetricKeysConfig []*struct {
	Kid        string
	PrivateKey string
	PublicKey  string
}

type JwtAccessTokenConfig struct {
	// Deprecated: no longer used since the algorithm has changed to RS256.
	SigningKey string
	Iss        string
	Exp        string
}

type JwtRefreshTokenConfig struct {
	// Deprecated: no longer used since the algorithm has changed to RS256.
	SigningKey  string
	Iss         string
	Exp         string
	ExtendedExp string
}

type Claims struct {
	UserId      uint64             `json:"user_id"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phone_number"`
	RoleId      uint64             `json:"role_id"`
	TokenType   constant.TokenType `json:"token_type"`
	Extended    bool               `json:"extended"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// MarshalBinary fulfills encoding.BinaryMarshaler implementation.
func (t Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

type GenerateTokenRequest struct {
	User      *User
	TokenType constant.TokenType
	Extended  bool
}

type Jwk struct {
	KeyType   string `json:"kty"`
	KeyID     string `json:"kid"`
	Usage     string `json:"use"`
	Algorithm string `json:"alg"`
	Modulus   string `json:"n"`
	Exponent  string `json:"e"`
}
