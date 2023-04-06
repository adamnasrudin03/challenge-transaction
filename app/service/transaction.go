package service

import (
	"adamnasrudin03/challenge-transaction/app/dto"
	"adamnasrudin03/challenge-transaction/app/repository"
	"math"

	"github.com/gin-gonic/gin"
)

// TransactionService is a contract about something that this service can do
type TransactionService interface {
	GetTransactions(ctx *gin.Context, queryparam dto.ParamTransactions) (result dto.ResponseList, err error)
}

type initTransactionService struct {
	transactionRepository repository.TransactionRepository
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &initTransactionService{
		transactionRepository: transactionRepo,
	}
}
func (u *initTransactionService) GetTransactions(ctx *gin.Context, queryparam dto.ParamTransactions) (result dto.ResponseList, err error) {

	result.Data, result.Total, err = u.transactionRepository.GetAll(ctx, queryparam)
	if err != nil {
		return result, err
	}
	result.Limit = uint(queryparam.Limit)
	result.Page = uint(queryparam.Page)

	result.LastPage = uint(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, nil
}