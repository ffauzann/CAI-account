package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/jmoiron/sqlx"
)

func (r *dbRepository) CreateUser(ctx context.Context, user *model.User, tx *sqlx.Tx) (err error) {
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
	"user"(name, email, username, phone_number, role_id, password, passcode, status, is_email_verified)
	VALUES(:name, :email, :username, :phone_number, :role_id, :password, :passcode, :status, :is_email_verified)
	RETURNING id
	`

	query, args, err := tx.BindNamed(query, user)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	err = tx.QueryRowxContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) IsUserExist(ctx context.Context, userIdType constant.UserIdType, userIdVal string) (isExist bool, err error) {
	if err = userIdType.Validate(); err != nil {
		return
	}

	var count int
	query := fmt.Sprintf(`SELECT COUNT(1) FROM "user" WHERE %s = $1 AND status != 'CLOSED'`, userIdType)
	if err = r.db.QueryRowxContext(ctx, query, userIdVal).Scan(&count); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return count > 0, nil
}

// GetUserByOneOfIdentifier returns user data if there's any match with one the username/email/phone_number values.
func (r *dbRepository) GetUserByOneOfIdentifier(ctx context.Context, val string) (user *model.User, err error) {
	user = new(model.User)
	query := `SELECT 
			u.id,
			u.name,
			u.email,
			u.phone_number,
			u.role_id,
			u.password,
			u.status
		FROM "user" u
		LEFT JOIN "role" r ON r.id = u.role_id 
		WHERE 
			(email = $1 OR phone_number = $1)
			AND u.deleted_at IS NULL
		GROUP BY u.id
		ORDER BY status
		LIMIT 1`

	if err = r.db.QueryRowxContext(ctx, query, val).StructScan(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, constant.ErrUserNotFound
		}
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) CloseUserAccount(ctx context.Context, req *model.CloseUserAccountRequest, tx *sqlx.Tx) (err error) {
	if tx == nil { // End tx as soon as this method finished if tx was not provided.
		defer func() { r.EndTx(ctx, tx, err) }() // This EndTx must NOT be called directly with defer since the arguments to deferred functions will be evaluated immediately.
	}
	tx, err = r.useOrInitTx(ctx, tx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	query := `
		UPDATE "user"
		SET
			status = :status,
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE
			id = :id AND deleted_at IS NULL
	`

	arg := map[string]interface{}{
		"id":         req.UserId,
		"status":     constant.UserStatusClosed,
		"updated_at": now(),
		"updated_by": req.UserId,
	}

	_, err = tx.NamedExecContext(ctx, query, arg)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}
