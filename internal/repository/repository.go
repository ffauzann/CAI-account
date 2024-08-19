package repository

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewDB(db *sqlx.DB, config *model.AppConfig, logger *zap.Logger) DBRepository {
	return &dbRepository{
		db: db,
		common: common{
			config: config,
			logger: logger,
		},
	}
}

func NewRedis(client *redis.Client, config *model.AppConfig, logger *zap.Logger) RedisRepository {
	return &redisRepository{
		redis: client,
		common: common{
			config: config,
			logger: logger,
		},
	}
}

func NewWhatsappClient(config *model.AppConfig, logger *zap.Logger) WhatsappClientRepository {
	return &whatsappClientRepository{
		client: &http.Client{},
		common: common{
			config: config,
			logger: logger,
		},
	}
}

type DBRepository interface {
	DBTxRepository
	DBUserRepository
	DBAccountRepository
}

type DBTxRepository interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sqlx.Tx, err error)
	EndTx(ctx context.Context, tx *sqlx.Tx, err error)
}

type DBUserRepository interface {
	CreateUser(ctx context.Context, user *model.User, tx *sqlx.Tx) error
	IsUserExist(ctx context.Context, userIdType constant.UserIdType, userIdVal string) (isExist bool, err error)
	GetUserByOneOfIdentifier(ctx context.Context, val string) (user *model.User, err error)
	CloseUserAccount(ctx context.Context, req *model.CloseUserAccountRequest, tx *sqlx.Tx) (err error)
}

type DBAccountRepository interface {
	GetListAccountByUserId(ctx context.Context, userId uint64) (accounts []*model.Account, err error)
	GetAccountByCategory(ctx context.Context, userId uint64, category constant.AccountCategory) (account *model.Account, err error)
	GetAccountById(ctx context.Context, id uint64) (account *model.Account, err error)
	CreateAccount(ctx context.Context, account *model.Account, tx *sqlx.Tx) (err error)
	UpdateBalance(ctx context.Context, accountId uint64, amount float64, tx *sqlx.Tx) (err error)
}

type RedisRepository interface {
	RegisterUserDevice(ctx context.Context, deviceId string, token *model.Token) error

	SetOTPRegister(ctx context.Context, data *model.RedisSetOTPRegisterData) (err error)
	GetOTPRegister(ctx context.Context, phoneNumber string) (data *model.RedisSetOTPRegisterData, err error)
	DeleteOTPRegister(ctx context.Context, phoneNumber string) (err error)

	SetOTPLogin(ctx context.Context, data *model.RedisSetOTPLoginData) (err error)
	GetOTPLogin(ctx context.Context, phoneNumber string) (data *model.RedisSetOTPLoginData, err error)
	DeleteOTPLogin(ctx context.Context, phoneNumber string) (err error)
}

type WhatsappClientRepository interface {
	Send(ctx context.Context, req *model.WhatsappClientSendTextRequest) (err error)
}

type common struct {
	config *model.AppConfig
	logger *zap.Logger
}

type dbRepository struct {
	db *sqlx.DB
	common
}

type redisRepository struct {
	redis *redis.Client
	common
}

type whatsappClientRepository struct {
	client *http.Client
	common
}

var now = time.Now // For mocking purpose later.
