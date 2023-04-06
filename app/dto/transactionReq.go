package dto

import (
	"adamnasrudin03/challenge-transaction/app/entity"
)

type TransactionReq struct {
	RequestID uint64               ` json:"request_id"`
	Data      []entity.Transaction ` json:"data"`
}

type ParamTransactions struct {
	Page   int `json:"page" valid:"Required"`
	Limit  int `json:"limit" valid:"Required"`
	Offset int `json:"offset"`
}
