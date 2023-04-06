package dto

import (
	"adamnasrudin03/challenge-transaction/app/entity"
)

type TransactionReq struct {
	RequestID uint64               ` json:"request_id"`
	Data      []entity.Transaction ` json:"data"`
}

type ParamTransactions struct {
	Page   uint64 `json:"page" valid:"Required"`
	Limit  uint64 `json:"limit" valid:"Required"`
	Offset uint64 `json:"offset"`
}
