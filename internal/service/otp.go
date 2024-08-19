package service

import (
	"context"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"go.uber.org/zap"
)

func (s *service) VerifyOTP(ctx context.Context, req *model.VerifyOTPRequest) (res *model.VerifyOTPResponse, err error) {
	switch req.Action {
	case constant.VOTPARegister:
		return s.verifyOTPRegister(ctx, req)
	case constant.VOTPALogin:
		return s.verifyOTPLogin(ctx, req)
	default:
		err = constant.ErrUnspecifiedAction
		return
	}
}

func (s *service) verifyOTPRegister(ctx context.Context, req *model.VerifyOTPRequest) (res *model.VerifyOTPResponse, err error) {
	// Get cache.
	registerData, err := s.repository.redis.GetOTPRegister(ctx, req.PhoneNumber)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Validate cache existence.
	if registerData == nil {
		err = constant.ErrNotFound
		util.LogContext(ctx).Warn("Cache not found")
		return
	}

	// Validate request aunthenticity.
	if registerData.AuthCode != req.AuthCode {
		err = constant.ErrUnauthenticated
		util.LogContext(ctx).Warn(
			"Mismatch authCode",
			zap.String("cachedAuthCode", registerData.AuthCode),
			zap.String("requestedAuthCode", req.AuthCode),
		)
		return
	}
	if registerData.OTP != req.OTP {
		err = constant.ErrInvalidOTP
		util.LogContext(ctx).Warn(
			"Mismatch authCode",
			zap.String("cachedOTP", registerData.OTP),
			zap.String("requestedOTP", req.OTP),
		)
		return
	}

	if err = s.repository.redis.DeleteOTPRegister(ctx, req.PhoneNumber); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Design response
	res = &model.VerifyOTPResponse{
		Status: "Success",
		Data:   util.CastStructToMap(registerData.RequestPayload),
	}

	return res, s.registerV2Commit(ctx, registerData.RequestPayload)
}

func (s *service) verifyOTPLogin(ctx context.Context, req *model.VerifyOTPRequest) (res *model.VerifyOTPResponse, err error) {
	// Get cache.
	loginData, err := s.repository.redis.GetOTPLogin(ctx, req.PhoneNumber)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Validate cache existence.
	if loginData == nil {
		err = constant.ErrNotFound
		util.LogContext(ctx).Warn("Cache not found")
		return
	}

	// Validate request aunthenticity.
	if loginData.AuthCode != req.AuthCode {
		err = constant.ErrUnauthenticated
		util.LogContext(ctx).Warn(
			"Mismatch authCode",
			zap.String("cachedAuthCode", loginData.AuthCode),
			zap.String("requestedAuthCode", req.AuthCode),
		)
		return
	}
	if loginData.OTP != req.OTP {
		err = constant.ErrInvalidOTP
		util.LogContext(ctx).Warn(
			"Mismatch authCode",
			zap.String("cachedOTP", loginData.OTP),
			zap.String("requestedOTP", req.OTP),
		)
		return
	}

	token, err := s.loginV2Commit(ctx, loginData)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if err = s.repository.redis.DeleteOTPLogin(ctx, req.PhoneNumber); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Design response
	res = &model.VerifyOTPResponse{
		Status: "Success",
		Data:   util.CastStructToMap(token),
	}

	return
}
