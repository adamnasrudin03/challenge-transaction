package dto

import (
	"adamnasrudin03/challenge-transaction/app/entity"
)

type TransactionReq struct {
	RequestID uint64               ` json:"request_id" validate:"required"`
	Data      []entity.Transaction ` json:"data"`
}

type ParamTransactions struct {
	Page   uint64 `json:"page" validate:"required"`
	Limit  uint64 `json:"limit" validate:"required"`
	Offset uint64 `json:"offset"`
}
