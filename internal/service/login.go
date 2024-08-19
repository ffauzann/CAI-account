package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req *model.LoginRequest) (res *model.LoginResponse, err error) {
	user, err := s.repository.db.GetUserByOneOfIdentifier(ctx, req.UserId)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if user.Status != constant.UserStatusActive {
		return nil, constant.ErrUserIsNotActive
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, constant.ErrInvalidUsernamePassword
		}
		util.LogContext(ctx).Error(err.Error())
		return
	}

	var token model.Token
	// Generate access_token
	token.AccessToken, err = s.generateToken(ctx, &model.GenerateTokenRequest{
		User:      user,
		TokenType: constant.TokenTypeAccess,
		Extended:  false,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Generate refresh_token
	token.RefreshToken, err = s.generateToken(ctx, &model.GenerateTokenRequest{
		User:      user,
		TokenType: constant.TokenTypeRefresh,
		Extended:  req.RememberMe,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = &model.LoginResponse{
		Token: token,
	}

	return
}

func (s *service) generateToken(ctx context.Context, req *model.GenerateTokenRequest) (token string, err error) { //nolint
	// Preserve variables.
	var (
		iss, exp, pk string
		d            time.Duration
		n            int64
	)

	// Predefine index value of keys.
	n = util.RandomNumericWithinRange(now().UnixNano(), 0, int64(len(s.config.Jwt.AsymmetricKeys)-1))

	// Determine which config to fetch based on its type.
	switch req.TokenType {
	case constant.TokenTypeAccess:
		iss = s.config.Jwt.AccessToken.Iss
		exp = s.config.Jwt.AccessToken.Exp
		pk = s.config.Jwt.AsymmetricKeys[n].PrivateKey

		d, err = time.ParseDuration(exp)
		if err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}
	case constant.TokenTypeRefresh:
		iss = s.config.Jwt.RefreshToken.Iss
		exp = s.config.Jwt.RefreshToken.Exp
		pk = s.config.Jwt.AsymmetricKeys[n].PrivateKey

		if req.Extended {
			exp = s.config.Jwt.RefreshToken.ExtendedExp
		}

		d, err = time.ParseDuration(exp)
		if err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}
	}

	// Decode base64 private key.
	bPrivateKey, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Construct base claims.
	claims := model.Claims{
		UserId:      req.User.Id,
		Name:        req.User.Name,
		Email:       req.User.Email,
		PhoneNumber: req.User.PhoneNumber,
		RoleId:      req.User.RoleId,
		TokenType:   req.TokenType,
		Extended:    req.Extended,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   req.User.Email,
			Issuer:    iss,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		},
	}

	// Create token & sign.
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t.Header["kid"] = s.config.Jwt.AsymmetricKeys[n].Kid

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bPrivateKey)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return t.SignedString(privateKey)
}

func (s *service) LoginV2(ctx context.Context, req *model.LoginV2Request) (res *model.LoginV2Response, err error) {
	t := now()

	res, err = s.validateIdempotencyLoginV2(ctx, req)
	if err != nil || res != nil {
		util.LogContext(ctx).Warn("Possible idempotent")
		return
	}

	user, err := s.getUserLoginV2(ctx, req)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	authCode, err := s.sendOTPLogin(ctx, req, user, t)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = &model.LoginV2Response{
		Status:   "Success",
		AuthCode: authCode,
	}

	return
}

// getUserLoginV2 gets user and groupCodes and also validate user status.
func (s *service) getUserLoginV2(ctx context.Context, req *model.LoginV2Request) (user *model.User, err error) {
	user, err = s.repository.db.GetUserByOneOfIdentifier(ctx, req.PhoneNumber)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if user.Status != constant.UserStatusActive {
		util.LogContext(ctx).Warn("Inactive user")
		err = constant.ErrUserIsNotActive
		return
	}

	return
}

// validateIdempotencyLoginV2 validates whether the same phone number is already requested to login within specific time period.
func (s *service) validateIdempotencyLoginV2(ctx context.Context, req *model.LoginV2Request) (res *model.LoginV2Response, err error) {
	existingReq, err := s.repository.redis.GetOTPLogin(ctx, req.PhoneNumber)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if existingReq != nil {
		util.LogContext(ctx).Warn("Duplicate request")
		res = &model.LoginV2Response{
			Status:   "Success",
			AuthCode: existingReq.AuthCode,
		}
		return
	}

	return
}

// sendOTPLogin sends OTP by Whatsapp and store it's value in cache.
func (s *service) sendOTPLogin(ctx context.Context, req *model.LoginV2Request, user *model.User, t time.Time) (authCode string, err error) {
	var otp, mockOTP string
	mockOTP = s.config.Dependency.Whatsapp.MockOTP
	otp = mockOTP

	// Generate & send OTP if not mocked.
	if mockOTP == "" {
		otp = util.RandomNumericSequence(t.UnixNano(), constant.DefaultOTPLength)
		content := fmt.Sprintf(s.config.Dependency.Whatsapp.LoginOTP.Content, otp)
		err = s.repository.whatsappClient.Send(ctx, &model.WhatsappClientSendTextRequest{
			PhoneNumber: req.PhoneNumber,
			Content:     content,
		})
		if err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}
	}

	// Generate authCode from random sequence.
	authCode = util.RandomAlphaSequence(t.UnixNano(), constant.DefaultAuthCodeLength)
	err = s.repository.redis.SetOTPLogin(ctx, &model.RedisSetOTPLoginData{
		PhoneNumber:    req.PhoneNumber,
		AuthCode:       authCode,
		OTP:            otp,
		RetryCount:     0,
		LastResend:     t,
		RequestPayload: req,
		User:           user,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (s *service) loginV2Commit(ctx context.Context, req *model.RedisSetOTPLoginData) (token *model.Token, err error) {
	token = new(model.Token)

	// Generate access_token
	token.AccessToken, err = s.generateToken(ctx, &model.GenerateTokenRequest{
		User:      req.User,
		TokenType: constant.TokenTypeAccess,
		Extended:  false,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Generate refresh_token
	token.RefreshToken, err = s.generateToken(ctx, &model.GenerateTokenRequest{
		User:      req.User,
		TokenType: constant.TokenTypeRefresh,
		Extended:  req.RequestPayload.RememberMe,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}
