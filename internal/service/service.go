package service

import (
	"context"
	"time"

	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/repository"

	"go.uber.org/zap"
)

type Service interface {
	AuthService
	UserService
	AccountService
}

type AuthService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (res *model.RegisterResponse, err error)
	RegisterV2(ctx context.Context, req *model.RegisterV2Request) (res *model.RegisterV2Response, err error)
	Login(ctx context.Context, req *model.LoginRequest) (res *model.LoginResponse, err error)
	LoginV2(ctx context.Context, req *model.LoginV2Request) (res *model.LoginV2Response, err error)
	RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (res *model.RefreshTokenResponse, err error)
	VerifyOTP(ctx context.Context, req *model.VerifyOTPRequest) (res *model.VerifyOTPResponse, err error)
	Jwks(ctx context.Context) (jwks []*model.Jwk, err error)
}

type UserService interface {
	CloseUserAccount(ctx context.Context, req *model.CloseUserAccountRequest) (res *model.CloseUserAccountResponse, err error)
	IsUserExist(ctx context.Context, req *model.IsUserExistRequest) (res *model.IsUserExistResponse, err error)
}

type AccountService interface {
	GetListAccount(ctx context.Context, req *model.GetListAccountRequest) (res *model.GetListAccountResponse, err error)
	UpdateBalance(ctx context.Context, req *model.UpdateBalanceRequest) (res *model.UpdateBalanceResponse, err error)
}

type service struct {
	config     *model.AppConfig
	logger     *zap.Logger
	repository repositoryWrapper
}

type repositoryWrapper struct {
	db             repository.DBRepository
	redis          repository.RedisRepository
	whatsappClient repository.WhatsappClientRepository
}

func New(db repository.DBRepository, redis repository.RedisRepository, whatsappClient repository.WhatsappClientRepository, config *model.AppConfig, logger *zap.Logger) Service {
	return &service{
		config: config,
		logger: logger,
		repository: repositoryWrapper{
			db:             db,
			redis:          redis,
			whatsappClient: whatsappClient,
		},
	}
}

var now = time.Now // For mocking purpose later.
