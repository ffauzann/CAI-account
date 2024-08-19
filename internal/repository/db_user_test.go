package repository

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/ffauzann/common/logger"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) { //nolint
	var (
		ctx    = context.Background()
		logger = logger.Setup(logger.EnvTesting)
		query  = `INSERT INTO users(name, username, email, phone_number, password, master_password) VALUES(?, ?, ?, ?, ?, ?)`
		user   = &model.User{
			Name:        "John Doe",
			Email:       "john@example.com",
			PhoneNumber: "6281222000212",
		}
	)

	// Temp structs
	type (
		arg struct {
			ctx  context.Context
			user *model.User
		}
		want struct {
			err error
		}
		dep struct {
			db sqlmock.Sqlmock
		}
		testModel struct {
			name string
			arg  arg
			want want
			proc func(dep *dep)
		}
	)

	tm := []testModel{
		{
			name: "success",
			arg: arg{
				ctx:  ctx,
				user: user,
			},
			want: want{
				err: nil,
			},
			proc: func(dep *dep) {
				dep.db.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "errConnClosed",
			arg: arg{
				ctx:  ctx,
				user: user,
			},
			want: want{
				err: sql.ErrConnDone,
			},
			proc: func(dep *dep) {
				dep.db.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)
			},
		},
	}

	for _, tt := range tm {
		tt := tt // Prevent race condition.
		t.Run(tt.name, func(t *testing.T) {
			db, sqlMock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := NewDB(sqlxDB, &model.AppConfig{}, logger)

			if tt.proc != nil {
				tt.proc(&dep{db: sqlMock})
			}

			err = repo.CreateUser(tt.arg.ctx, tt.arg.user, nil)
			assert.Equalf(t, tt.want.err, err, "error is not equal.\nexpected: %v\nbut got: %v\n", tt.want.err, err)
		})
	}
}

func TestIsUserExists(t *testing.T) { //nolint
	var (
		ctx    = context.Background()
		logger = logger.Setup(logger.EnvTesting)
	)

	// Temp structs
	type (
		arg struct {
			ctx        context.Context
			userIdType constant.UserIdType
			userIdVal  string
		}
		want struct {
			isExist bool
			err     error
		}
		dep struct {
			db sqlmock.Sqlmock
		}
		testModel struct {
			name string
			arg  arg
			want want
			proc func(dep *dep)
		}
	)

	tm := []testModel{
		{
			name: "success",
			arg: arg{
				ctx:        ctx,
				userIdType: constant.UserIdTypeEmail,
				userIdVal:  "john@example.com",
			},
			want: want{
				isExist: true,
				err:     nil,
			},
			proc: func(dep *dep) {
				query := fmt.Sprintf("SELECT COUNT(1) FROM users WHERE %s = ?", constant.UserIdTypeEmail)
				dep.db.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(sqlmock.NewRows([]string{"COUNT(1)"}).AddRow(1))
			},
		},
		{
			name: "errInvalidUserIdType",
			arg: arg{
				ctx:        ctx,
				userIdType: "random string",
				userIdVal:  "johndoe",
			},
			want: want{
				err: constant.ErrInvalidUserIdType,
			},
			proc: func(dep *dep) {},
		},
		{
			name: "errConnClosed",
			arg: arg{
				ctx:        ctx,
				userIdType: constant.UserIdTypeEmail,
				userIdVal:  "john@example.com",
			},
			want: want{
				err: sql.ErrConnDone,
			},
			proc: func(dep *dep) {
				query := fmt.Sprintf("SELECT COUNT(1) FROM users WHERE %s = ?", constant.UserIdTypeEmail)
				dep.db.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)
			},
		},
	}

	for _, tt := range tm {
		tt := tt // Prevent race condition.
		t.Run(tt.name, func(t *testing.T) {
			db, sqlMock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := NewDB(sqlxDB, &model.AppConfig{}, logger)

			if tt.proc != nil {
				tt.proc(&dep{db: sqlMock})
			}

			isExist, err := repo.IsUserExist(tt.arg.ctx, tt.arg.userIdType, tt.arg.userIdVal)
			assert.Equalf(t, tt.want.isExist, isExist, "return is not equal.\nexpected: %v\nbut got: %v\n", tt.want.isExist, isExist)
			assert.Equalf(t, tt.want.err, err, "error is not equal.\nexpected: %v\nbut got: %v\n", tt.want.err, err)
		})
	}
}
