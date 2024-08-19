package repository

import (
	"context"
	"database/sql"

	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/jmoiron/sqlx"
)

func (r *dbRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sqlx.Tx, err error) {
	tx, err = r.db.BeginTxx(ctx, opts)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) EndTx(ctx context.Context, tx *sqlx.Tx, err error) {
	if tx == nil {
		util.LogContext(ctx).Error("No active tx.")
		return
	}

	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			util.LogContext(ctx).Error(errRollback.Error())
			return
		}
	}

	if errCommit := tx.Commit(); errCommit != nil {
		util.LogContext(ctx).Error(errCommit.Error())
		return
	}
}

func (r *dbRepository) useOrInitTx(ctx context.Context, tx *sqlx.Tx) (*sqlx.Tx, error) {
	if tx != nil {
		return tx, nil
	}

	return r.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
}
