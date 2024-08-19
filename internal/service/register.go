package service

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
)

func (s *service) Register(ctx context.Context, req *model.RegisterRequest) (res *model.RegisterResponse, err error) {
	// Validate user existence.
	isUserExist, err := s.IsUserExist(ctx, &model.IsUserExistRequest{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if isUserExist.IsExist {
		res = &model.RegisterResponse{
			StatusCode: constant.RSCFailed,
			Reasons:    isUserExist.Reasons,
		}
		return
	}

	// Begin tx.
	tx, err := s.repository.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	defer func() { s.repository.db.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.

	// Begin to register new user.
	user, err := s.createUser(ctx, tx, req)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Open debit account automatically.
	_, err = s.openAccount(ctx, &model.OpenAccountRequest{
		UserId:   user.Id,
		Category: constant.AccountCategoryDebit,
		Balance:  0,
	}, tx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = &model.RegisterResponse{
		StatusCode: constant.RSCSucceed,
	}

	return
}

func (s *service) createUser(ctx context.Context, tx *sqlx.Tx, req *model.RegisterRequest) (user *model.User, err error) {
	// Prepare concurrent for hashing due it could take quite some times.
	var wg sync.WaitGroup
	chErr := make(chan error, 1) //nolint
	fnHash := func(pwd, credType string) {
		defer wg.Done()
		hashed, err := util.HashBCrypt(pwd)
		if err != nil {
			chErr <- err
			return
		}
		switch credType {
		case "password":
			req.UserPassword = hashed
		case "passcode":
			// req.UserPasscode = hashed
		}
	}

	// Begin concurrent.
	wg.Add(1) //nolint
	// go fnHash(s.config.Encryption.MasterPassword, constant.MasterPasswordType)
	go fnHash(req.PlainPassword, "password")
	// go fnHash(req.UserPasscode, "passcode")
	wg.Wait()

	// Begin non-blocking read channel.
	select {
	case err = <-chErr: // Error occurred
		util.LogContext(ctx).Error(err.Error())
		return
	default: // No error, moving on.
	}

	user = util.CastStruct[model.User](req)
	user.Password = req.UserPassword
	// user.Passcode = req.User.UserPasscode
	user.Status = constant.UserStatusActive
	if err = s.repository.db.CreateUser(ctx, user, tx); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (s *service) RegisterV2(ctx context.Context, req *model.RegisterV2Request) (res *model.RegisterV2Response, err error) {
	t := now()

	res, err = s.validateIdempotencyRegisterV2(ctx, req)
	if err != nil || res != nil {
		util.LogContext(ctx).Warn("Possible idempotent")
		return
	}

	authCode, err := s.sendOTPRegister(ctx, req, t)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = &model.RegisterV2Response{
		Status:   "Success",
		AuthCode: authCode,
	}

	return
}

// validateIdempotencyRegisterV2 validates whether the same phone number is already requested to register within specific time period.
func (s *service) validateIdempotencyRegisterV2(ctx context.Context, req *model.RegisterV2Request) (res *model.RegisterV2Response, err error) {
	existingReq, err := s.repository.redis.GetOTPRegister(ctx, req.PhoneNumber)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if existingReq != nil {
		util.LogContext(ctx).Warn("Duplicate request")
		res = &model.RegisterV2Response{
			Status:   "Success",
			AuthCode: existingReq.AuthCode,
		}
		return
	}

	// Validate user existence.
	isUserExist, err := s.IsUserExist(ctx, &model.IsUserExistRequest{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if isUserExist.IsExist {
		err = constant.NewDynamicError(codes.AlreadyExists, isUserExist.Reasons[0])
		util.LogContext(ctx).Warn(err.Error())
		return
	}

	return
}

func (s *service) createUserV2(ctx context.Context, tx *sqlx.Tx, req *model.RegisterV2Request) (user *model.User, err error) {
	user = util.CastStruct[model.User](req)
	user.Status = constant.UserStatusActive

	if err = s.repository.db.CreateUser(ctx, user, tx); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (s *service) sendOTPRegister(ctx context.Context, req *model.RegisterV2Request, t time.Time) (authCode string, err error) {
	var otp, mockOTP string
	mockOTP = s.config.Dependency.Whatsapp.MockOTP
	otp = mockOTP

	// Generate & send OTP if not mocked.
	if mockOTP == "" {
		otp = util.RandomNumericSequence(t.UnixNano(), constant.DefaultOTPLength)
		content := fmt.Sprintf(s.config.Dependency.Whatsapp.RegisterOTP.Content, otp)
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
	err = s.repository.redis.SetOTPRegister(ctx, &model.RedisSetOTPRegisterData{
		PhoneNumber:    req.PhoneNumber,
		AuthCode:       authCode,
		OTP:            otp,
		RetryCount:     0,
		LastResend:     t,
		RequestPayload: req,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (s *service) registerV2Commit(ctx context.Context, req *model.RegisterV2Request) (err error) {
	// Begin tx.
	tx, err := s.repository.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	defer func() { s.repository.db.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.

	// Begin to register new user
	_, err = s.createUserV2(ctx, tx, req)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	err = s.repository.redis.DeleteOTPRegister(ctx, req.PhoneNumber)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}
