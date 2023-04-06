package service

import (
	"adamnasrudin03/challenge-transaction/app/dto"
	"adamnasrudin03/challenge-transaction/app/entity"
	"adamnasrudin03/challenge-transaction/app/repository"
	"log"
	"math"

	"github.com/gin-gonic/gin"
)

// TransactionService is a contract about something that this service can do
type TransactionService interface {
	GetTransactions(ctx *gin.Context, queryparam dto.ParamTransactions) (result dto.ResponseList, err error)
	Create(ctx *gin.Context, input dto.TransactionReq) (res entity.Transaction, err error)
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
func (srv *initTransactionService) GetTransactions(ctx *gin.Context, queryparam dto.ParamTransactions) (result dto.ResponseList, err error) {
	result.Limit = queryparam.Limit
	result.Page = queryparam.Page

	result.Data, result.Total, err = srv.transactionRepository.GetAll(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.LastPage = uint64(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, nil
}

func (srv *initTransactionService) Create(ctx *gin.Context, input dto.TransactionReq) (res entity.Transaction, err error) {

	for _, v := range input.Data {
		go func(transaktion entity.Transaction) {
			log.Println("Running process request_id: ", input.RequestID)
			// transaktion.ID = 0
			_, _ = srv.transactionRepository.Create(transaktion)
		}(v)
	}

	return entity.Transaction{}, nil
}
