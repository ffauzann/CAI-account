package repository

import (
	"context"
	"database/sql"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/jmoiron/sqlx"
)

func (r *dbRepository) GetListAccountByUserId(ctx context.Context, userId uint64) (accounts []*model.Account, err error) {
	query := `SELECT * FROM "account" a WHERE a.user_id = $1 AND a.deleted_at IS NULL`

	if err = r.db.SelectContext(ctx, &accounts, query, userId); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) GetAccountById(ctx context.Context, id uint64) (account *model.Account, err error) {
	account = new(model.Account)
	query := `SELECT * FROM "account" a WHERE a.id = $1 AND a.deleted_at IS NULL`

	if err = r.db.QueryRowxContext(ctx, query, id).StructScan(account); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) GetAccountByCategory(ctx context.Context, userId uint64, category constant.AccountCategory) (account *model.Account, err error) {
	account = new(model.Account)
	query := `SELECT * FROM "account" a WHERE a.user_id = $1 AND a.category = $2 AND a.deleted_at IS NULL`

	if err = r.db.QueryRowxContext(ctx, query, userId, category).StructScan(account); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) CreateAccount(ctx context.Context, account *model.Account, tx *sqlx.Tx) (err error) {
	if tx == nil { // End tx as soon as this method finished if tx was not provided.
		defer func() { r.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.
	}
	tx, err = r.useOrInitTx(ctx, tx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	query := `
	INSERT INTO 
	"account"(user_id, category, balance, created_by, updated_by)
	VALUES(:user_id, :category, :balance, :user_id, :user_id)
	RETURNING id
	`

	query, args, err := tx.BindNamed(query, account)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	err = tx.QueryRowxContext(ctx, query, args...).Scan(&account.Id)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) UpdateBalance(ctx context.Context, accountId uint64, amount float64, tx *sqlx.Tx) (err error) {
	if tx == nil { // End tx as soon as this method finished if tx was not provided.
		defer func() { r.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.
	}
	tx, err = r.useOrInitTx(ctx, tx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Update account balance.
	if err = r.updateAccountBalance(ctx, accountId, amount, tx); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Create audit trail.
	if err = r.createAccountHistory(ctx, accountId, amount, tx); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) updateAccountBalance(ctx context.Context, accountId uint64, amount float64, tx *sqlx.Tx) (err error) {
	if tx == nil { // End tx as soon as this method finished if tx was not provided.
		defer func() { r.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.
	}
	tx, err = r.useOrInitTx(ctx, tx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	query := `UPDATE "account" SET balance = balance + :amount WHERE id = :account_id`
	arg := map[string]interface{}{
		"account_id": accountId,
		"amount":     amount,
	}

	if _, err = tx.NamedExecContext(ctx, query, arg); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) createAccountHistory(ctx context.Context, accountId uint64, amount float64, tx *sqlx.Tx) (err error) {
	if tx == nil { // End tx as soon as this method finished if tx was not provided.
		defer func() { r.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.
	}
	tx, err = r.useOrInitTx(ctx, tx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	query := `INSERT INTO "account_history"(account_id, amount) VALUES(:account_id, :amount)`
	arg := map[string]interface{}{
		"account_id": accountId,
		"amount":     amount,
	}

	if _, err = tx.NamedExecContext(ctx, query, arg); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}
