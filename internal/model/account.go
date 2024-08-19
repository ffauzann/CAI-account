package model

import "github.com/ffauzann/CAI-account/internal/constant"

type Account struct {
	CommonModel

	UserId   uint64                   `json:"user_id" db:"user_id"`
	Category constant.AccountCategory `json:"category" db:"category"`
	Balance  float64                  `json:"balance" db:"balance"`
}

type GetListAccountRequest struct {
	UserId uint64
}

type GetListAccountDataResponse struct {
	Id       uint64                   `json:"id"`
	Category constant.AccountCategory `json:"category"`
	Balance  float64                  `json:"balance"`
}

type GetListAccountResponse struct {
	Accounts []*GetListAccountDataResponse `json:"accounts"`
}

type OpenAccountRequest struct {
	UserId   uint64                   `json:"user_id"`
	Category constant.AccountCategory `json:"category"`
	Balance  float64                  `json:"balance"`
}

type UpdateBalanceRequest struct {
	UserId               uint64
	SourceAccountId      uint64  `json:"source_account_id"`
	DestinationAccountId uint64  `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
}

type UpdateBalanceResponse struct {
	Status string `json:"status"`
}
