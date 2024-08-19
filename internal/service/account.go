package service

import (
	"context"
	"database/sql"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/jmoiron/sqlx"
)

func (s *service) GetListAccount(ctx context.Context, req *model.GetListAccountRequest) (res *model.GetListAccountResponse, err error) {
	accounts, err := s.repository.db.GetListAccountByUserId(ctx, req.UserId)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = new(model.GetListAccountResponse)
	res.Accounts = util.CastSlice[[]*model.GetListAccountDataResponse](accounts)

	return
}

func (s *service) openAccount(ctx context.Context, req *model.OpenAccountRequest, tx *sqlx.Tx) (account *model.Account, err error) {
	// Check whether account under the same category is exist.
	account, err = s.repository.db.GetAccountByCategory(ctx, req.UserId, req.Category)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Return an existing account if exist.
	if account != nil {
		return account, nil
	}

	// Create account if not exist.
	if err = s.repository.db.CreateAccount(ctx, util.CastStruct[model.Account](req), tx); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (s *service) UpdateBalance(ctx context.Context, req *model.UpdateBalanceRequest) (res *model.UpdateBalanceResponse, err error) {
	// Begin tx.
	tx, err := s.repository.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	defer func() { s.repository.db.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.

	// Validate request and get source & destination account details.
	sourceAccount, destinationAccount, err := s.validateUpdateBalance(ctx, req)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Update source account balance.
	if err = s.repository.db.UpdateBalance(ctx, sourceAccount.Id, req.Amount, tx); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Update destination account balance if set.
	if req.DestinationAccountId != 0 {
		if err = s.repository.db.UpdateBalance(ctx, destinationAccount.Id, (req.Amount * -1), tx); err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}
	}

	// Construct response.
	res = &model.UpdateBalanceResponse{
		Status: "OK",
	}

	return
}

func (s *service) validateUpdateBalance(ctx context.Context, req *model.UpdateBalanceRequest) (sourceAccount, destinationAccount *model.Account, err error) {
	// Check whether source account is exist.
	sourceAccount, err = s.repository.db.GetAccountById(ctx, req.SourceAccountId)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Validate existence & ownership.
	if sourceAccount == nil || sourceAccount.UserId != req.UserId {
		err = constant.ErrAccountNotFound
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Validate amount.
	sourceAccount.Balance += req.Amount
	if sourceAccount.Balance < 0 {
		err = constant.ErrInsufficientAccountBalance
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Check whether destination account is set.
	if req.DestinationAccountId != 0 {
		// Check whether destination account is exist.
		destinationAccount, err = s.repository.db.GetAccountById(ctx, req.DestinationAccountId)
		if err != nil {
			util.LogContext(ctx).Error(err.Error())
			return
		}

		// destination account is not found.
		if destinationAccount == nil {
			err = constant.ErrAccountNotFound
			util.LogContext(ctx).Error(err.Error())
			return
		}

		// Validate amount.
		destinationAccount.Balance += (req.Amount * -1)
		if destinationAccount.Balance < 0 {
			err = constant.ErrInsufficientAccountBalance
			util.LogContext(ctx).Error(err.Error())
			return
		}
	}

	return
}
