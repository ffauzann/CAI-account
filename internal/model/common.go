package model

import (
	"database/sql"
	"time"
)

type CommonModel struct {
	Id uint64 `json:"id" db:"id"`

	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	CreatedBy sql.NullInt64 `json:"created_by" db:"created_by"`

	UpdatedAt sql.NullTime  `json:"updated_at" db:"updated_at"`
	UpdatedBy sql.NullInt64 `json:"updated_by" db:"updated_by"`

	DeletedAt sql.NullTime  `json:"deleted_at" db:"deleted_at"`
	DeletedBy sql.NullInt64 `json:"deleted_by" db:"deleted_by"`
}

type PaginationMD struct {
	Search    string `json:"search"`
	Limit     uint64 `json:"limit"`
	Page      uint64 `json:"page"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`
	Next      bool   `json:"next"`
	Previous  bool   `json:"previous"`
}
