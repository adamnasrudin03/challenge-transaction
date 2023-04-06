package app

import (
	"adamnasrudin03/challenge-transaction/app/repository"
	"adamnasrudin03/challenge-transaction/app/service"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		Transaction: repository.NewTransactionRepository(db),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		Transaction: service.NewTransactionService(repo.Transaction),
	}
}
