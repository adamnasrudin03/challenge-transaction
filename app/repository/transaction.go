package repository

import (
	"adamnasrudin03/challenge-transaction/app/dto"
	"adamnasrudin03/challenge-transaction/app/entity"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TransactionRepository is contract what TransactionRepository can do to db
type TransactionRepository interface {
	Create(input entity.Transaction) (res entity.Transaction, err error)
	GetAll(ctx *gin.Context, queryparam dto.ParamTransactions) (result []entity.Transaction, total int64, err error)
}

type initTransactionRepo struct {
	DB *gorm.DB
}

// NewTransactionRepository is creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &initTransactionRepo{
		DB: db,
	}
}

func (repo *initTransactionRepo) Create(input entity.Transaction) (res entity.Transaction, err error) {
	if err := repo.DB.Create(&input).Error; err != nil {
		log.Printf("[TransactionRepository-Create] error create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *initTransactionRepo) GetAll(ctx *gin.Context, queryparam dto.ParamTransactions) (result []entity.Transaction, total int64, err error) {
	offset := queryparam.Limit * (queryparam.Page - 1)

	query := repo.DB.WithContext(ctx)
	err = query.Model(&entity.Transaction{}).Count(&total).Error
	if err != nil {
		log.Printf("[TransactionRepository-GetAll] error count total data: %+v \n", err)
		return
	}

	err = query.Offset(offset).Limit(queryparam.Limit).Find(&result).Error
	if err != nil {
		log.Printf("[TransactionRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}
