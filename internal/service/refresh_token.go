package service

import (
	"context"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
)

func (s *service) RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (res *model.RefreshTokenResponse, err error) {
	// Get jwks.
	jwks, err := s.Jwks(ctx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Extract and validate token string.
	claims, ok := util.ExtractClaimsFromString(ctx, req.RefreshToken, jwks)
	if !ok {
		err = constant.ErrInvalidToken
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Construct user model from claims.
	user := &model.User{
		CommonModel: model.CommonModel{
			Id: claims.UserId,
		},
		Name:        claims.Name,
		Email:       claims.Email,
		PhoneNumber: claims.PhoneNumber,
	}

	// Generate new access_token.
	accessToken, err := s.generateToken(ctx, &model.GenerateTokenRequest{
		User:      user,
		TokenType: constant.TokenTypeAccess,
		Extended:  false,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Generate new refresh_token.
	refreshToken, err := s.generateToken(ctx, &model.GenerateTokenRequest{
		User:      user,
		TokenType: constant.TokenTypeRefresh,
		Extended:  claims.Extended,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Construct response.
	res = &model.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}
