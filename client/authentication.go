package client

import (
	"context"

	"github.com/ffauzann/CAI-account/proto/gen"
)

func (c *userClient) Register(ctx context.Context, req *gen.RegisterRequest) (*gen.RegisterResponse, error) {
	return c.authClient.Register(ctx, req)
}

func (c *userClient) RegisterV2(ctx context.Context, req *gen.RegisterV2Request) (*gen.RegisterV2Response, error) {
	return c.authClient.RegisterV2(ctx, req)
}

func (c *userClient) VerifyOTP(ctx context.Context, req *gen.VerifyOTPRequest) (*gen.VerifyOTPResponse, error) {
	return c.authClient.VerifyOTP(ctx, req)
}

func (c *userClient) Login(ctx context.Context, req *gen.LoginRequest) (*gen.LoginResponse, error) {
	return c.authClient.Login(ctx, req)
}

func (c *userClient) LoginV2(ctx context.Context, req *gen.LoginV2Request) (*gen.LoginV2Response, error) {
	return c.authClient.LoginV2(ctx, req)
}

func (c *userClient) RefreshToken(ctx context.Context, req *gen.RefreshTokenRequest) (*gen.RefreshTokenResponse, error) {
	return c.authClient.RefreshToken(ctx, req)
}
